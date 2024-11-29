package pdc_swagger

import (
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

func NewSchema(data interface{}) *Schema {
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

			return schema
		}

		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)

			dataModel := reflect.Zero(field.Type).Interface()

			switch field.Type.Kind() {
			case reflect.Pointer:
				schemaProperties := NewSchema(dataModel)
				schema.Properties = schemaProperties.Properties

			default:
				nameField := field.Tag.Get("json")
				if nameField == "" {
					nameField = field.Name
				}

				schemaProperties := NewSchema(dataModel)
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

		additionalProperties := NewSchema(dataModel)
		schema.AdditionalProperties = additionalProperties

	case reflect.Pointer:
		dataModel := reflect.Zero(dataType.Elem()).Interface()
		result := NewSchema(dataModel)

		schema = result
		return schema

	case reflect.Invalid:
		return schema
	}

	return schema
}
