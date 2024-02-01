package types

import (
	"fmt"
	"reflect"
	"strings"
)

type ContextKey string

type Range struct {
	Upper int
	Lower int
}

func NewRange(upper int, lower int) *Range {
	return &Range{Upper: upper, Lower: lower}
}

type Relation interface {
	GetRelationName() string
	GetPrimaryKeyName() string
}

func GeneratInsertQuery(r Relation) (string, []interface{}) {
	var variables []interface{}
	var columns []string
	var placeholders []string
	StructVal := reflect.Indirect(reflect.ValueOf(r))
	StructType := StructVal.Type()
	counter := 0

	for i := 0; i < StructVal.NumField(); i++ {
		counter += 1
		field := StructVal.Field(i)
		fieldType := StructType.Field(i)
		columns = append(columns, fieldType.Tag.Get("db"))
		variables = append(variables, field.Interface())
		placeholders = append(placeholders, fmt.Sprintf("$%d", counter))
	}

	columnString := "(" + strings.Join(columns, ",") + ")"
	placeholderString := "(" + strings.Join(placeholders, ",") + ")"
	insertStatement := fmt.Sprintf("INSERT INTO %s ", r.GetRelationName())
	query := insertStatement + columnString + " VALUES " + placeholderString

	return query, variables
}

func GenerateQueryByIDString(r Relation) (string, []interface{}) {
	var pointers []interface{}
	var columns []string
	StructVal := reflect.Indirect(reflect.ValueOf(r))
	StructType := StructVal.Type()

	for i := 0; i < StructVal.NumField(); i++ {
		fieldValue := StructVal.Field(i)
		pointer := fieldValue.Addr().Interface()
		pointers = append(pointers, pointer)

		fieldType := StructType.Field(i)
		columns = append(columns, fieldType.Tag.Get("db"))
	}

	query := strings.Join(columns, ",")
	query = "SELECT " + query + " FROM " + r.GetRelationName() + " WHERE " + r.GetPrimaryKeyName() + "=$1"

	return query, pointers
}

type Filter interface {
	GetRelationForFilterName() string
}

func GenerateFilterQueryString(f Filter) (string, []interface{}) {
	var variables []interface{}
	queryString := ""
	filterVal := reflect.ValueOf(f)
	filterType := filterVal.Type()
	var queryFilter string
	counter := 0
	for i := 0; i < filterVal.NumField(); i++ {
		field := filterVal.Field(i)
		fieldType := filterType.Field(i) // Get tag
		dbColumn := fieldType.Tag.Get("db")
		if field.Type().Name() == "Range" {
			max := field.FieldByName("Upper").Interface().(int)
			min := field.FieldByName("Lower").Interface().(int)
			if min != 0 || max != 0 {
				max_ph := "$" + fmt.Sprint(counter+1)
				min_ph := "$" + fmt.Sprint(counter+2)
				queryFilter = fmt.Sprintf(" %s BETWEEN %s AND %s ", dbColumn, max_ph, min_ph)
				counter = counter + 2
				variables = append(variables, max, min)
				if queryString != "" {
					queryString = queryString + " AND " + queryFilter
				} else {
					queryString = queryFilter
				}

			}
		} else if field.Type().Kind() == reflect.String {
			if field.Interface().(string) != "" {
				counter = counter + 1
				queryFilter = fmt.Sprintf(" %s LIKE $%d ", dbColumn, counter)
				variables = append(variables, fmt.Sprintf("%%%s%%", field.Interface().(string)))
				if queryString != "" {
					queryString = queryString + " AND " + queryFilter
				} else {
					queryString = queryFilter
				}

			}
		} else if field.Type().Kind() == reflect.Int {
			if field.Interface().(int) != 0 {

				counter = counter + 1
				queryFilter = fmt.Sprintf("%s=$%d", dbColumn, counter)
				variables = append(variables, field.Interface())
				if queryString != "" {
					queryString = queryString + " AND " + queryFilter
				} else {
					queryString = queryFilter
				}
			}
		}

	}

	return queryString, variables
}

// func (r Relation) GenerateSelectionQuery() (string, []interface{}) {

// }
