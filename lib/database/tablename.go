package database

import (
	"log"
	"sync"

	"gorm.io/gorm/schema"
)

func GetColumnNames(model interface{}) (res map[string]string, err error) {
	s, err := getSchema(model)
	if nil != err {
		return
	}

	res = make(map[string]string)
	for _, field := range s.Fields {
		columnName := field.Tag.Get("json")
		if columnName == "" || columnName == "-" {
			columnName = field.DBName
		}
		fieldName := field.Name
		res[fieldName] = columnName
	}

	return
}

func GetDBTableName(model interface{}) (tableName string, err error) {
	s, err := getSchema(model)
	if nil != err {
		return
	}
	return s.Table, nil
}

func GetModelName(model interface{}) (tableName string, err error) {
	s, err := getSchema(model)
	if nil != err {
		return
	}
	return s.Name, nil
}

func getSchema(model interface{}) (s *schema.Schema, err error) {
	s, err = schema.Parse(&model, &sync.Map{}, schema.NamingStrategy{})
	if nil != err {
		log.Print("failed to create scema")
		log.Print(err)
		return
	}
	return
}
