package project

import (
	"fmt"
	"os"
	"path"

	"github.com/hanymamdouh82/operatree/internal/units/archive"
	"gopkg.in/yaml.v3"
)

type unitLoader struct {
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
	UnitPath   string `yaml:"unitPath"`
}

type projectLoader struct {
	Name    string       `yaml:"name"`
	BaseDir string       `yaml:"baseDir"`
	Tags    []string     `yaml:"tags"`
	Units   []unitLoader `yaml:"units"`
}

// Loads a project by reading project metadata file and sets project structrue
// Path represents project root path
func Load(pth string) error {

	b, err := os.ReadFile(path.Join(pth, METADATA_FILE))
	if err != nil {
		return err
	}

	var pl projectLoader
	if err := yaml.Unmarshal(b, &pl); err != nil {
		return err
	}

	j, _ := yaml.Marshal(pl)
	fmt.Printf("%s\n", j)

	p, err := build(pl)
	k, _ := yaml.Marshal(p)
	fmt.Printf("%s\n", k)

	return nil
}

func build(pl projectLoader) (Project, error) {

	p := Project{
		Name:    pl.Name,
		BaseDir: pl.BaseDir,
		Tags:    pl.Tags,
	}

	// load units here
	for _, v := range pl.Units {
		switch v.Type {
		case string(UnitArchive):
			// To-Do: implement LoadUnit instead of AddUnit
			p.AddUnit(&archive.UnitArchive{}, UnitArchive)
		}
	}

	return p, nil
}
