# pdc_swagger
This repo is used for golang developer to build api documentation with standard open api.
The documentation can be access with swagger mode or redoc mode.

Example Usage:
```	
doc := pdc_swagger.NewPdcOpenApi("Test Documentation API", "Description test documentation api", "1.0.0")

r := gin.Default()
sdk := doc_api.NewApiSdk(r).
		UseDocumentation(doc).
		UseRedocDocumentation("/data_doc", "/redoc")
```