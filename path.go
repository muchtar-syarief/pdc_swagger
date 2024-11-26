package pdc_swagger

type ResponseContentItem struct {
	Schema interface{} `yaml:"schema" json:"schema"`
}

type MediaType string
type ResponseContent map[MediaType]ResponseContentItem

type ItemResponse struct {
	Description string          `yaml:"description" json:"description"`
	Content     ResponseContent `yaml:"content" json:"content"`
}

type HTTPStatusCode int
type PathItemResponses map[HTTPStatusCode]interface{}

type PathItemParameters struct {
	Name        string      `yaml:"name" json:"name"`
	In          string      `yaml:"in" json:"in"`
	Description string      `yaml:"description" json:"description"`
	Required    bool        `yaml:"required" json:"required"`
	Schema      interface{} `yaml:"schema" json:"schema"`
}

type PdcSwaggerPathItem struct {
	Summary     string             `yaml:"summary" json:"summary"`
	Description string             `yaml:"description" json:"description"`
	Parameters  PathItemParameters `yaml:"parameters" json:"parameters"`
	Responses   PathItemResponses  `yaml:"responses" json:"responses"`
}

type HTTPMethod string
type PdcSwaggerPathData map[HTTPMethod]*PdcSwaggerPathItem

func NewPathData() PdcSwaggerPathData {
	return PdcSwaggerPathData{}
}

func (p PdcSwaggerPathData) AddPathData(method HTTPMethod, pathData *PdcSwaggerPathItem) PdcSwaggerPathData {
	p[method] = pathData
	return p
}

type Router string
type PdcSwaggerPath map[Router]PdcSwaggerPathData

func NewPath() PdcSwaggerPath {
	return PdcSwaggerPath{}
}

func (p PdcSwaggerPath) AddRouter(route Router, pathData PdcSwaggerPathData) PdcSwaggerPath {
	p[route] = pathData
	return p
}
