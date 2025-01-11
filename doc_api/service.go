package doc_api

import "github.com/muchtar-syarief/pdc_swagger"

type Documentation interface {
	AddToDocumentation(api pdc_swagger.Api)
	RegisterDataDocumentation(url string, handler func(method, path string))
	RegisterSwaggerDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error)))
	RegisterRedocDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error)))
}
