package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/muchtar-syarief/pdc_swagger/doc_api"
	"github.com/muchtar-syarief/pdc_swagger/doc_api/echo_sdk"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func echoApi() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}	${method}	${status}		 ${uri}\n",
	}))

	doc := pdc_swagger.NewPdcOpenApi("Test Documentation API Using Echo", "Description test documentation api using Echo", "1.0.0")
	doc.RegisterDataDocumentation("/doc_data", func(method, path string) {
		e.Add(method, path, func(ctx echo.Context) error {
			raw, err := yaml.Marshal(doc)
			if err != nil {
				return ctx.JSON(500, echo.Map{"status": "error"})
			}

			return ctx.Blob(http.StatusOK, "application/x-yaml", raw)
		})
	})

	doc.RegisterSwaggerDocumentation("/doc_data", "/docs", func(method, path string, responseTemplate func() (string, error)) {
		e.Add(method, path, func(ctx echo.Context) error {
			template, err := responseTemplate()
			if err != nil {
				return ctx.JSON(500, echo.Map{"status": "error"})
			}

			return ctx.HTML(200, template)
		})
	})

	doc.RegisterRedocDocumentation("/doc_data", "/redoc", func(method, path string, responseTemplate func() (string, error)) {
		e.Add(method, path, func(ctx echo.Context) error {
			template, err := responseTemplate()
			if err != nil {
				return ctx.JSON(500, echo.Map{"status": "error"})
			}

			return ctx.HTML(200, template)
		})
	})

	sdk := echo_sdk.NewEchoApiSdk(e).
		Use(doc.AddToDocumentation)

	sdk.Register(&doc_api.ApiData{
		Payload:      PayloadDataDD{},
		Method:       http.MethodPost,
		RelativePath: "/users",
		Tags:         []string{"Users"},
	}, func(ctx echo.Context) error {

		return nil
	})

	datag := sdk.Group("/data")
	datag.Register(&doc_api.ApiData{
		Method: http.MethodGet,
		Tags:   []string{"Data"},
	}, func(ctx echo.Context) error {
		return nil
	})

	usrg := datag.Group("/user")
	usrg.Register(&doc_api.ApiData{
		Method:       http.MethodPost,
		RelativePath: "create",
		Tags:         []string{"User"},
		Response:     ResponseData{},
	}, func(ctx echo.Context) error {
		return nil
	})

	sdk.RegisterGroup("/product", func(group *echo.Group, register echo_sdk.RegisterFunc) {
		register(&doc_api.ApiData{
			Payload:      PayloadDataDD{},
			Method:       http.MethodPost,
			RelativePath: "/create",
			Tags:         []string{"Product"},
		}, func(ctx echo.Context) error {
			return nil
		})
	})

	sdk.RegisterGroup("/product_data", func(group *echo.Group, register echo_sdk.RegisterFunc) {
		register(&doc_api.ApiData{
			Payload:      []*PayloadDataDD{},
			Method:       http.MethodPost,
			Response:     []string{},
			RelativePath: "/create",
			Tags:         []string{"Product"},
		}, func(ctx echo.Context) error {
			return nil
		})
	})

	// doc.Save("./doc.json")
	e.Logger.Fatal(e.Start(":8300"))
}

func EchoApiCli() *cli.Command {
	return &cli.Command{
		Name:    "echo server",
		Aliases: []string{"echo"},
		Action: func(ctx *cli.Context) error {
			echoApi()
			return nil
		},
	}
}
