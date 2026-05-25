package template

type ModuleTemplate struct {
	Description string           `yaml:"description"`
	Type        string           `yaml:"type"`
	Name        string           `yaml:"name"`
	Modules     []ModuleTemplate `yaml:"modules"`
	SubDirs     []string         `yaml:"subDirs"`
}

type OTTemplate struct {
	Name        string           `yaml:"name"`
	Modules     []ModuleTemplate `yaml:"modules"`
	Description string           `yaml:"description"`
}

type tmpltMap map[string]string

var Templates tmpltMap = tmpltMap{
	"general":    "template_general.yml",
	"dev":        "template_dev.yml",
	"consulting": "template_consulting.yml",
	"research":   "template_research.yml",
}
