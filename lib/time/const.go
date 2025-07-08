package internaltime

// Go's reference layout
const (
	DayMonthYearHourMinuteSecondWithSlashLayout = "02/01/2006 15:04:05" // "DD/MM/YYYY HH:mm:ss"
	DayMonthYearWithSlashLayout                 = "02/01/2006"          // "DD/MM/YYYY"
	DayMonthYearWithDashLayout                  = "02-01-2006"          // "DD-MM-YYYY"

	YearMonthDayHourMinuteSecondWithDashLayout = "2006-01-02 15:04:05" // "YYYY-MM-DD HH:mm:ss"
)
