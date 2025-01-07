package pdc_swagger

type ResponseObject struct {
	Description string                         `yaml:"description,omitempty" json:"description,omitempty"`
	Headers     map[string]interface{}         `yaml:"headers,omitempty" json:"headers,omitempty"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
	Links       map[string]interface{}         `yaml:"links,omitempty" json:"links,omitempty"`
}

type RequestBodyObject struct {
	Description string                         `yaml:"description,omitempty" json:"description,omitempty"`
	Required    bool                           `yaml:"required,omitempty" json:"required,omitempty"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
}

type OperationObject struct {
	Tags        []string                           `yaml:"tags,omitempty" json:"tags,omitempty"`
	Summary     string                             `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string                             `yaml:"description,omitempty" json:"description,omitempty"`
	OperationId string                             `yaml:"operationId,omitempty" json:"operationId,omitempty"`
	Parameters  []*ParameterObject                 `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	RequestBody *RequestBodyObject                 `yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	Responses   map[HTTPStatusCode]*ResponseObject `yaml:"responses,omitempty" json:"responses,omitempty"`
	Callbacks   interface{}                        `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	Deprecated  bool                               `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	Security    interface{}                        `yaml:"security,omitempty" json:"security,omitempty"`
	Servers     []*ServerObject                    `yaml:"servers,omitempty" json:"servers,omitempty"`
}

func NewOperationObject(tags []string, summary, desc, operationID string) *OperationObject {
	return &OperationObject{
		Tags:        tags,
		Summary:     summary,
		Description: desc,
		OperationId: operationID,
	}
}

func (o *OperationObject) SetParameters(data interface{}) *OperationObject {
	if data == nil {
		return o
	}

	if o.Parameters == nil {
		o.Parameters = []*ParameterObject{}
	}

	parameters := NewListParametersObject(data)
	o.Parameters = append(o.Parameters, parameters...)

	return o
}

func (o *OperationObject) SetRequestBody(data interface{}) *OperationObject {
	if data == nil {
		return o
	}

	if o.RequestBody == nil {
		o.RequestBody = &RequestBodyObject{}
	}

	schemaPayload := NewSchema(data)

	o.RequestBody = &RequestBodyObject{
		Description: "",
		Required:    false,
		Content: map[MediaType]*MediaTypeObject{
			MediaTypeJson: {
				Schema: schemaPayload,
			},
		},
	}

	return o
}

func (o *OperationObject) SetResponse(code string, data interface{}) *OperationObject {
	if data == nil {
		return o
	}

	if o.Responses == nil {
		o.Responses = map[HTTPStatusCode]*ResponseObject{}
	}

	responseSchema := NewSchema(data)
	o.Responses[HTTPStatusCode(code)] = &ResponseObject{
		Content: map[MediaType]*MediaTypeObject{
			MediaTypeJson: {
				Schema: responseSchema,
			},
		},
	}

	return o
}
