package pdc_swagger

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/muchtar-syarief/pdc_swagger/view"
	"gopkg.in/yaml.v3"
)

type PdcOpenApi struct {
	OpenApi string                     `yaml:"openapi" json:"openapi"`
	Info    *Info                      `yaml:"info" json:"info"`
	Servers []*ServerObject            `yaml:"servers,omitempty" json:"servers,omitempty"`
	Paths   map[string]*PathItemObject `yaml:"paths,omitempty" json:"paths,omitempty"`
}

func NewPdcOpenApi(title, description, apiVersion string) *PdcOpenApi {
	info := NewInfo(title, description, apiVersion)

	return &PdcOpenApi{
		OpenApi: "3.0.0",
		Info:    info,
	}
}

func (d *PdcOpenApi) addApi(fullPathUri string, item *PathItemObject) *PdcOpenApi {
	if d.Paths == nil {
		d.Paths = map[string]*PathItemObject{}
	}

	if !strings.HasPrefix(fullPathUri, "/") {
		fullPathUri = "/" + fullPathUri
	}

	d.Paths[fullPathUri] = item

	return d
}

type Api interface {
	GetFullUriPath() string
	GetTags() []string
	GetSummary() string
	GetDescription() string
	GetKeyName() string
	GetQuery() any
	GetPayload() any
	GetResponse() any
	GetMethod() string
	GetGroupPath() string
	GetRelativePath() string

	SetGroupPath(path string)
}

func (d *PdcOpenApi) AddToDocumentation(api Api) {
	operationObject := NewOperationObject(
		api.GetTags(),
		api.GetSummary(),
		api.GetDescription(),
		api.GetKeyName(),
	)

	operationObject.
		SetParameters(api.GetQuery()).
		SetRequestBody(api.GetPayload()).
		SetResponse("200", api.GetResponse())

	pathItem := NewPathItemObjectDefault()
	pathItem.
		SetOperationObject(
			api.GetMethod(),
			operationObject,
		)

	d.addApi(api.GetFullUriPath(), pathItem)
}

func (d *PdcOpenApi) RegisterDataDocumentation(url string, handler func(method, path string)) {
	if url == "" {
		url = "/doc_data"
	}

	handler(http.MethodGet, url)
}

func (d *PdcOpenApi) RegisterSwaggerDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error))) {
	if urlData == "" {
		urlData = "/doc_data"
	}

	if urlDoc == "" {
		urlDoc = "/docs"
	}

	getTemplate := func() (string, error) {
		return view.GetSwaggerViewTemplate(&view.ViewTemplateConfig{
			Title: d.Info.Title,
			Url:   urlData,
		})
	}

	handler(http.MethodGet, urlDoc, getTemplate)
}

func (d *PdcOpenApi) RegisterRedocDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error))) {
	if urlData == "" {
		urlData = "/doc_data"
	}

	if urlDoc == "" {
		urlDoc = "/redoc"
	}

	getTemplate := func() (string, error) {
		return view.GetRedocViewTemplate(&view.ViewTemplateConfig{
			Title: d.Info.Title,
			Url:   urlData,
		})
	}

	handler(http.MethodGet, urlDoc, getTemplate)
}

func (d *PdcOpenApi) Save(filename string) error {
	var err error

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		return err
	}

	var data []byte

	ext := filepath.Ext(filename)
	switch ext {
	case "yaml", "yml":
		raw, err := yaml.Marshal(d)
		if err != nil {
			return err
		}

		data = raw

	default:
		raw, err := json.MarshalIndent(d, "", "	")
		if err != nil {
			return err
		}

		data = raw
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return err
}
