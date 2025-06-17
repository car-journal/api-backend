package database

import (
	"fmt"
	"strings"

	"github.com/car-journal/config"
	cardvaultstrings "github.com/car-journal/lib/strings"
	"gorm.io/gorm"
)

func GenerateAffectedRecords() string {
	return "COUNT (*) OVER() AS affected_records"
}

func GeneratePaginationQuery(db *gorm.DB, limit int, offset int) *gorm.DB {
	return db.Limit(limit).Offset(offset)
}

func GenerateOrderQuery(db *gorm.DB, keys []string, validSorts map[string]string) *gorm.DB {
	uniqueFields := make(map[string]bool)
	for _, key := range keys {
		keyParts := cardvaultstrings.Split(key, ":")
		if len(keyParts) != 2 {
			continue
		}

		currKey, direction := strings.ToLower(keyParts[0]), strings.ToUpper(keyParts[1])
		field, okField := validSorts[currKey]
		if (!okField || (direction != config.ASCENDING && direction != config.DESCENDING)) && currKey != config.RANDOM {
			continue
		}

		_, okUniqueFields := uniqueFields[field]
		if okUniqueFields {
			continue
		}

		var orderQuery string
		switch {
		case direction == "": // no direction provided
			orderQuery = field
		case strings.Contains(field, config.NULLS_LAST): // sort withs nulls last
			field = strings.TrimSpace(strings.ReplaceAll(field, config.NULLS_LAST, ""))
			orderQuery = fmt.Sprintf("%s %s %s", field, direction, config.NULLS_LAST)

		default: // typical sort query
			orderQuery = fmt.Sprintf("%s %s", field, direction)
		}

		db = db.Order(orderQuery)
		uniqueFields[field] = true
	}

	return db
}

func CountAffectedRecords(db *gorm.DB, tableName string) *gorm.DB {
	return db.Select(fmt.Sprintf("%s.*, COUNT (*) Over() AS affected_records", tableName))
}

func GenerateJoinQuery(db *gorm.DB, joinType string, rightTableName string, rightTableAlias string, rightTableKey string, leftTableAlias string, leftTableKey string) *gorm.DB {
	return db.Joins(fmt.Sprintf(`%s JOIN %s AS %s ON %s."%s" = %s."%s"`, joinType, rightTableName, rightTableAlias, leftTableAlias, leftTableKey, rightTableAlias, rightTableKey))
}

func EqualsTo(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", columnName), value)
}

func AndOr(db *gorm.DB, andColumnName string, andValue interface{}, orColumnName string, orValue interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", andColumnName), andValue).Or(fmt.Sprintf("%s = ?", orColumnName), orValue)
}

func UnequalTo(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s <> ?", columnName), value)
}

func LowerThanAndEqualsTo(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s <= ?", columnName), value)
}

func GreaterThanAndEqualsTo(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s >= ?", columnName), value)
}

func GreaterThan(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s > ?", columnName), value)
}

func In(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s IN ?", columnName), value)
}

func ILike(db *gorm.DB, columnName string, value interface{}) *gorm.DB {
	return db.Where(fmt.Sprintf("%s ILIKE ?", columnName), fmt.Sprintf("%%%s%%", value))
}

// func CreateBatchSession(db *gorm.DB) *gorm.DB {
// 	return db.Session(&gorm.Session{CreateBatchSize: helper.MustStringToInt(config.Get(config.POSTGRES_CREATE_BATCH_SIZE))})
// }

// func GenerateILikePattern(values []string) string {
// 	var patterns []string
// 	for _, value := range values {
// 		patterns = append(patterns, "%"+value+"%")
// 	}
// 	return "{" + strings.Join(patterns, ",") + "}"
// }
