package subject

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	primary   = lipgloss.Color("#7B61FF")
	subtle    = lipgloss.Color("#888888")
	highlight = lipgloss.Color("#F1C40F")
	white     = lipgloss.Color("#FFFFFF")
	green     = lipgloss.Color("#2ECC71")

	typeBadgeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primary)

	nameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(white)

	dividerStyle = lipgloss.NewStyle().
			Foreground(subtle)

	labelStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Width(16)

	valueStyle = lipgloss.NewStyle().
			Foreground(white)

	tagStyle = lipgloss.NewStyle().
			Foreground(highlight)

	notesLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(subtle)

	notesValueStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Italic(true).
			PaddingLeft(1)

	statusStyle = map[string]lipgloss.Style{
		"done":        lipgloss.NewStyle().Bold(true).Foreground(green),
		"in-progress": lipgloss.NewStyle().Bold(true).Foreground(highlight),
		"pending":     lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#E74C3C")),
	}
)

func row(label, value string) string {
	if strings.TrimSpace(value) == "" {
		return ""
	}
	return lipgloss.JoinHorizontal(lipgloss.Top,
		labelStyle.Render(label),
		valueStyle.Render(value),
	) + "\n"
}

func rowList(label string, items []string) string {
	// filter empty
	var clean []string
	for _, v := range items {
		if strings.TrimSpace(v) != "" {
			clean = append(clean, v)
		}
	}
	if len(clean) == 0 {
		return ""
	}
	return row(label, strings.Join(clean, "  ·  "))
}

func renderTags(tags []string) string {
	var clean []string
	for _, t := range tags {
		if strings.TrimSpace(t) != "" {
			clean = append(clean, tagStyle.Render("#"+strings.TrimSpace(t)))
		}
	}
	if len(clean) == 0 {
		return ""
	}
	return lipgloss.JoinHorizontal(lipgloss.Top,
		labelStyle.Render("Tags"),
		strings.Join(clean, "  "),
	) + "\n"
}

func renderStatus(status string) string {
	if strings.TrimSpace(status) == "" {
		return ""
	}
	s, ok := statusStyle[strings.ToLower(status)]
	if !ok {
		s = lipgloss.NewStyle().Foreground(subtle)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top,
		labelStyle.Render("Status"),
		s.Render(strings.ToUpper(status)),
	) + "\n"
}

func describe(s *Subject) {
	var b strings.Builder

	// Header
	badge := typeBadgeStyle.Render(strings.ToUpper(string(s.Type)))
	name := nameStyle.Render(s.Name)
	b.WriteString("\n " + badge + dividerStyle.Render("  ·  ") + name + "\n")
	b.WriteString(dividerStyle.Render(" "+strings.Repeat("─", 40)) + "\n\n")

	// Core fields
	b.WriteString(row("Path", s.DirName))
	b.WriteString(row("Date", s.Date))
	b.WriteString(row("Location", s.Location))
	b.WriteString(row("Owner", s.Owner))
	b.WriteString(renderStatus(s.Status))
	b.WriteString(rowList("Participants", s.Participants))
	b.WriteString(renderTags(s.Tags))

	// Notes
	if strings.TrimSpace(s.Notes) != "" {
		b.WriteString("\n")
		b.WriteString(" " + notesLabelStyle.Render("Notes") + "\n")
		b.WriteString(dividerStyle.Render(" "+strings.Repeat("─", 10)) + "\n")
		b.WriteString(notesValueStyle.Render(s.Notes) + "\n")
	}

	// Relations
	hasRelations := s.RelatedObjective != "" || len(s.RelatedEvents) > 0 || len(s.Outputs) > 0
	if hasRelations {
		b.WriteString("\n")
		b.WriteString(row("Objective", s.RelatedObjective))
		b.WriteString(rowList("Related Events", s.RelatedEvents))
		b.WriteString(rowList("Outputs", s.Outputs))
	}

	b.WriteString("\n")
	fmt.Println(b.String())
}
