# pdc_swagger
When using golang to create backend application with `Gin`, `Echo`, and other. They doesn't have auto generated api documentation to show result api was made it.
This repo is used for golang developer to build api documentation with standard open api.
The documentation can be access with swagger mode or redoc mode.

## Swagger And Redoc Documentation
Example Usage :
```	
doc := pdc_swagger.NewPdcOpenApi("Test Documentation API", "Description test documentation api", "1.0.0")

r := gin.Default()
sdk := doc_api.NewApiSdk(r).
		UseDocumentation(doc).
		UseRedocDocumentation("/data_doc", "/redoc").
		UseSwaggerDocumentation("/data_doc", "/docs")

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
