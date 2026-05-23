package project

import (
	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

const (
	METADATA_FILE   = "METADATA.yml"
	ARCHIVED_DEST   = "closed_tasks"
	TMPLT_GENERAL   = "general"
	TMPLT_DEV       = "dev"
	TMPL_CONSULTING = "consulting"
	TMPL_RESEARCH   = "research"
)

type Project struct {
	Name     string          `yaml:"name"`
	Template string          `yaml:"template"`
	BaseDir  string          `yaml:"baseDir"`
	Tags     []string        `yaml:"tags"`
	Modules  []module.Module `yaml:"modules"`
}

// project templates map
type tmpltMap map[string]func(name string, bpth string) Project

var Templates tmpltMap = tmpltMap{
	TMPLT_GENERAL:   tmpltGeneral,
	TMPLT_DEV:       tmpltDev,
	TMPL_CONSULTING: tmpltConsulting,
	TMPL_RESEARCH:   tmpltResearch,
}

// SubjectModuleMap maps each subject type to its corresponding storage module
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
	subject.SubjectEvent:     module.ModuleEvents,
	subject.SubjectTask:      module.ModuleTasks,
	subject.SubjectTopic:     module.ModuleTopics,
	subject.SubjectObjective: module.ModuleObjectives,
}
