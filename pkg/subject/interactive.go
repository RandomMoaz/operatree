package subject

import (
	"time"

	"github.com/charmbracelet/huh"
	"github.com/hanymamdouh82/operatree/internal/metadata"
)

func interactiveCLI(st SubjectType, s *Subject, ss []Subject) error {

	var name, date, tags, notes string
	date = time.Now().Format("2006-01-02")

	// Standard fields — all types
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Value(&name),

			huh.NewInput().
				Title("Date").
				Value(&date).
				Placeholder(time.Now().Format("2006-01-02")),

			huh.NewInput().
				Title("Tags").
				Description("comma-separated").
				Value(&tags),

			huh.NewText().
				Title("Notes").
				Value(&notes),
		),
	).Run()
	if err != nil {
		return err
	}

	s.Name = name
	s.Date = date
	s.Tags = metadata.ParseTags(tags)
	s.Notes = notes

	// Event-specific fields
	if st == SubjectEvent {
		var location, participants string

		err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Location").
					Value(&location),

				huh.NewInput().
					Title("Participants").
					Description("comma-separated").
					Value(&participants),
			),
		).Run()
		if err != nil {
			return err
		}

		s.Location = location
		s.Participants = metadata.ParseParticipants(participants)
	}

	// Task-specific fields
	if st == SubjectTask && len(ss) > 0 {
		var owner, status string

		// Build options for related events select
		eventOptions := make([]huh.Option[string], len(ss))
		for i, v := range ss {
			eventOptions[i] = huh.NewOption(v.Name, v.Name)
		}

		err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Owner").
					Value(&owner),

				huh.NewSelect[string]().
					Title("Status").
					Options(
						huh.NewOption("Planned", "planned"),
						huh.NewOption("In Progress", "in-progress"),
						huh.NewOption("Postponed", "postponed"),
						huh.NewOption("Done", "done"),
					).
					Value(&status),

				huh.NewMultiSelect[string]().
					Title("Related Subject").
					Options(eventOptions...).
					Value(&s.RelatedEvents), // stays empty slice if nothing selected
			),
		).Run()
		if err != nil {
			return err
		}

		s.Owner = metadata.ParsePersonName(owner)
		s.Status = status
		// s.RelatedEvents = append(s.RelatedEvents, relatedEvent)
	}

	// Data-Source-specific fields
	if st == SubjectDataSource {
		var source, sourceLink, sourceObjective, sourceDataSize string

		err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Source").
					Value(&source),
				huh.NewInput().
					Title("Link").
					Value(&sourceLink),
				huh.NewInput().
					Title("Objective").
					Value(&sourceObjective),
				huh.NewInput().
					Title("Data Size").
					Value(&sourceDataSize),
			),
		).Run()
		if err != nil {
			return err
		}

		s.Source = source
		s.SourceLink = sourceLink
		s.SourceObjective = sourceObjective
		s.SourceDataSize = sourceDataSize
	}

	return nil
}
