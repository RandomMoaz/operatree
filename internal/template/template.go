package template

import (
	"embed"

	"gopkg.in/yaml.v3"
)

//go:embed *.yml
var FS embed.FS

// Reads template file from embedded templates directory (go:embed)
// returns parsed and validated template that is parsable by all factory functions
func Load(name string) (OTTemplate, error) {
	// name = "template_general.yml"

	// Read temaplte file from embedded template files
	data, err := FS.ReadFile(name)
	if err != nil {
		return OTTemplate{}, err
	}

	// parse yaml into OTTemplate
	var pt OTTemplate
	if err := yaml.Unmarshal(data, &pt); err != nil {
		return pt, err
	}

	return pt, nil
}

func (t *OTTemplate) TemplateDescription() string {

	return t.Description
}

func (m *ModuleTemplate) ModuleDescription() string {

	return m.Description
}

func (t *OTTemplate) FullDescription() string {
	// To-Do: replace with full implementation
	return "auto-generated-description-place-holder"
}
