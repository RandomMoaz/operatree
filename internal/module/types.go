package module

import "github.com/hanymamdouh82/operatree/internal/subject"

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

type ModuleDirPrefix map[ModuleType]string

var (
	ModuleDirPrefixMap ModuleDirPrefix = ModuleDirPrefix{
		ModuleAdmin:             "00",
		ModuleEvents:            "01",
		ModuleProjectManagement: "02",
		ModuleLegal:             "03",
		ModuleResearch:          "04",
		ModuleEngineering:       "05",
		ModuleData:              "06",
		ModuleTasks:             "05",
		ModuleIndex:             "06",
		ModuleTopics:            "07",
		ModuleObjectives:        "08",
		ModuleSummaries:         "09",
		ModuleReferences:        "10",
		ModuleAudioNotes:        "11",
		ModuleAttachements:      "12",
		ModuleDataSources:       "13",
		ModulePublications:      "14",
		ModuleMediaLibrary:      "97",
		ModuleDeliverables:      "98",
		ModuleArchive:           "99",
	}
)

type Module struct {
	Type     ModuleType        `yaml:"type"`
	Name     string            `yaml:"name"`
	AbsPath  string            `yaml:"-"`        // abs path of the module not relative to project. Important for subjects and standlone modules
	Modules  []Module          `yaml:"modules"`  // use when subjects are nested into a sub-dir
	Subjects []subject.Subject `yaml:"subjects"` // use when module contains direct subjects such as 01_EVENTS
	SubDirs  []string          `yaml:"subDirs"`  // only flat dirs, they are not created initially by operatree, but can be created by Topic factory
}
