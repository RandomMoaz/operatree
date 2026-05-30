package project

import (
	"fmt"
	"os"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

func EditSubjectMetadata(p *Project, s subject.Subject) error {
	if err := s.EditMetadata(); err != nil {
		return err
	}

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionEdit,
		string(s.Type),
		s.Name,
	); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}
