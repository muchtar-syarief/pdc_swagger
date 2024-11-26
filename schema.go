package pdc_swagger

import (
	"errors"
	"reflect"
)

type Schema struct {
	Type       DataType `yaml:"type" json:"type"`
	Properties *Schema  `yaml:"properties,omitempty" json:"properties,omitempty"`
	Items      *Schema  `yaml:"items,omitempty" json:"items,omitempty"`
	Format     string   `yaml:"format,omitempty" json:"format,omitempty"`
}

func BuildSchema(data interface{}) (*Schema, error) {
	schema := &Schema{}

	dataType := reflect.TypeOf(data)
	schema.Type = GetDataTypeMapper(dataType.Kind())
	if schema.Type == DataTypeUnknown {
		err := errors.New("invalid data type")
		return schema, err
	}

	// switch schema.Type {
	// case DataTypeObject:
	// 	for _, item := range dataType.Elem() {
	// 	schema.Properties = BuildSchema(item)
	// 	}
	// }

	return schema, nil
}
