package pdc_swagger

import (
	"net/http"
	"strings"

	"github.com/muchtar-syarief/pdc_swagger/view"
)

type PdcSwagger struct {
	OpenApi string                     `yaml:"openapi" json:"openapi"`
	Info    *Info                      `yaml:"info" json:"info"`
	Servers []*ServerObject            `yaml:"servers,omitempty" json:"servers,omitempty"`
	Paths   map[string]*PathItemObject `yaml:"paths,omitempty" json:"paths,omitempty"`
}

func NewPdcSwagger(title, description, version string) *PdcSwagger {
	info := NewInfo(title, description, version)

	return &PdcSwagger{
		OpenApi: "3.0.0",
		Info:    info,
	}
}

func (s *PdcSwagger) AddServer(server *ServerObject) *PdcSwagger {
	if s.Servers == nil {
		s.Servers = []*ServerObject{}
	}

	s.Servers = append(s.Servers, server)

	return s
}

func (s *PdcSwagger) AddApi(fullPathUri string, item *PathItemObject) *PdcSwagger {
	if s.Paths == nil {
		s.Paths = map[string]*PathItemObject{}
	}

	if !strings.HasPrefix(fullPathUri, "/") {
		fullPathUri = "/" + fullPathUri
	}

	s.Paths[fullPathUri] = item

	return s
}

func (s *PdcSwagger) RegisterDataDocumentation(url string, handler func(method, path string)) *PdcSwagger {
	if url == "" {
		url = "/doc_data"
	}

	handler(http.MethodGet, url)

	return s
}

func (s *PdcSwagger) RegisterSwaggerDocumentation(url string, handler func(method, path, responseTemplate string)) *PdcSwagger {
	if url == "" {
		url = "/doc_data"
	}

	template, err := view.GetSwaggerViewTemplate(&view.ViewTemplateConfig{
		Title: s.Info.Title,
		Url:   url,
	})
	if err != nil {
		handler(http.MethodGet, "/docs", "error register swagger documentation")
		return s
	}

	handler(http.MethodGet, "/docs", template)

	return s
}

func (s *PdcSwagger) RegisterRedocDocumentation(url string, handler func(method, path, responseTemplate string)) *PdcSwagger {
	if url == "" {
		url = "/doc_data"
	}

	template, err := view.GetSwaggerViewTemplate(&view.ViewTemplateConfig{
		Title: s.Info.Title,
		Url:   url,
	})
	if err != nil {
		handler(http.MethodGet, "/docs", "error register redoc documentation")
		return s
	}

	handler(http.MethodGet, "/redoc", template)

	return s
}
