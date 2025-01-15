# pdc_swagger
When using golang to create backend application with `Gin`, `Echo`, and other. They doesn't have auto generated api documentation to show result api was made it.
This repo is used for golang developer to build api documentation with standard open api.
The documentation can be access with swagger mode or redoc mode.

## Swagger And Redoc Documentation
Example Usage :
```	
	r := gin.Default()
	doc := pdc_swagger.NewPdcOpenApi(
		"Test Documentation API Using Gin", 
		"Description test documentation api using gin", 
		"1.0.0",
		)
		
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
	Method:       http.MethodPost,
	RelativePath: "/users",
	Payload:      Payload{},
	Response: 	  Response{},
	Tags:         []string{"Users"},
}, func(ctx *gin.Context) {

})


sdkGroup := sdk.Group("v1")
sdkGroup.Register(&doc_api.ApiData{
	Method: 		http.MethodGet,
	RelativePath:   "/users",
	Query: 			Query{},
	Tags:   		[]string{"Data"},
}, func(ctx *gin.Context) {

})
```

## Type Can Generated
Type can use in this repo is:
- String
- Int
- Uint 
- Boolean
- Map
- Struct
- Pointer
- Generic
- Type Alias. Exclude: time alias


## Bug
- Type Alias for time.Time.<br>
Example:<br>
```
type TimeAlias time.Time

type Response struct {
	day TimeAlias `json:"day"`
}
```
