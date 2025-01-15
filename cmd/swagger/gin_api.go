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
	r := gin.Default()
	doc := pdc_swagger.NewPdcOpenApi("Test Documentation API Using Gin", "Description test documentation api using gin", "1.0.0")
	r.Handle(http.MethodGet, "/doc_data", func(ctx *gin.Context) {
		doc.RegisterDataDocumentation("/doc_data", func(method, path string) {
			ctx.YAML(200, doc)
		})
	})

	doc.RegisterSwaggerDocumentation("/doc_data", "/docs", func(method, path string, responseTemplate func() (string, error)) {
		r.Handle(method, path, func(ctx *gin.Context) {
			template, err := responseTemplate()
			if err != nil {
				ctx.JSON(500, gin.H{"status": "error"})
				return
			}

			ctx.Data(200, "text/html", []byte(template))
		})
	})

	doc.RegisterRedocDocumentation("/doc_data", "/redoc", func(method, path string, responseTemplate func() (string, error)) {
		r.Handle(method, path, func(ctx *gin.Context) {
			template, err := responseTemplate()
			if err != nil {
				ctx.JSON(500, gin.H{"status": "error"})
				return
			}

			ctx.Data(200, "text/html", []byte(template))
		})
	})

	sdk := gin_sdk.NewGinApiSdk(r).
		Use(doc.AddToDocumentation)

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
