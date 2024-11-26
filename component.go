package pdc_swagger

type ComponentName string

type ComponentSchema map[ComponentName]interface{}

type PdcSwaggerComponent struct {
	Schemas string `yaml:"schemas" json:"schemas"`
}
