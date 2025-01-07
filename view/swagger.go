package view

import (
	"bytes"
	"text/template"
)

var ViewTemplate = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" type="text/css" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css">
  <title>{{ .Title }}</title>

<body>
  <div id="app" />

  <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function () {
      const ui = SwaggerUIBundle({
        url: "{{ .Url }}",
        dom_id: "#app",
        deepLinking: true,
      })
    }
  </script>
</body>
</html>`

func GetSwaggerViewTemplate(config *ViewTemplateConfig) (string, error) {
	var out bytes.Buffer

	viewTemplate := template.New("open_api_template")
	viewTemplate.Parse(ViewTemplate)

	err := viewTemplate.Execute(&out, config)
	return out.String(), err

}
