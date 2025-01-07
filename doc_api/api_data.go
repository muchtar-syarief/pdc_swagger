package doc_api

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

type ApiData struct {
	Method       string
	RelativePath string
	Payload      interface{}
	Response     interface{}
	Query        interface{}
	GroupPath    string
	Deprecated   bool
	Tags         []string
	Summary      string
	Description  string
}

// GetDescription implements pdc_swagger.Api.
func (a *ApiData) GetDescription() string {
	return a.Description
}

// GetGroupPath implements pdc_swagger.Api.
func (a *ApiData) GetGroupPath() string {
	return a.GroupPath
}

// GetMethod implements pdc_swagger.Api.
func (a *ApiData) GetMethod() string {
	return a.Method
}

// GetPayload implements pdc_swagger.Api.
func (a *ApiData) GetPayload() any {
	return a.Payload
}

// GetQuery implements pdc_swagger.Api.
func (a *ApiData) GetQuery() any {
	return a.Query
}

// GetRelativePath implements pdc_swagger.Api.
func (a *ApiData) GetRelativePath() string {
	return a.RelativePath
}

// GetResponse implements pdc_swagger.Api.
func (a *ApiData) GetResponse() any {
	return a.Response
}

// GetSummary implements pdc_swagger.Api.
func (a *ApiData) GetSummary() string {
	return a.Summary
}

// GetTags implements pdc_swagger.Api.
func (a *ApiData) GetTags() []string {
	return a.Tags
}

// SetGroupPath implements pdc_swagger.Api.
func (a *ApiData) SetGroupPath(path string) {
	a.GroupPath = path
}

func (api *ApiData) GetFullUriPath() string {
	name, _ := url.JoinPath(api.GroupPath, api.RelativePath)
	return name
}

func (api *ApiData) GetKeyName() string {
	name, _ := url.JoinPath(api.GroupPath, api.RelativePath)
	name = strings.TrimPrefix(name, `/`)

	fnname := ""
	funcname := filepath.Join(strings.ToLower(api.Method), name)
	funcs := strings.Split(funcname, `\`)
	for _, u := range funcs {
		fnname += strcase.ToCamel(u)
	}

	return fnname
}
