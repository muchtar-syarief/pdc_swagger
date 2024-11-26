package pdc_swagger

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
	Title       string   `yaml:"title" json:"title"`
	Summary     string   `yaml:"summary" json:"summary"`
	Description string   `yaml:"description" json:"description"`
	License     *License `yaml:"license,omitempty" json:"license,omitempty"`
	Version     string   `yaml:"version" json:"version"`
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

func (i *PdcSwaggerInfo) SetLicence(license *License) *PdcSwaggerInfo {
	i.License = license
	return i
}
