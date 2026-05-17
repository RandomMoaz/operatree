package project

import "path"

type Unit interface {
	Bootstrap() error
}

type Project struct {
	name    string   `yaml:"name"`
	baseDir string   `yaml:"baseDir"`
	Tags    []string `yaml:"tags"`
	Units   []Unit
}

func (p *Project) ProjectName() string {
	return p.name
}

// Returns root dir of the project. It is the dir where project resides
func (p *Project) ProjectRootDir() string {
	return p.baseDir
}

// Returns full project path
func (p *Project) ProjectDir() string {
	return path.Join(p.baseDir, p.name)
}
