package helper

import (
	"reflect"
	"strings"
)

func GetFieldName(field reflect.StructField, tag string) string {
	value := field.Tag.Get(tag)
	if value == "" {
		value = field.Name
	} else {
		values := strings.Split(value, ",")
		value = values[0]
	}

	return value
}

func GetTagValue(field reflect.StructField, tag string) string {
	return field.Tag.Get(tag)
}

func GetTagValues(field reflect.StructField, tag string) []string {
	values := []string{}

	value := field.Tag.Get(tag)
	if value == "" {
		return values
	}

	return strings.Split(value, ",")
}

func IterateTagValues(field reflect.StructField, tag string, handler func(val string)) {
	values := GetTagValues(field, tag)
	for _, value := range values {
		handler(value)
	}
}
