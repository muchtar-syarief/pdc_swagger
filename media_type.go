package pdc_swagger

type MediaType string

const (
	MediaTypeJson      MediaType = "application/json"
	MediaTypeTextPlain MediaType = "text/plain; charset=utf-8"
)

type MediaTypeContent struct {
	Schema *Schema `yaml:"schema" json:"schema"`
}

type MediaTypeObject map[MediaType]*MediaTypeContent
