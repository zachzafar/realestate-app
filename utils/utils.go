package utils

import (
	"reflect"
)

func ObjContainsZeroOrEmptyStrings(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Struct {

		if val.IsZero() {
			return true
		}
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Struct {

			return ObjContainsZeroOrEmptyStrings(field.Interface())
		}

		if !field.IsZero() {
			return false
		}

	}

	return true
}

func IsLast(i int, list interface{}) bool {
	if reflect.TypeOf(list).Kind() != reflect.Slice {
		return false
	}

	length := reflect.ValueOf(list).Len()
	if length-1 == i {
		return true
	}

	return false
}
