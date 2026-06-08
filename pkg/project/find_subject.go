package project

import (
	"fmt"

	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func FindSubjectByID(p *Project, id string) (*subject.Subject, error) {
	result := findSubjectByIDInModules(p.Modules, id)
	if result == nil {
		return nil, fmt.Errorf("subject with ID %s not found", id)
	}
	return result, nil
}

func findSubjectByIDInModules(modules []module.Module, id string) *subject.Subject {
	for i, m := range modules {
		// Check subjects at this level
		for j, s := range m.Subjects {
			if s.UUID == id {
				return &modules[i].Subjects[j]
			}
		}

		// Not found here — go deeper
		if found := findSubjectByIDInModules(m.Modules, id); found != nil {
			return found
		}
	}

	return nil
}
