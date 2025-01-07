package pdc_swagger

import "net/http"

type HTTPStatusCode string

type PathItemObject struct {
	Summary     string             `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string             `yaml:"description,omitempty" json:"description,omitempty"`
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

func NewPathItemObjectDefault() *PathItemObject {
	return &PathItemObject{}
}

func NewPathItemObject(summary, desc string) *PathItemObject {
	return &PathItemObject{
		Summary:     summary,
		Description: desc,
	}
}

func (i *PathItemObject) SetOperationObject(method string, operation *OperationObject) *PathItemObject {
	switch method {
	case http.MethodGet:
		i.Get = operation
	case http.MethodPost:
		i.Post = operation
	case http.MethodPut:
		i.Put = operation
	case http.MethodDelete:
		i.Put = operation
	}

	return i
}

func (i *PathItemObject) SetParameters(data interface{}) *PathItemObject {
	if i.Parameters == nil {
		i.Parameters = []*ParameterObject{}
	}

	parameters := NewListParametersObject(data)
	i.Parameters = append(i.Parameters, parameters...)

	return i
}
