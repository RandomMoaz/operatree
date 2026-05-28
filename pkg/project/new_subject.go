package project

import (
	"fmt"
	"os"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

// NewSubject creates a new subject of the specified type in its corresponding module.
// Parameters:
//   - p: target project
//   - subjectName: name of the subject
//   - subjectDate: date associated with the subject
//   - st: subject type (Event, Task, Topic, or Objective)
func NewSubject(p *Project, subjectName, subjectDate string, st subject.SubjectType) error {
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
		Type: st,
		Name: subjectName,
		Date: subjectDate,
	}

	// Use factory to build subject with validation
	s, err := subject.SubjectFactory(is, tm.AbsPath, ss)
	if err != nil {
		return err
	}

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

// findModule recursively searches for a module by type within the project hierarchy.
func findModule(modules []module.Module, tmt module.ModuleType) (*module.Module, error) {
	for i, m := range modules {
		// Check if this module matches the target type
		if m.Type == tmt {
			return &modules[i], nil
		}

		// Recursively search submodules
		if len(m.Modules) > 0 {
			if found, err := findModule(m.Modules, tmt); err == nil && found != nil {
				return found, nil
			}
		}
	}

	return nil, fmt.Errorf("module type %s not found in project", string(tmt))
}
