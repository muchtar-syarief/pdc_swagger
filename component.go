package pdc_swagger

type Component struct {
	Schemas map[string]*Schema `yaml:"schemas" json:"schemas"`
}
type PdcSwaggerComponent struct {
	Components *Component `yaml:"components,omitempty" json:"components,omitempty"`
}

func NewComponent() *PdcSwaggerComponent {
	return &PdcSwaggerComponent{
		Components: &Component{
			Schemas: map[string]*Schema{},
		},
	}
}

func (c *PdcSwaggerComponent) AddComponent(name string, data interface{}) *PdcSwaggerComponent {
	schema := NewSchema(data)
	c.Components.Schemas[name] = schema

	return c
}
