// this file contains the most comming primary modules
package module

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/subject"
)

func FactoryAdmin(ppth string) Module {

	n := "00_ADMIN"
	m := Module{
		Type:     ModuleAdmin,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"contacts",
			"governance",
			"guidelines",
			"templates",
		},
	}

	return m
}

func FactoryEvents(ppth string) Module {

	n := "01_EVNTS"
	m := Module{
		Type:     ModuleEvents,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs:  []string{},
	}

	return m
}

func FactoryProjectManagement(ppth string) Module {

	n := "02_PROJECT_MANAGEMENT"
	m := Module{
		Type:    ModuleProjectManagement,
		Name:    n,
		AbsPath: path.Join(ppth, n),
		Modules: []Module{
			{
				Type:     ModuleTasks,
				Name:     "01_TASKS",
				AbsPath:  path.Join(ppth, n, "01_TASKS"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
		},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"budgets",
			"communications",
			"planning",
			"reports",
			"risks",
		},
	}

	return m
}

// Legal module
func FactoryLegal(ppth string) Module {

	n := "03_LEGAL"
	m := Module{
		Type:     ModuleLegal,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"contracts",
			"ndas",
			"compliance",
			"approvals",
			"templates",
		},
	}

	return m
}

// Research module, basically it is a collection of submodules
func FactoryResearch(ppth string) Module {

	n := "04_RESEARCH"
	m := Module{
		Type:    ModuleResearch,
		Name:    n,
		AbsPath: path.Join(ppth, n),
		Modules: []Module{
			{
				Type:     ModuleIndex,
				Name:     "00_INDEX",
				AbsPath:  path.Join(ppth, n, "00_INDEX"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
			{
				Type:     ModuleTopics,
				Name:     "01_TOPICS",
				AbsPath:  path.Join(ppth, n, "01_TOPICS"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
			{
				Type:     ModuleObjectives,
				Name:     "02_OBJECTIVES",
				AbsPath:  path.Join(ppth, n, "02_OBJECTIVES"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
			{
				Type:     ModuleSummaries,
				Name:     "03_SUMMARIES",
				AbsPath:  path.Join(ppth, n, "03_SUMMARIES"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
			{
				Type:     ModuleReferences,
				Name:     "04_REFERENCES",
				AbsPath:  path.Join(ppth, n, "04_REFERENCES"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs: []string{
					"articles",
					"books",
					"papers",
					"standards",
					"vendor_docs",
				},
			},
			{
				Type:     ModuleAudioNotes,
				Name:     "05_AUDIO_NOTES",
				AbsPath:  path.Join(ppth, n, "05_AUDIO_NOTES"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs: []string{
					"raw",
					"transcriptions",
					"indexed",
				},
			},
			{
				Type:     ModuleAttachements,
				Name:     "06_ATTACHEMENTS",
				AbsPath:  path.Join(ppth, n, "06_ATTACHEMENTS"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
		},
		Subjects: []subject.Subject{},
		SubDirs:  []string{},
	}

	return m
}

// Engineering module
func FactoryEngineering(ppth string) Module {

	n := "05_ENGINEERING"
	m := Module{
		Type:     ModuleEngineering,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"architecture",
			"decisions",
			"diagrams",
			"prototypes",
			"simulations",
			"specifications",
			"templates",
		},
	}

	return m
}

// Data module, this is always found in `dev` template
// Although subdirs looks like modules; they are actually non-managed modules
// Sources is a sub-module, since it will be controlled from CLI and it has metadata to index Sources
// and ensure they are searchable
func FactoryData(ppth string) Module {

	n := "06_DATA"
	m := Module{
		Type:    ModuleData,
		Name:    n,
		AbsPath: path.Join(ppth, n),
		Modules: []Module{
			{
				Type:     ModuleDataSources,
				Name:     "00_SOURCES",
				AbsPath:  path.Join(ppth, n, "00_SOURCES"),
				Modules:  []Module{},
				Subjects: []subject.Subject{},
				SubDirs:  []string{},
			},
		},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"01_RAW",
			"02_STAGING",
			"03_PROCESSED",
			"04_ANALYTICS",
			"05_MODELS",
			"06_EXPORTS",
			"99_ARCHIVE",
		},
	}

	return m
}

// Media Library module
func FactoryMediaLib(ppth string) Module {

	n := "07_MEDIA_LIBRARY"
	m := Module{
		Type:     ModuleMediaLibrary,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"photos_clean",
			"videos",
			"diagrams",
			"presentation_assets",
			"branding",
		},
	}

	return m
}

// Media Deliverables
func FactoryDeliverables(ppth string) Module {

	n := "08_DELIVERABLES"
	m := Module{
		Type:     ModuleDeliverables,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"client_documents",
			"presentations",
			"reports",
			"submissions",
		},
	}

	return m
}

// Archive module, just an empty dir. It is managed by special CLI command `archive` that moves
// any file to the archive dir
func FactoryArchive(ppth string) Module {

	n := "09_ARCHIVE"
	m := Module{
		Type:     ModuleArchive,
		Name:     n,
		AbsPath:  path.Join(ppth, n),
		Modules:  []Module{},
		Subjects: []subject.Subject{},
		SubDirs: []string{
			"old_versions",
			"closed_tasks",
			"deprecated",
		},
	}

	return m
}
