// Package filter handles meta data for pagination
package filter

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/uuid"
	"github.com/jinzhu/now"
)

type Page struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

func ParsePage(queryParams url.Values, defaultLimit int) *Page {
	limit := defaultLimit
	page := 1
	if limitQuery := queryParams.Get("limit"); limitQuery != "" {
		limit, _ = strconv.Atoi(limitQuery)
	}
	if pageQuery := queryParams.Get("page"); pageQuery != "" {
		tmp, _ := strconv.Atoi(pageQuery)
		page = int(math.Max(float64(1), float64(tmp)))
	}
	return BuildPage(limit, page)
}

func ParsePageBody(pageBody int, limitBody int, defaultLimit int) *Page {
	limit := defaultLimit
	page := 1
	if limitBody != 0 {
		limit = limitBody
	}
	if pageBody != 0 {
		tmp := pageBody
		page = int(math.Max(float64(1), float64(tmp)))
	}
	return BuildPage(limit, page)
}

func BuildPage(limit int, page int) *Page {
	offset := (page - 1) * limit
	if limit == -1 {
		offset = 0
	}
	return &Page{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}
}

type PageResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}

func (pr PageResponse) Count() int {
	return reflect.ValueOf(pr.Data).Len()
}

type MetaResponse struct {
	AffectedRecords int64 `json:"affected_records"`
	LastPage        int64 `json:"last_page"`
	Count           int   `json:"count"`
	HasNext         bool  `json:"has_next"`
}

func BuildPageResponse(ctx context.Context, page *Page, data interface{}) (*PageResponse, error) {
	// Validate if data is a slice
	if reflect.Slice != reflect.TypeOf(data).Kind() {
		return nil, errors.New("paginated data must be a slice")
	}

	// Extract the value of data
	s := reflect.ValueOf(data)

	// Determine the length of data and the real length of data (remove the last index)
	length := s.Len()

	// Extract the value of data as an interface
	result := s.Interface()

	// If data length is 0, result is set to empty array (not null value)
	if length == 0 {
		result = []struct{}{}
	}

	// Create MetaData
	metaResponse, err := CreateMetadata(ctx, page, s, length)
	if err != nil {
		return nil, err
	}

	return &PageResponse{
		Data: result,
		Meta: metaResponse,
	}, nil
}

func CreateMetadata(ctx context.Context, page *Page, s reflect.Value, length int) (MetaResponse, error) {
	var metaResponse MetaResponse

	if length > 0 {
		// extract data from the first index
		val := s.Index(0).Elem()
		// extract AffectedRecords value as string
		affectedRecordsString := val.FieldByName("AffectedRecords")

		// Handle error when extracting AffectedRecords
		if !affectedRecordsString.IsValid() {
			return metaResponse, httperror.New(errortype.INTERNAL_SERVER, fmt.Errorf("affectedrecords is not in the result struct"))
		}

		// set affectedRecords as integer and assign to metaResponse.AffectedRecords
		affectedRecords := affectedRecordsString.Int()
		metaResponse.AffectedRecords = affectedRecords

		// calculate lastPage value and assign it to metaResponse.LastPage
		lastPage := int64(0)
		if page.Limit < 0 {
			lastPage = 1
		} else if page.Limit > 0 {
			lastPage = int64(math.Ceil(float64(affectedRecords) / float64(page.Limit)))
		}
		metaResponse.LastPage = lastPage

		// determine hasNext value and assign it to metaResponse.HasNext
		hasNext := page.Page < int(lastPage)
		if page.Limit == -1 {
			hasNext = false
		}

		metaResponse.HasNext = hasNext
		metaResponse.Count = length
	}

	return metaResponse, nil
}

func ParseBoolOrNil(val string) *bool {
	if val == "" {
		return nil
	}
	var res bool
	if val == "1" || val == "true" {
		res = true
	}
	return &res
}

type TimestampRange struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
}

func (tr TimestampRange) EndDate() *time.Time {
	if tr.End == nil {
		return nil
	}
	endDate := time.Date(tr.End.Year(), tr.End.Month(), tr.End.Day(), 23, 59, 59, 0, tr.End.Location())
	return &endDate
}

func (tr TimestampRange) EndStartDiffDay() int {
	if tr.Start == nil || tr.End == nil {
		return 0
	}
	diff := tr.End.Sub(*tr.Start)
	return int(diff.Hours() / 24)
}

func (tr TimestampRange) LastEndDate() *time.Time {
	if tr.End == nil {
		return nil
	}
	last := tr.EndDate().AddDate(0, 0, -tr.EndStartDiffDay()-1)
	return &last
}

func (tr TimestampRange) LastStartDate() *time.Time {
	if tr.Start == nil {
		return nil
	}
	last := tr.StartDate().AddDate(0, 0, -tr.EndStartDiffDay()-1)
	return &last
}

func (tr TimestampRange) StartDate() *time.Time {
	if tr.Start == nil {
		return nil
	}
	startDate := time.Date(tr.Start.Year(), tr.Start.Month(), tr.Start.Day(), 0, 0, 0, 0, tr.Start.Location())
	return &startDate
}

func ParseTimestampRange(queryParams url.Values, defaultStart *time.Time) (*TimestampRange, error) {
	return ParseTimestampRangeInLocation(queryParams, defaultStart, time.UTC)
}

func ParseTimestampRangeInLocation(queryParams url.Values, defaultStart *time.Time, loc *time.Location) (*TimestampRange, error) {
	start := queryParams.Get("start")
	if start == "" && defaultStart == nil {
		return nil, nil
	}

	currentTime := time.Now()
	timestampRange := &TimestampRange{
		Start: defaultStart,
		End:   &currentTime,
	}

	if start != "" {
		startTime, err := ParseTimestampInLocation(start, loc)
		if err != nil {
			return nil, err
		}

		// add one month after start
		endTime := startTime.AddDate(0, 1, 0)

		timestampRange.Start = startTime
		timestampRange.End = &endTime
	}

	if end := queryParams.Get("end"); end != "" {
		endTime, err := ParseTimestampInLocation(end, loc)
		if err != nil {
			return nil, err
		}

		// set endTime as end of day if time set as 00:00:00
		endTimeEndOfDay := *endTime
		if endTime.Hour() != 0 || endTime.Minute() == 0 && endTime.Second() == 0 {
			endTimeEndOfDay = now.With(endTimeEndOfDay).EndOfDay()
		}

		timestampRange.End = &endTimeEndOfDay
	}

	return timestampRange, nil
}

func ParseTimestamp(arg string) (*time.Time, error) {
	return ParseTimestampInLocation(arg, time.UTC)
}

func ParseTimestampInLocation(arg string, loc *time.Location) (*time.Time, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.000",
	}
	for _, f := range formats {
		t, err := time.ParseInLocation(f, arg, loc)
		if nil == err {
			return &t, nil
		}
	}
	return nil, errors.New("unsupported timestamp format")
}

func ParseUUID(arg string) uuid.UUID {
	if arg == "" {
		return uuid.UUID{}
	}
	return uuid.FromString(arg)
}
