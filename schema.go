package pdc_swagger

import (
	"errors"
	"reflect"
	"time"
)

type Schema struct {
	Type                 DataType           `yaml:"type" json:"type"`
	Properties           map[string]*Schema `yaml:"properties,omitempty" json:"properties,omitempty"`
	Items                *Schema            `yaml:"items,omitempty" json:"items,omitempty"`
	Format               string             `yaml:"format,omitempty" json:"format,omitempty"`
	AdditionalProperties *Schema            `yaml:"additionalProperties,omitempty" json:"additionalProperties,omitempty"`
}

func NewSchema(data interface{}) (*Schema, error) {
	var err error
	schema := &Schema{}

	dataType := reflect.TypeOf(data)
	kind := dataType.Kind()

	schema.Type = GetDataTypeMapper(kind)

	switch kind {
	case reflect.Struct:
		if schema.Properties == nil {
			schema.Properties = map[string]*Schema{}
		}

		switch data.(type) {
		case time.Time:
			schema = &Schema{
				Type: DataTypeString,
			}

			return schema, err
		}

		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)

			dataModel := reflect.Zero(field.Type).Interface()

			switch field.Type.Kind() {
			case reflect.Pointer:
				schemaProperties, err := NewSchema(dataModel)
				if err != nil {
					return schema, err
				}
				schema.Properties = schemaProperties.Properties

			default:
				nameField := field.Tag.Get("json")
				if nameField == "" {
					nameField = field.Name
				}

				schemaProperties, err := NewSchema(dataModel)
				if err != nil {
					return schema, err
				}
				schema.Properties[nameField] = schemaProperties

				formatField := field.Tag.Get("fmt")
				if formatField != "" {
					schema.Properties[nameField].Format = formatField
				}

			}
		}

	case reflect.Map:
		valueType := reflect.TypeOf(data).Elem()

		dataModel := reflect.Zero(valueType).Interface()
		additionalProperties, err := NewSchema(dataModel)
		if err != nil {
			return schema, err
		}

		schema.AdditionalProperties = additionalProperties

	case reflect.Pointer:
		dataModel := reflect.Zero(dataType.Elem()).Interface()
		result, err := NewSchema(dataModel)
		if err != nil {
			return schema, err
		}

		schema = result
		return schema, err

	case reflect.Invalid:
		err := errors.New("error unknown data type")
		return schema, err
	}

	return schema, err
}
