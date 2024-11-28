package pdc_swagger

type MediaType string

const (
	MediaTypeJson      MediaType = "application/json"
	MediaTypeTextPlain MediaType = "text/plain; charset=utf-8"
)

type MediaTypeObject struct {
	Schema *Schema `yaml:"schema" json:"schema"`
}
