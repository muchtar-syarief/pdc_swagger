package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/muchtar-syarief/pdc_swagger/doc_api"
	"github.com/muchtar-syarief/pdc_swagger/doc_api/gin_sdk"
	"github.com/urfave/cli/v2"
)

func ginApi() {
	doc := pdc_swagger.NewPdcOpenApi("Test Documentation API Using Gin", "Description test documentation api using gin", "1.0.0")

	r := gin.Default()
	sdk := gin_sdk.NewGinApiSdk(r).
		UseDocumentation(doc).
		UseRedocDocumentation("/data_doc", "/redoc").
		UseSwaggerDocumentation("/data_doc", "/docs")

	sdk.Register(&doc_api.ApiData{
		Payload:      PayloadDataDD{},
		Method:       http.MethodPost,
		RelativePath: "/users",
		Tags:         []string{"Users"},
	}, func(ctx *gin.Context) {

	})

	datag := sdk.Group("/data")
	datag.Register(&doc_api.ApiData{
		Method: http.MethodGet,
		Tags:   []string{"Data"},
	}, func(ctx *gin.Context) {

	})

	usrg := datag.Group("/user")
	usrg.Register(&doc_api.ApiData{
		Method:       http.MethodPost,
		RelativePath: "create",
		Tags:         []string{"User"},
		Response:     ResponseData{},
	}, func(ctx *gin.Context) {})

	sdk.RegisterGroup("/product", func(group *gin.RouterGroup, register gin_sdk.RegisterFunc) {
		register(&doc_api.ApiData{
			Payload:      PayloadDataDD{},
			Method:       http.MethodPost,
			RelativePath: "/create",
			Tags:         []string{"Product"},
		})
	})

	sdk.RegisterGroup("/product_data", func(group *gin.RouterGroup, register gin_sdk.RegisterFunc) {
		register(&doc_api.ApiData{
			Payload:      []*PayloadDataDD{},
			Method:       http.MethodPost,
			Response:     []string{},
			RelativePath: "/create",
			Tags:         []string{"Product"},
		})
	})

	sdk.R.Run(":8200")
}

func GinApiCli() *cli.Command {
	return &cli.Command{
		Name:    "gin server",
		Aliases: []string{"gin"},
		Action: func(ctx *cli.Context) error {
			ginApi()
			return nil
		},
	}
}
