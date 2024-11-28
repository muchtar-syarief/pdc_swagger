package pdc_swagger

type ServerVariable struct {
	Enum        []string `yaml:"enum" json:"enum"`
	Default     string   `yaml:"default" json:"default"`
	Description string   `yaml:"description" json:"description"`
}

type ServerVariables map[string]*ServerVariable

type PdcSwaggerServer struct {
	Url         string          `yaml:"url" json:"url"`
	Description string          `yaml:"description" json:"description"`
	Variables   ServerVariables `yaml:"variables,omitempty" json:"variables,omitempty"`
}

func NewServer(url, desc string) *PdcSwaggerServer {
	return &PdcSwaggerServer{
		Url:         url,
		Description: desc,
	}
}

func (s *PdcSwaggerServer) SetVariables(name, desc, data string, enums []string) *PdcSwaggerServer {
	if s.Variables == nil {
		s.Variables = ServerVariables{}
	}

	variable := ServerVariable{
		Description: desc,
		Enum:        enums,
		Default:     data,
	}

	s.Variables[name] = &variable
	return s
}

type PdcSwaggerServers []*PdcSwaggerServer
