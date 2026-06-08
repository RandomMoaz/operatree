package project

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"gopkg.in/yaml.v3"
)

func (p *Project) ProjectName() string {
	return p.Name
}

// Returns base dir of the project. It is the dir where project resides
func (p *Project) ProjectBaseDir() string {
	bd := strings.TrimSuffix(p.absDir, p.Name)
	return bd
}

// Returns full project path including project name.
// Ex: /mnt/repos/porjects/my_project
// never use baseDir property, always use reciever function whenever project path is required
func (p *Project) ProjectDir() string {
	// return path.Join(p.BaseDir, p.Name)
	return p.absDir
}

// Prints project contents on stdout
// It is useful since users of CLI always needs a way to display all project details
// the displayed information should follow UNIX/Linux style so it can be piped to
// or chained with GNU tools such as `sed`, `cut`, `grep`, etc
func (p *Project) Describe(plain bool) error {
	if plain {
		y, err := yaml.Marshal(p)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", y)
		return nil
	}

	describeProject(p)
	return nil
}

func (p *Project) WriteMetadata() error {

	fn := filepath.Join(p.ProjectDir(), METADATA_FILE)
	if err := filesystem.StructToFile(p, fn); err != nil {
		return err
	}

	return nil
}

// Confirms module exists, if exists it returns the module.
// It checks root level only, not nested modules.
func (p *Project) ModuleExists(name string) (module.Module, error) {

	midx := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Name == name
	})

	if midx == -1 {
		return module.Module{}, fmt.Errorf("project doesn't contain module %s", name)
	}

	return p.Modules[midx], nil
}

// Archives a subject by moving to project Archive module
func (p *Project) Archive(s subject.Subject) error {

	if err := Archive(p, s); err != nil {
		return err
	}

	return nil
}

// finds subject within a project and renames it, and updates project METADATA.yml
func (p *Project) RenameSubject(st, term, newName string, uuid string) error {

	var s subject.Subject
	var err error

	// find the required subject using interactive CLI for user to select the required one
	if uuid == "" {
		s, err = FindSubject(p, st, term)
		if err != nil {
			return err
		}
	} else {
		sp, err := findSubjectByID(p, uuid)
		if err != nil {
			return err
		}

		s = *sp
	}

	if s.Type == "" {
		return fmt.Errorf("couldn't identify subject type")
	}

	// rename subject. internally it updates subject METADATA also
	if err := s.Rename(newName); err != nil {
		return err
	}

	// find subject within the project to update the project metadata
	ps, err := findSubjectByID(p, s.UUID)
	if err != nil {
		return err
	}

	// deference edited subject with found in project
	*ps = s

	// update project metadata for the subject
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	return nil
}
