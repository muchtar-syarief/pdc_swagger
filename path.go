package pdc_swagger

type ResponseObject struct {
	Description string                         `yaml:"description" json:"description"`
	Headers     map[string]interface{}         `yaml:"headers,omitempty" json:"headers,omitempty"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
	Links       map[string]interface{}         `yaml:"links,omitempty" json:"links,omitempty"`
}

type RequestBodyObject struct {
	Description string                         `yaml:"description" json:"description"`
	Required    bool                           `yaml:"required" json:"required"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
}

type ParameterObject struct {
	Name        string  `yaml:"name" json:"name"`
	In          string  `yaml:"in" json:"in"`
	Description string  `yaml:"description" json:"description"`
	Required    bool    `yaml:"required" json:"required"`
	Deprecated  bool    `yaml:"deprecated" json:"deprecated"`
	Schema      *Schema `yaml:"schema,omitempty" json:"schema,omitempty"`
}

type HTTPStatusCode string

type OperationObject struct {
	Tags        []string                           `yaml:"tags" json:"tags"`
	Summary     string                             `yaml:"summary" json:"summary"`
	Description string                             `yaml:"description" json:"description"`
	OperationId string                             `yaml:"operationId" json:"operationId"`
	Parameters  []*ParameterObject                 `yaml:"parameters" json:"parameters"`
	RequestBody *RequestBodyObject                 `yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	Responses   map[HTTPStatusCode]*ResponseObject `yaml:"responses,omitempty" json:"responses,omitempty"`
	Callbacks   interface{}                        `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	Deprecated  bool                               `yaml:"deprecated" json:"deprecated"`
	Security    interface{}                        `yaml:"security,omitempty" json:"security,omitempty"`
	Servers     []*ServerObject                    `yaml:"servers,omitempty" json:"servers,omitempty"`
}

type PathItemObject struct {
	Summary     string             `yaml:"summary" json:"summary"`
	Description string             `yaml:"description" json:"description"`
	Get         *OperationObject   `yaml:"get,omitempty" json:"get,omitempty"`
	Put         *OperationObject   `yaml:"put,omitempty" json:"put,omitempty"`
	Post        *OperationObject   `yaml:"post,omitempty" json:"post,omitempty"`
	Delete      *OperationObject   `yaml:"delete,omitempty" json:"delete,omitempty"`
	Options     *OperationObject   `yaml:"options,omitempty" json:"options,omitempty"`
	Head        *OperationObject   `yaml:"head,omitempty" json:"head,omitempty"`
	Patch       *OperationObject   `yaml:"patch,omitempty" json:"patch,omitempty"`
	Trace       *OperationObject   `yaml:"trace,omitempty" json:"trace,omitempty"`
	Servers     []*ServerObject    `yaml:"servers,omitempty" json:"servers,omitempty"`
	Parameters  []*ParameterObject `yaml:"parameters,omitempty" json:"parameters,omitempty"`
}

type Router string
type PdcSwaggerPath map[Router]*PathItemObject

func NewPath() PdcSwaggerPath {
	return PdcSwaggerPath{}
}

func (p PdcSwaggerPath) AddRouter(route Router, pathData *PathItemObject) PdcSwaggerPath {
	p[route] = pathData
	return p
}
