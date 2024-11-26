package pdc_swagger

type PdcSwagger struct {
	OpenApi string            `yaml:"openapi" json:"openapi"`
	Info    *PdcSwaggerInfo   `yaml:"info" json:"info"`
	Servers PdcSwaggerServers `yaml:"servers" json:"servers"`
	// Paths
}

func NewPdcSwagger(title, description, version string) *PdcSwagger {
	info := NewPdcSwaggerInfo(title, description, version)

	return &PdcSwagger{
		OpenApi: "3.0.0",
		Info:    info,
		Servers: PdcSwaggerServers{},
	}
}

func (s *PdcSwagger) AddServer(server *PdcSwaggerServer) *PdcSwagger {
	s.Servers = append(s.Servers, server)

	return s
}
