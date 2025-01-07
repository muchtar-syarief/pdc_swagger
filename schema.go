package pdc_swagger

import (
	"reflect"
	"time"

	"github.com/muchtar-syarief/pdc_swagger/helper"
)

type Schema struct {
	Type                 DataType           `yaml:"type" json:"type"`
	Properties           map[string]*Schema `yaml:"properties,omitempty" json:"properties,omitempty"`
	Items                *Schema            `yaml:"items,omitempty" json:"items,omitempty"`
	Format               string             `yaml:"format,omitempty" json:"format,omitempty"`
	AdditionalProperties *Schema            `yaml:"additionalProperties,omitempty" json:"additionalProperties,omitempty"`
	Required             []string           `yaml:"required,omitempty" json:"required,omitempty"`

	// // type int
	// Minimum int `yaml:"minimum,omitempty" json:"minimum,omitempty"`
	// Maximum int `yaml:"maximum,omitempty" json:"maximum,omitempty"`

	// // type string
	// MinLength int `yaml:"minLength,omitempty" json:"minLength,omitempty"`
	// MaxLength int `yaml:"maxLength,omitempty" json:"maxLength,omitempty"`

	// // type array
	// MinItems int `yaml:"minItems,omitempty" json:"minItems,omitempty"`
	// MaxItems int `yaml:"maxItems,omitempty" json:"maxItems,omitempty"`
}

func NewSchema(data interface{}) *Schema {
	schema := &Schema{}

	dataType := reflect.TypeOf(data)
	kind := dataType.Kind()

	schemaType := GetDataTypeMapper(kind)
	if schemaType == DataTypeUnknown {
		return schema
	}

	schema.Type = schemaType

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

			return schema
		}

		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)

			dataModel := reflect.Zero(field.Type).Interface()

			switch field.Type.Kind() {
			case reflect.Pointer:
				schemaProperties := NewSchema(dataModel)
				for key, value := range schemaProperties.Properties {
					schema.Properties[key] = value
				}

			default:
				fieldName := helper.GetFieldName(field, "json")
				schemaProperties := NewSchema(dataModel)

				schema.Properties[fieldName] = schemaProperties
			}
		}

	case reflect.Map:
		valueType := dataType.Elem()

		dataModel := reflect.Zero(valueType).Interface()

		additionalProperties := NewSchema(dataModel)
		schema.AdditionalProperties = additionalProperties

	case reflect.Pointer:
		dataModel := reflect.Zero(dataType.Elem()).Interface()
		result := NewSchema(dataModel)

		schema = result
		return schema

	case reflect.Array, reflect.Slice:
		valueType := dataType.Elem()

		dataModel := reflect.Zero(valueType).Interface()
		properties := NewSchema(dataModel)
		schema.Items = properties

	case reflect.Invalid:
		return schema
	}

	return schema
}
