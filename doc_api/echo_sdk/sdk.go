package echo_sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/muchtar-syarief/pdc_swagger/doc_api"
	"gopkg.in/yaml.v3"
)

type EchoApiSdk struct {
	E   *echo.Echo
	doc doc_api.Documentation

	isApiDocRegistered bool
}

func NewEchoApiSdk(e *echo.Echo) *EchoApiSdk {
	return &EchoApiSdk{
		E: e,
	}
}

func (sdk *EchoApiSdk) GetGinEngine() *echo.Echo {
	return sdk.E
}

func (sdk *EchoApiSdk) Group(relativePath string) *SdkGroup {
	newGroup := SdkGroup{
		sdk:      sdk,
		G:        sdk.E.Group(relativePath),
		Basepath: relativePath,
	}

	return &newGroup
}

type RegisterFunc func(api pdc_swagger.Api, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *echo.Route

func (sdk *EchoApiSdk) RegisterGroup(relativePath string, groupHandler func(group *echo.Group, register RegisterFunc)) {
	e := sdk.E.Group(relativePath)

	var registfn RegisterFunc = func(api pdc_swagger.Api, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *echo.Route {
		api.SetGroupPath(relativePath)

		if sdk.doc != nil {
			sdk.doc.AddToDocumentation(api)
		}

		return e.Add(api.GetMethod(), api.GetRelativePath(), handler, middlewares...)
	}

	groupHandler(e, registfn)
}

func (sdk *EchoApiSdk) UseDocumentation(doc doc_api.Documentation) *EchoApiSdk {
	sdk.doc = doc
	return sdk
}

func (sdk *EchoApiSdk) UseSwaggerDocumentation(dataUri, docUri string) *EchoApiSdk {
	if sdk.doc == nil {
		return sdk
	}

	if !sdk.isApiDocRegistered {
		sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
			sdk.E.Add(method, path, func(ctx echo.Context) error {
				raw, err := yaml.Marshal(sdk.doc)
				if err != nil {
					return ctx.JSON(500, gin.H{"status": "error"})
				}

				return ctx.Blob(http.StatusOK, "application/x-yaml", raw)
			})
		})
	}

	sdk.doc.RegisterSwaggerDocumentation(dataUri, docUri, func(method, path string, responseTemplate func() (string, error)) {
		sdk.E.Add(method, path, func(ctx echo.Context) error {
			template, err := responseTemplate()
			if err != nil {
				return ctx.JSON(500, gin.H{"status": "error"})
			}

			return ctx.HTML(200, template)
		})
	})

	sdk.isApiDocRegistered = true

	return sdk
}

func (sdk *EchoApiSdk) UseRedocDocumentation(dataUri, docUri string) *EchoApiSdk {
	if sdk.doc == nil {
		return sdk
	}

	if !sdk.isApiDocRegistered {
		sdk.doc.RegisterDataDocumentation(dataUri, func(method, path string) {
			sdk.E.Add(method, path, func(ctx echo.Context) error {
				raw, err := yaml.Marshal(sdk.doc)
				if err != nil {
					return ctx.JSON(500, gin.H{"status": "error"})
				}

				return ctx.Blob(http.StatusOK, "application/x-yaml", raw)
			})
		})
	}

	sdk.doc.RegisterRedocDocumentation(dataUri, docUri, func(method, path string, responseTemplate func() (string, error)) {
		sdk.E.Add(method, path, func(ctx echo.Context) error {
			template, err := responseTemplate()
			if err != nil {
				return ctx.JSON(500, gin.H{"status": "error"})
			}

			return ctx.HTML(200, template)
		})
	})

	sdk.isApiDocRegistered = true

	return sdk
}

func (sdk *EchoApiSdk) Register(api pdc_swagger.Api, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *echo.Route {

	if sdk.doc != nil {
		sdk.doc.AddToDocumentation(api)
	}

	return sdk.E.Add(api.GetMethod(), api.GetRelativePath(), handler, middlewares...)
}
