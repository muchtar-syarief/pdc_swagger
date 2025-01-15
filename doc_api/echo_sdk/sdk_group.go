package echo_sdk

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/muchtar-syarief/pdc_swagger"
)

type SdkGroup struct {
	sdk      *EchoApiSdk
	G        *echo.Group
	Basepath string
}

func (grp *SdkGroup) GetGinEngine() *echo.Echo {
	return grp.sdk.E
}

func (grp *SdkGroup) Register(api pdc_swagger.Api, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *echo.Route {
	api.SetGroupPath(grp.Basepath)

	for _, middlewareFunc := range grp.sdk.middlewares {
		middlewareFunc(api)
	}

	return grp.G.Add(api.GetMethod(), api.GetRelativePath(), handler, middlewares...)
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
