package doc_api

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/muchtar-syarief/pdc_swagger"
)

type ApiSdk struct {
	R   *gin.Engine
	doc Documentation
}

type Documentation interface {
	AddToDocumentation(api pdc_swagger.Api)
	RegisterDataDocumentation(url string, handler func(method, path string))
	RegisterSwaggerDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error)))
	RegisterRedocDocumentation(urlData, urlDoc string, handler func(method, path string, responseTemplate func() (string, error)))
}

func (sdk *ApiSdk) UseDocumentation(doc Documentation) *ApiSdk {
	sdk.doc = doc
	return sdk
}

func (sdk *ApiSdk) UseSwaggerDocumentation(dataUri, docUri string) *ApiSdk {
	sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
		sdk.R.Handle(method, path, func(ctx *gin.Context) {
			ctx.YAML(200, sdk.doc)
		})
	})

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

	return sdk
}

func (sdk *ApiSdk) UseRedocDocumentation(dataUri, docUri string) *ApiSdk {
	sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
		sdk.R.Handle(method, path, func(ctx *gin.Context) {
			ctx.YAML(200, sdk.doc)
		})
	})

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

	return sdk
}

type RegisterFunc func(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes

func (sdk *ApiSdk) Register(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes {

	if sdk.doc != nil {
		sdk.doc.AddToDocumentation(api)
	}

	return sdk.R.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
}

type SdkGroup struct {
	sdk      *ApiSdk
	G        *gin.RouterGroup
	Basepath string
}

func (grp *SdkGroup) Register(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes {
	api.SetGroupPath(grp.Basepath)

	if grp.sdk.doc != nil {
		grp.sdk.doc.AddToDocumentation(api)
	}

	return grp.G.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
}

func (grp *SdkGroup) Group(path string) *SdkGroup {
	base, _ := url.JoinPath(grp.Basepath, path)
	newGroup := SdkGroup{
		sdk:      grp.sdk,
		G:        grp.G.Group(path),
		Basepath: base,
	}

	return &newGroup
}

func (sdk *ApiSdk) Group(relativePath string) *SdkGroup {
	newGroup := SdkGroup{
		sdk:      sdk,
		G:        sdk.R.Group(relativePath),
		Basepath: relativePath,
	}

	return &newGroup
}

func (sdk *ApiSdk) RegisterGroup(relativePath string, groupHandler func(group *gin.RouterGroup, register RegisterFunc)) {
	r := sdk.R.Group(relativePath)

	var registfn RegisterFunc = func(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes {
		api.SetGroupPath(relativePath)

		if sdk.doc != nil {
			sdk.doc.AddToDocumentation(api)
		}

		return r.Handle(api.GetMethod(), api.GetRelativePath(), handlers...)
	}

	groupHandler(r, registfn)
}

func NewApiSdk(r *gin.Engine) *ApiSdk {
	sdk := &ApiSdk{
		R: r,
	}
	return sdk
}
