package project

import (
	"fmt"
	"os"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

// NewSubject creates a new subject of the specified type in its corresponding module.
// Parameters:
//   - p: target project
//   - subjectName: name of the subject
//   - subjectDate: date associated with the subject
//   - st: subject type (Event, Task, Topic, or Objective)
func NewSubject(p *Project, cliSubject subject.Subject, st subject.SubjectType) error {
	ss := ListSubjects(p, "")

	// Validate subject type and get corresponding module type
	tmt, exists := SubjectModuleMap[st]
	if !exists {
		return fmt.Errorf("unsupported subject type: %s", string(st))
	}

	// Find the target module for this subject type
	tm, err := findModule(p.Modules, tmt)
	if err != nil {
		return err
	}

	// Create the subject instance
	is := subject.Subject{
		Type:             st,
		Name:             cliSubject.Name,
		Date:             cliSubject.Date,
		Notes:            cliSubject.Notes,
		Tags:             cliSubject.Tags,
		Participants:     cliSubject.Participants,
		Location:         cliSubject.Location,
		Owner:            cliSubject.Owner,
		Status:           cliSubject.Status,
		RelatedObjective: cliSubject.RelatedObjective,
		RelatedEvents:    cliSubject.RelatedEvents,
		Outputs:          cliSubject.Outputs,
		Source:           cliSubject.Source,
		SourceLink:       cliSubject.SourceLink,
		SourceObjective:  cliSubject.SourceObjective,
		SourceDataSize:   cliSubject.SourceDataSize,
	}

	// Use factory to build subject with validation
	s, err := subject.SubjectFactory(is, tm.AbsPath, ss)
	if err != nil {
		return err
	}

	fmt.Printf("uuid: %s\n", s.UUID)

	// Persist subject to filesystem
	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// Update project metadata with new subject
	tm.Subjects = append(tm.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	// Confirm creation to user
	fmt.Printf("%s created: %s\n", string(st), s.Name)
	fmt.Printf("%s path: %s\n", string(st), s.DirName)

	// Log the action for audit trail
	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionCreate,
		string(st),
		s.Name,
	); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}
