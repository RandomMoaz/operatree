package engineering

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

const (
	UNIT_NAME = "05_ENGINEERING"
)

var (
	SubDirs = map[string]string{
		"architecture":   "architecture description here",
		"decisions":      "decisions description here",
		"diagrams":       "diagrams description here",
		"prototypes":     "prototypes description here",
		"simulations":    "simulations description here",
		"specifications": "specifications description here",
		"templates":      "templates description here",
	}
)

type UnitEngineering struct {
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
	UnitPath   string `yaml:"unitPath"`
}

func (u *UnitEngineering) SetUnitType(t string) {
	u.Type = t
}

func (u *UnitEngineering) SetUnitName() {
	u.Name = UNIT_NAME
}

func (u *UnitEngineering) SetParentDir(pth string) {
	u.ParentPath = pth
}

// Used with loaders
func (u *UnitEngineering) SetUnitDir() {
	u.UnitPath = u.UnitDir()
}

func (u UnitEngineering) UnitDir() string {
	return path.Join(u.ParentPath, UNIT_NAME)
}

// Cannot use *Event since it will not implement the interface
func (u UnitEngineering) Bootstrap(ppth string) error {
	if err := filesystem.CreateDir(u.UnitDir()); err != nil {
		return err
	}

	// create sub dirs, ADMIN doesn't include units, instead it includes only flat sub dirs
	errs := make([]error, 0)
	for k, v := range SubDirs {
		if err := filesystem.CreateDir(path.Join(u.UnitDir(), k)); err != nil {
			errs = append(errs, err)
			continue
		}
		fmt.Printf("Created sub-dir: %s\n", v)
	}

	return nil
}
