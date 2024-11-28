package pdc_swagger

type ServerVariable struct {
	Enum        []string `yaml:"enum" json:"enum"`
	Default     string   `yaml:"default" json:"default"`
	Description string   `yaml:"description" json:"description"`
}

type ServerObject struct {
	Url         string                     `yaml:"url" json:"url"`
	Description string                     `yaml:"description" json:"description"`
	Variables   map[string]*ServerVariable `yaml:"variables,omitempty" json:"variables,omitempty"`
}

func NewServer(url, desc string) *ServerObject {
	return &ServerObject{
		Url:         url,
		Description: desc,
	}
}

func (s *ServerObject) SetVariables(name, desc, data string, enums []string) *ServerObject {
	if s.Variables == nil {
		s.Variables = map[string]*ServerVariable{}
	}

	variable := ServerVariable{
		Description: desc,
		Enum:        enums,
		Default:     data,
	}

	s.Variables[name] = &variable
	return s
}
