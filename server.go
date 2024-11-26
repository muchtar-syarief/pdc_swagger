package pdc_swagger

type ServerVariable struct {
	Enum        []string `yaml:"enum" json:"enum"`
	Default     string   `yaml:"default" json:"default"`
	Description string   `yaml:"description" json:"description"`
}

type ServerVariables map[string]ServerVariable

type PdcSwaggerServer struct {
	Url         string          `yaml:"url" json:"url"`
	Description string          `yaml:"description" json:"description"`
	Variables   ServerVariables `yaml:"variables,omitempty" json:"variables,omitempty"`
}

type PdcSwaggerServers []*PdcSwaggerServer
