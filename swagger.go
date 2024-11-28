package pdc_swagger

type PdcSwagger struct {
	OpenApi string                     `yaml:"openapi" json:"openapi"`
	Info    *Info                      `yaml:"info" json:"info"`
	Servers []*ServerObject            `yaml:"servers,omitempty" json:"servers,omitempty"`
	Paths   map[string]*PathItemObject `yaml:"paths" json:"paths"`
}

func NewPdcSwagger(title, description, version string) *PdcSwagger {
	info := NewInfo(title, description, version)

	return &PdcSwagger{
		OpenApi: "3.0.0",
		Info:    info,
		Servers: []*ServerObject{},
	}
}

func (s *PdcSwagger) AddServer(server *ServerObject) *PdcSwagger {
	s.Servers = append(s.Servers, server)

	return s
}
