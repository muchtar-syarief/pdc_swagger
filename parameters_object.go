package pdc_swagger

import "reflect"

type ParameterObject struct {
	Name        string  `yaml:"name" json:"name"`
	In          string  `yaml:"in" json:"in"`
	Description string  `yaml:"description,omitempty" json:"description,omitempty"`
	Required    bool    `yaml:"required,omitempty" json:"required,omitempty"`
	Deprecated  bool    `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	Schema      *Schema `yaml:"schema,omitempty" json:"schema,omitempty"`
}

func NewListParametersObject(data interface{}) []*ParameterObject {
	results := []*ParameterObject{}

	dataType := reflect.TypeOf(data)
	kind := dataType.Kind()

	switch kind {
	case reflect.Struct:
		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)

			dataModel := reflect.Zero(field.Type).Interface()

			switch field.Type.Kind() {
			case reflect.Pointer:
				schema := NewSchema(dataModel)
				for key, item := range schema.Properties {
					results = append(results, &ParameterObject{
						Name:   key,
						In:     "query",
						Schema: item,
					})
				}

			default:
				nameField := field.Tag.Get("json")
				if nameField == "" {
					nameField = field.Name
				}

				results = append(results, &ParameterObject{
					Name:   nameField,
					In:     "query",
					Schema: NewSchema(dataModel),
				})
			}
		}
	}

	return results
}
