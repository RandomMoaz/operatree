package project

import (
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

const (
	METADATA_FILE = "METADATA.yml"
	ARCHIVED_DEST = "closed_tasks"
)

type Project struct {
	Name     string          `yaml:"name"`
	Template string          `yaml:"template"`
	absDir   string          `yaml:"-"` // project absolute directory, hydrated during load
	Tags     []string        `yaml:"tags"`
	Modules  []module.Module `yaml:"modules"`
}

// SubjectModuleMap maps each subject type to its corresponding storage module
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
	subject.SubjectEvent:      module.ModuleEvents,
	subject.SubjectTask:       module.ModuleTasks,
	subject.SubjectTopic:      module.ModuleTopics,
	subject.SubjectObjective:  module.ModuleObjectives,
	subject.SubjectDataSource: module.ModuleDataSources,
}
