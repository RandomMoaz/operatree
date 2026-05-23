package module

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

type ModuleType string

const (
	ModuleAdmin             ModuleType = "ADMIN"
	ModuleEvents            ModuleType = "EVENTS"
	ModuleProjectManagement ModuleType = "PROJECT_MANAGEMENT"
	ModuleTasks             ModuleType = "TASKS"
	ModuleLegal             ModuleType = "LEGAL"
	ModuleResearch          ModuleType = "RESEARCH"
	ModuleIndex             ModuleType = "INDEX"
	ModuleTopics            ModuleType = "TOPICS"
	ModuleObjectives        ModuleType = "OBJECTIVES"
	ModuleSummaries         ModuleType = "SUMMARIES"
	ModuleReferences        ModuleType = "REFERENCES"
	ModuleAudioNotes        ModuleType = "AUDIO_NOTES"
	ModuleAttachements      ModuleType = "ATTACHEMENTS"
	ModuleEngineering       ModuleType = "ENGINEERING"
	ModuleData              ModuleType = "DATA"
	ModuleDataSources       ModuleType = "DATA_SOURCES"
	ModuleMediaLibrary      ModuleType = "MEDIA_LIBRARY"
	ModuleDeliverables      ModuleType = "DELIVERABLES"
	ModulePublications      ModuleType = "PUBLICATIONS"
	ModuleArchive           ModuleType = "ARCHIVE"
)

type Module struct {
	Type     ModuleType        `yaml:"type"`
	Name     string            `yaml:"name"`
	AbsPath  string            `yaml:"absPath"`  // abs path of the module not relative to project. Important for subjects and standlone modules
	Modules  []Module          `yaml:"modules"`  // use when subjects are nested into a sub-dir
	Subjects []subject.Subject `yaml:"subjects"` // use when module contains direct subjects such as 01_EVENTS
	SubDirs  []string          `yaml:"subDirs"`  // only flat dirs, they are not created initially by operatree, but can be created by Topic factory
}

// A method to create module directory
func (m *Module) MkDir() error {

	if err := filesystem.CreateDir(m.AbsPath); err != nil {
		return err
	}

	return nil
}

// A method to create module sub directories
func (m *Module) MkSubDirs() error {

	for _, v := range m.SubDirs {
		sdp := path.Join(m.AbsPath, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

// Creates module dirs, subdirs, nested modules, metadata templates, etc
// This is to be used during project bootstrapping, or new module bootstrapping
func (m *Module) Bootstrap() error {
	// Create module directory
	if err := m.MkDir(); err != nil {
		return err
	}

	// Recursive bootstrapping for submodules
	for _, v := range m.Modules {
		v.Bootstrap()
	}

	// create module subdirs
	if err := m.MkSubDirs(); err != nil {
		return err
	}

	return nil
}
