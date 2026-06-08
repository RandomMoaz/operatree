// internal/project/summary.go
package project

import (
	"fmt"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/ui"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

type subjectStats struct {
	total         int
	byType        map[subject.SubjectType]int
	byStatus      map[string]int
	recentDate    string // latest date found across all subjects
	recentSubject string // latest subject by latest date found across all subjects
}

func (p *Project) Summary() {
	stats := collectStats(p)
	printSummary(p, stats)
}

func collectStats(p *Project) subjectStats {
	stats := subjectStats{
		byType:   make(map[subject.SubjectType]int),
		byStatus: make(map[string]int),
	}

	db := BuildSearchDB(p)
	for _, entry := range db {
		s := entry.Subject
		stats.total++
		stats.byType[s.Type]++

		if s.Status != "" {
			stats.byStatus[s.Status]++
		}

		// track most recent date
		if s.Date > stats.recentDate {
			stats.recentDate = s.Date
			stats.recentSubject = s.Name
		}
	}

	return stats
}

func printSummary(p *Project, stats subjectStats) {
	var b strings.Builder

	// Header
	b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + p.Name + ui.AnsiReset)
	b.WriteString(ui.AnsiGray + "  ·  summary" + ui.AnsiReset + "\n")
	b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n\n")

	// Totals
	b.WriteString(label("Total Subjects") + "  " + ui.AnsiYellow + fmt.Sprintf("%d", stats.total) + ui.AnsiReset + "\n")
	if stats.recentDate != "" {
		b.WriteString(label("Latest Activity") + "  " + value(stats.recentDate) + "\n")
		b.WriteString(label("Latest Subject") + "  " + value(stats.recentSubject) + "\n")
	}

	// By type
	if len(stats.byType) > 0 {
		b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + "By Type" + ui.AnsiReset + "\n")
		b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n")
		for t, count := range stats.byType {
			bar := renderBar(count, stats.total, 20)
			b.WriteString(fmt.Sprintf("  %-14s %s%s%s  %s%d%s\n",
				string(t),
				ui.AnsiPurple, bar, ui.AnsiReset,
				ui.AnsiYellow, count, ui.AnsiReset,
			))
		}
	}

	// By status (only if any subjects have status)
	if len(stats.byStatus) > 0 {
		b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + "By Status" + ui.AnsiReset + "\n")
		b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n")
		for st, count := range stats.byStatus {
			color := statusColor(st)
			bar := renderBar(count, stats.total, 20)
			b.WriteString(fmt.Sprintf("  %-14s %s%s%s  %s%d%s\n",
				st,
				color, bar, ui.AnsiReset,
				ui.AnsiYellow, count, ui.AnsiReset,
			))
		}
	}

	// Module breakdown
	b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + "Modules" + ui.AnsiReset + "\n")
	b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n")
	for _, m := range p.Modules {
		count := countModuleSubjects(m)
		if count == 0 {
			continue // skip empty modules
		}
		b.WriteString(fmt.Sprintf("  %-24s %s%d subject(s)%s\n",
			m.Name,
			ui.AnsiYellow, count, ui.AnsiReset,
		))
		// sub-modules with subjects
		for _, sub := range m.Modules {
			subCount := countModuleSubjects(sub)
			if subCount == 0 {
				continue
			}
			b.WriteString(fmt.Sprintf("    %-22s %s%d%s\n",
				"↳ "+sub.Name,
				ui.AnsiGray, subCount, ui.AnsiReset,
			))
		}
	}

	b.WriteString("\n")
	fmt.Print(b.String())
}

func countModuleSubjects(m module.Module) int {
	count := len(m.Subjects)
	for _, sub := range m.Modules {
		count += countModuleSubjects(sub)
	}
	return count
}

func renderBar(count, total, width int) string {
	if total == 0 {
		return strings.Repeat("░", width)
	}
	filled := (count * width) / total
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func statusColor(status string) string {
	switch strings.ToLower(status) {
	case "done":
		return ui.AnsiGreen
	case "in-progress":
		return ui.AnsiYellow
	case "pending", "planned":
		return ui.AnsiGray
	case "postponed":
		return "\033[38;5;203m" // orange-red
	default:
		return ui.AnsiGray
	}
}
