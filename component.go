package pdc_swagger

type Component struct {
	Schemas map[string]*Schema `yaml:"schemas" json:"schemas"`
}
type PdcGoDocumentationComponent struct {
	Components *Component `yaml:"components,omitempty" json:"components,omitempty"`
}

func NewComponent() *PdcGoDocumentationComponent {
	return &PdcGoDocumentationComponent{
		Components: &Component{
			Schemas: map[string]*Schema{},
		},
	}
}

func (c *PdcGoDocumentationComponent) AddComponent(name string, data interface{}) *PdcGoDocumentationComponent {
	schema := NewSchema(data)
	c.Components.Schemas[name] = schema

	return c
}
