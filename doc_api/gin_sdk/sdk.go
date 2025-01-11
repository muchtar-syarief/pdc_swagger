package gin_sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/muchtar-syarief/pdc_swagger/doc_api"
)

type GinApiSdk struct {
	R   *gin.Engine
	doc doc_api.Documentation

	isApiDocRegistered bool
}

func NewGinApiSdk(r *gin.Engine) *GinApiSdk {
	sdk := &GinApiSdk{
		R: r,
	}
	return sdk
}

func (sdk *GinApiSdk) UseDocumentation(doc doc_api.Documentation) *GinApiSdk {
	sdk.doc = doc
	return sdk
}

func (sdk *GinApiSdk) GetGinEngine() *gin.Engine {
	return sdk.R
}

func (sdk *GinApiSdk) UseSwaggerDocumentation(dataUri, docUri string) *GinApiSdk {
	if !sdk.isApiDocRegistered {
		sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
			sdk.R.Handle(method, path, func(ctx *gin.Context) {
				ctx.YAML(200, sdk.doc)
			})
		})
	}

	sdk.doc.RegisterSwaggerDocumentation(dataUri, docUri, func(method, path string, responseTemplate func() (string, error)) {
		sdk.R.Handle(method, path, func(ctx *gin.Context) {
			template, err := responseTemplate()
			if err != nil {
				ctx.JSON(500, gin.H{"status": "error"})
				return
			}

			ctx.Data(200, "text/html", []byte(template))
		})
	})

	sdk.isApiDocRegistered = true

	return sdk
}

func (sdk *GinApiSdk) UseRedocDocumentation(dataUri, docUri string) *GinApiSdk {
	if !sdk.isApiDocRegistered {
		sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
			sdk.R.Handle(method, path, func(ctx *gin.Context) {
				ctx.YAML(200, sdk.doc)
			})
		})
	}

	sdk.doc.RegisterRedocDocumentation(dataUri, docUri, func(method, path string, responseTemplate func() (string, error)) {
		sdk.R.Handle(method, path, func(ctx *gin.Context) {
			template, err := responseTemplate()
			if err != nil {
				ctx.JSON(500, gin.H{"status": "error"})
				return
			}

			ctx.Data(200, "text/html", []byte(template))
		})
	})

	sdk.isApiDocRegistered = true

	return sdk
}

func (sdk *GinApiSdk) Register(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes {

	if sdk.doc != nil {
		sdk.doc.AddToDocumentation(api)
	}

	return sdk.R.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
}
