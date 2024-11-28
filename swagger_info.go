package pdc_swagger

type Contact struct {
	Name  string `yaml:"name" json:"name"`
	Url   string `yaml:"url" json:"url"`
	Email string `yaml:"email" json:"email"`
}

type License struct {
	Name       string `yaml:"name" json:"name"`
	Identifier string `yaml:"identifier" json:"identifier"`
	Url        string `yaml:"url" json:"url"`
}

var LicenceApache = &License{
	Name:       "Apache 2.0",
	Identifier: "Apache-2.0",
	Url:        "https://www.apache.org/licenses/LICENSE-2.0.html",
}

type PdcSwaggerInfo struct {
	Title          string   `yaml:"title" json:"title"`
	Summary        string   `yaml:"summary" json:"summary"`
	Description    string   `yaml:"description" json:"description"`
	TermsOfService string   `yaml:"termsOfService" json:"termsOfService"`
	Contact        *Contact `yaml:"contact,omitempty" json:"contact,omitempty"`
	License        *License `yaml:"license,omitempty" json:"license,omitempty"`
	Version        string   `yaml:"version" json:"version"`
}

func NewPdcSwaggerInfo(
	title,
	description,
	version string,
) *PdcSwaggerInfo {
	return &PdcSwaggerInfo{
		Title:       title,
		Description: description,
		Version:     version,
	}
}

func (i *PdcSwaggerInfo) SetSummary(summary string) *PdcSwaggerInfo {
	i.Summary = summary
	return i
}

func (i *PdcSwaggerInfo) SetTermOfService(termOfService string) *PdcSwaggerInfo {
	i.TermsOfService = termOfService
	return i
}

func (i *PdcSwaggerInfo) SetLicence(license *License) *PdcSwaggerInfo {
	i.License = license
	return i
}

func (i *PdcSwaggerInfo) SetContact(name, url, email string) *PdcSwaggerInfo {
	i.Contact = &Contact{
		Name:  name,
		Url:   url,
		Email: email,
	}
	return i
}
