package view

import (
	"bytes"
	"text/template"
)

var RedocViewTemplate = `<!DOCTYPE html>
<html>
  <head>
    <title>{{ .Title }}</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

    <!--
    Redoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='{{ .Url }}'></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
  </body>
</html>`

func GetRedocViewTemplate(config *ViewTemplateConfig) (string, error) {
	var out bytes.Buffer

	viewTemplate := template.New("redoc_template")
	viewTemplate.Parse(RedocViewTemplate)

	err := viewTemplate.Execute(&out, config)
	return out.String(), err

}
