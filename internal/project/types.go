package project

import (
	"path"
)

const (
	METADATA_FILE = "metadata.yml"
)

type UnitType string

const (
	UnitAdmin        UnitType = "ADMIN"
	UnitArchive      UnitType = "ARCHIVE"
	UnitDeliverables UnitType = "DELIVERABLES"
	UnitEngineering  UnitType = "ENGINEERING"
	UnitEvents       UnitType = "EVENTS"
	UnitLegal        UnitType = "LEGAL"
	UnitMediaLib     UnitType = "MEDIALIB"
)

type Unit interface {
	Bootstrap(pth string) error
	UnitDir() string
	SetParentDir(pth string)
	SetUnitName()
	SetUnitDir()
	SetUnitType(t string)
}

type Project struct {
	Name    string   `yaml:"name"`
	BaseDir string   `yaml:"baseDir"`
	Tags    []string `yaml:"tags"`
	Units   []Unit   `yaml:"units"`
}

func (p *Project) ProjectName() string {
	return p.Name
}

func (p *Project) MetadataFile() string {
	return path.Join(p.ProjectDir(), METADATA_FILE)
}

// Returns base dir of the project. It is the dir where project resides
func (p *Project) ProjectBaseDir() string {
	return p.BaseDir
}

// Returns full project path including project name.
// Ex: /mnt/repos/porjects/my_project
// never use baseDir property, always use reciever function whenever project path is required
func (p *Project) ProjectDir() string {
	return path.Join(p.BaseDir, p.Name)
}

// Use to add unit to project, it is reponsible to set required project properties into unit
// t: unit type, it is very important since it defines unit types during loading
func (p *Project) AddUnit(u Unit, t UnitType) {
	u.SetUnitType(string(t))
	u.SetUnitName()
	u.SetParentDir(p.ProjectDir())
	u.SetUnitDir()
	p.Units = append(p.Units, u)
}
