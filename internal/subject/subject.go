package subject

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/runner"
)

type SubjectType string

const (
	METADATA_FILE = "METADATA.yml"
)

const (
	SubjectEvent     SubjectType = "EVENT"
	SubjectTask      SubjectType = "TASK"
	SubjectTopic     SubjectType = "TOPIC"
	SubjectObjective SubjectType = "OBJECTIVE"
)

var (
	SubDirs map[SubjectType][]string = map[SubjectType][]string{
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
	Files map[SubjectType][]string = map[SubjectType][]string{
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
	Type             SubjectType `yaml:"type"`
	Name             string      `yaml:"name"`
	DirName          string      `yaml:"dirName"`
	SubDirs          []string    `yaml:"subDirs"`
	Files            []string    `yaml:"-"`
	Date             string      `yaml:"date"`
	Tags             []string    `yaml:"tags"`
	Notes            string      `yaml:"notes"`
	Paricipants      []string    `yaml:"paricipants,omitempty"` // omitempty guarantees that field written only for Subject that needs it
	Location         string      `yaml:"location,omitempty"`    // omitempty guarantees that field written only for Subject that needs it
	Owner            string      `yaml:"owner,omitempty"`
	Status           string      `yaml:"status,omitempty"`
	RelatedObjective string      `yaml:"related_objective,omitempty"`
	RelatedEvents    []string    `yaml:"related_events,omitempty"`
	Outputs          []string    `yaml:"outputs,omitempty"`
}

// A method to create module directory
func (s *Subject) MkDir() error {

	if err := filesystem.CreateDir(s.DirName); err != nil {
		return err
	}

	return nil
}

// A method to create module sub directories
func (s *Subject) MkSubDirs() error {

	for _, v := range s.SubDirs {
		sdp := path.Join(s.DirName, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

// Creates empty files on disk
func (s *Subject) WriteFiles() error {
	for _, f := range s.Files {
		sdp := path.Join(s.DirName, f)
		t := fmt.Sprintf("# %s\n", s.Name)
		if err := filesystem.TextToMDFile(t, sdp); err != nil {
			return err
		}
	}

	return nil
}

// Writes metadata.yml file for the subject at subject dir
func (s *Subject) WriteMetadata() error {

	fn := path.Join(s.DirName, METADATA_FILE)
	if err := filesystem.StructToFile(s, fn); err != nil {
		return err
	}

	return nil
}

// Writes subject to disk
func (s *Subject) WriteToDisk() error {

	// make dir
	if err := s.MkDir(); err != nil {
		return err
	}

	// make subdirs
	if err := s.MkSubDirs(); err != nil {
		return err
	}

	// Create emtpy subject files
	if err := s.WriteFiles(); err != nil {
		return err
	}

	// write metadata file
	if err := s.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

// Prints subject on screen in a prettey style
// Not all properties are displayed, such as subDirs not included
// This function is intended to be used for subject briefing only
func (s *Subject) Describe() {
	describe(s)
}

func (s *Subject) EditMetadata() error {

	// call config to get editor

	// run editor

	fn := path.Join(s.DirName, METADATA_FILE)
	if err := runner.EditFile(fn); err != nil {
		return err
	}

	return nil
}
