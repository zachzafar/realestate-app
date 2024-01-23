package types

import (
	"fmt"
	"reflect"
)

type PropertyFilter struct {
	PriceRange    Range `db:"price"`
	Property_type int   `db:"property_type_id"`
	Country       int   `db:"country_id"`
	Parish        int   `db:"parish_id"`
	Beds          int   `db:"bedrooms"`
	Bathrooms     int   `db:"bathrooms"`
	UserId        int   `db:"user_id"`
}

type Range struct {
	Upper int
	Lower int
}

func NewRange(upper int, lower int) *Range {
	return &Range{Upper: upper, Lower: lower}
}

func NewPropertyFilter(property_type int, country int, parish int, priceRange Range, beds int, bathrooms int, userId int) *PropertyFilter {

	return &PropertyFilter{
		Beds:          beds,
		Bathrooms:     bathrooms,
		Property_type: property_type,
		Country:       country,
		Parish:        parish,
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
		} else {
			if field.Interface().(int) != 0 {
				counter = counter + 1
				queryFilter = fmt.Sprintf("%s=$%d", dbColumn, counter)
				variables = append(variables, field.Interface())
			}
		}

	}

	return queryString, variables
}
