package types

import (
	"fmt"
	"reflect"
)

type PropertyFilter struct {
	PriceRange    Range  `db:"price"`
	Property_type string `db:"property_type"`
	Address       string `db:"address"`
	BedRange      Range  `db:"bedrooms"`
	UserId        int    `db:"user_id"`
}

type Range struct {
	Upper int
	Lower int
}

func NewRange(upper int, lower int) *Range {
	return &Range{Upper: upper, Lower: lower}
}

func NewPropertyFilter(property_type string, address string, priceRange Range, bedRange Range, userId int) *PropertyFilter {

	return &PropertyFilter{
		BedRange:      bedRange,
		Property_type: property_type,
		Address:       address,
		PriceRange:    priceRange,
		UserId:        userId,
	}
}

func (p PropertyFilter) GenerateQueryString() (string, []interface{}) {
	var variables []interface{}
	queryString := ""
	filterVal := reflect.ValueOf(p)
	filterType := filterVal.Type()
	var queryFilter string
	for i := 0; i < filterVal.NumField(); i++ {
		field := filterVal.Field(i)
		fieldType := filterType.Field(i) // Get tag
		dbColumn := fieldType.Tag.Get("db")
		if field.Type().Name() == "Range" {
			max := field.FieldByName("Upper").Interface().(int)
			min := field.FieldByName("Lower").Interface().(int)
			if min != 0 || max != 0 {
				queryFilter = fmt.Sprintf(" %s BETWEEN ? AND ? ", dbColumn)
				variables = append(variables, max, min)
				if queryString != "" {
					queryString = queryString + " AND " + queryFilter
				} else {
					queryString = queryFilter
				}

			}
		} else if field.Type().Kind() == reflect.String {
			if field.Interface().(string) != "" {
				queryFilter = fmt.Sprintf(" %s LIKE ? ", dbColumn)
				variables = append(variables, fmt.Sprintf("%%%s%%", field.Interface().(string)))
				if queryString != "" {
					queryString = queryString + " AND " + queryFilter
				} else {
					queryString = queryFilter
				}

			}
		} else {
			if field.Interface().(int) != 0 {
				queryFilter = fmt.Sprintf("%s=?", dbColumn)
				variables = append(variables, field.Interface())
			}
		}

	}

	return queryString, variables
}
