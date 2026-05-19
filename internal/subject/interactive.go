package subject

import (
	"time"

	"github.com/hanymamdouh82/operatree/internal/metadata"
	"github.com/manifoldco/promptui"
)

func interactiveCLI(st SubjectType, s *Subject) error {

	// Standard props
	prompt := promptui.Prompt{
		Label: "Event name",
	}
	name, err := prompt.Run()
	if err != nil {
		return err
	}

	// Date prompt
	prompt = promptui.Prompt{
		Label:   "Date",
		Default: time.Now().Format("2006-01-02"),
	}
	date, err := prompt.Run()
	if err != nil {
		return err
	}

	// Tags prompt
	prompt = promptui.Prompt{
		Label: "Tags (comma-separated)",
	}
	tags, err := prompt.Run()
	if err != nil {
		return err
	}

	// Notes prompt
	prompt = promptui.Prompt{
		Label: "Notes",
	}
	notes, err := prompt.Run()
	if err != nil {
		return err
	}

	s.Name = name
	s.Date = date
	s.Tags = metadata.ParseTags(tags)
	s.Notes = notes

	// Custom props based on type
	if st == SubjectEvent {

		// Location prompt
		prompt = promptui.Prompt{
			Label: "Location",
		}
		location, err := prompt.Run()
		if err != nil {
			return err
		}

		// Participants prompt
		prompt = promptui.Prompt{
			Label: "Participants (comma-separated)",
		}
		participants, err := prompt.Run()
		if err != nil {
			return err
		}

		s.Location = location
		s.Paricipants = metadata.ParseParticipants(participants)
	}

	return nil
}
