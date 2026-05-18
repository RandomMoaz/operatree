package admin

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

const (
	UNIT_NAME = "01_ADMIN"
)

var (
	SubDirs = map[string]string{
		"contacts":   "contacts description here",
		"governance": "governance description here",
		"guidelines": "guidelines description here",
		"templates":  "templates description here",
	}
)

type UnitAdmin struct {
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
	UnitPath   string `yaml:"unitPath"`
}

func (u *UnitAdmin) SetUnitType(t string) {
	u.Type = t
}

func (u *UnitAdmin) SetUnitName() {
	u.Name = UNIT_NAME
}

func (u *UnitAdmin) SetParentDir(pth string) {
	u.ParentPath = pth
}

// Used with loaders
func (u *UnitAdmin) SetUnitDir() {
	u.UnitPath = u.UnitDir()
}

func (u UnitAdmin) UnitDir() string {
	return path.Join(u.ParentPath, UNIT_NAME)
}

// Cannot use *Event since it will not implement the interface
func (u UnitAdmin) Bootstrap(ppth string) error {
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
