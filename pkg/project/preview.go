package project

import (
	"strings"

	"github.com/hanymamdouh82/operatree/internal/ui"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func label(s string) string {
	return ui.AnsiPurple + ui.AnsiBold + s + ui.AnsiReset
}

func value(s string) string {
	return ui.AnsiReset + s
}

func dim(s string) string {
	return ui.AnsiGray + ui.AnsiItalic + s + ui.AnsiReset
}

func formatPreview(s subject.Subject) string {
	var b strings.Builder

	// Header
	b.WriteString(ui.AnsiBold + ui.AnsiPurple + s.Name + ui.AnsiReset)
	b.WriteString(ui.AnsiGray + "  ·  " + string(s.Type) + ui.AnsiReset + "\n")
	b.WriteString(ui.AnsiGray + strings.Repeat("─", 35) + ui.AnsiReset + "\n\n")

	// Core fields
	if s.Date != "" {
		b.WriteString(label("Date      ") + value(s.Date) + "\n")
	}
	if s.Location != "" {
		b.WriteString(label("Location  ") + value(s.Location) + "\n")
	}
	if s.Owner != "" {
		b.WriteString(label("Owner     ") + value(s.Owner) + "\n")
	}
	if s.Status != "" {
		b.WriteString(label("Status    ") + value(s.Status) + "\n")
	}
	if len(s.Participants) > 0 {
		b.WriteString(label("People    ") + value(strings.Join(s.Participants, ", ")) + "\n")
	}

	// Tags
	if len(s.Tags) > 0 {
		tags := make([]string, len(s.Tags))
		for i, t := range s.Tags {
			tags[i] = ui.AnsiYellow + "#" + t + ui.AnsiReset
		}
		b.WriteString(label("Tags      ") + strings.Join(tags, "  ") + "\n")
	}

	// Notes
	if s.Notes != "" {
		b.WriteString("\n" + label("Notes") + "\n")
		b.WriteString(dim(s.Notes) + "\n")
	}

	// Relations
	hasRelations := s.RelatedObjective != "" || len(s.RelatedEvents) > 0 || len(s.Outputs) > 0
	if hasRelations {
		b.WriteString("\n" + ui.AnsiGray + strings.Repeat("─", 35) + ui.AnsiReset + "\n")
		if s.RelatedObjective != "" {
			b.WriteString(label("Objective ") + value(s.RelatedObjective) + "\n")
		}
		if len(s.RelatedEvents) > 0 {
			b.WriteString(label("Events    ") + value(strings.Join(s.RelatedEvents, ", ")) + "\n")
		}
		if len(s.Outputs) > 0 {
			b.WriteString(label("Outputs   ") + value(strings.Join(s.Outputs, ", ")) + "\n")
		}
	}

	return b.String()
}
