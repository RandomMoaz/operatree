package subject

type SubjectType string
type SubjectDirMap map[SubjectType][]string
type SubjectFilesMap map[SubjectType][]string

// Defines subject metadata file name. Search, sycn, etc depends on this name.
const (
	METADATA_FILE = "METADATA.yml"
)

// Subjects, each entry defines a new subject
const (
	SubjectEvent     SubjectType = "EVENT"
	SubjectTask      SubjectType = "TASK"
	SubjectTopic     SubjectType = "TOPIC"
	SubjectObjective SubjectType = "OBJECTIVE"
)

// Default subdirs for subject type. These dirs are created automatically during bootstrapping of a project
var (
	SubDirs SubjectDirMap = SubjectDirMap{
		SubjectEvent: {
			"01_AGENDA",
			"02_MEDIA",
			"03_NOTES",
			"04_DOCUMENTS",
			"05_OUTCOMES",
		},
		SubjectTask: {
			"01_INPUTS",
			"02_WORKING",
			"03_REVIEW",
			"04_FINAL",
		},
	}

	// Those are empty files to be created inside subject dir
	Files SubjectFilesMap = SubjectFilesMap{
		SubjectTopic: {
			"overview.md",
			"notes.md",
		},
		SubjectObjective: {
			"definitions.md",
			"findings.md",
			"strategy.md",
		},
	}
)

// managed by operatree, can add/delete/edit, etc
// searchable, indexable, parsed by describe()
// This is like 01_RAW inside 06_DATA module
type Subject struct {
	Type    SubjectType `yaml:"type"`
	Name    string      `yaml:"name"`
	DirName string      `yaml:"-"`
	SubDirs []string    `yaml:"subDirs"`
	Files   []string    `yaml:"-"`
	Date    string      `yaml:"date"`
	Tags    []string    `yaml:"tags"`
	Notes   string      `yaml:"notes"`
	// Custom fields based on subject type, you must use `omitempty` to avoid parsing if not used
	Paricipants      []string `yaml:"paricipants,omitempty"` // omitempty guarantees that field written only for Subject that needs it
	Location         string   `yaml:"location,omitempty"`    // omitempty guarantees that field written only for Subject that needs it
	Owner            string   `yaml:"owner,omitempty"`
	Status           string   `yaml:"status,omitempty"`
	RelatedObjective string   `yaml:"related_objective,omitempty"`
	RelatedEvents    []string `yaml:"related_events,omitempty"`
	Outputs          []string `yaml:"outputs,omitempty"`
}
