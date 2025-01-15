package echo_sdk

import (
	"github.com/labstack/echo/v4"
	"github.com/muchtar-syarief/pdc_swagger"
)

type EchoApiSdk struct {
	E *echo.Echo

	middlewares []func(pdc_swagger.Api)
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

		for _, middlewareFunc := range sdk.middlewares {
			middlewareFunc(api)
		}

		return e.Add(api.GetMethod(), api.GetRelativePath(), handler, middlewares...)
	}

	groupHandler(e, registfn)
}

func (sdk *EchoApiSdk) Use(handler func(pdc_swagger.Api)) *EchoApiSdk {
	sdk.middlewares = append(sdk.middlewares, handler)
	return sdk
}

func (sdk *EchoApiSdk) Register(api pdc_swagger.Api, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *echo.Route {

	for _, middlewareFunc := range sdk.middlewares {
		middlewareFunc(api)
	}

	return sdk.E.Add(api.GetMethod(), api.GetRelativePath(), handler, middlewares...)
}
