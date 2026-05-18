package legal

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

const (
	UNIT_NAME = "03_LEGAL"
)

var (
	SubDirs = map[string]string{
		"approvals":  "approvals description here",
		"compliance": "compliance description here",
		"contracts":  "contracts description here",
		"ndas":       "ndas description here",
		"templates":  "templates description here",
	}
)

type UnitLegal struct {
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
	UnitPath   string `yaml:"unitPath"`
}

func (u *UnitLegal) SetUnitType(t string) {
	u.Type = t
}

func (u *UnitLegal) SetUnitName() {
	u.Name = UNIT_NAME
}

func (u *UnitLegal) SetParentDir(pth string) {
	u.ParentPath = pth
}

// Used with loaders
func (u *UnitLegal) SetUnitDir() {
	u.UnitPath = u.UnitDir()
}

func (u UnitLegal) UnitDir() string {
	return path.Join(u.ParentPath, UNIT_NAME)
}

// Cannot use *Event since it will not implement the interface
func (u UnitLegal) Bootstrap(ppth string) error {
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
