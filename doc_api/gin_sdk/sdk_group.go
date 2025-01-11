package gin_sdk

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/muchtar-syarief/pdc_swagger"
)

type SdkGroup struct {
	sdk      *GinApiSdk
	G        *gin.RouterGroup
	Basepath string
}

func (grp *SdkGroup) GetGinEngine() *gin.Engine {
	return grp.sdk.R
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

func (sdk *GinApiSdk) Group(relativePath string) *SdkGroup {
	newGroup := SdkGroup{
		sdk:      sdk,
		G:        sdk.R.Group(relativePath),
		Basepath: relativePath,
	}

	return &newGroup
}

type RegisterFunc func(api pdc_swagger.Api, handlers ...gin.HandlerFunc) gin.IRoutes

func (sdk *GinApiSdk) RegisterGroup(relativePath string, groupHandler func(group *gin.RouterGroup, register RegisterFunc)) {
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
