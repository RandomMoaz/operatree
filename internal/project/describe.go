package project

import (
	"fmt"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/ui"
)

func describeProject(p *Project) {
	var b strings.Builder

	// Header
	b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + p.Name + ui.AnsiReset)
	b.WriteString(ui.AnsiGray + "  ·  project" + ui.AnsiReset + "\n")
	b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n\n")

	// Core fields
	b.WriteString(label("BaseDir") + "  " + value(p.BaseDir) + "\n")
	if len(p.Tags) > 0 {
		tags := make([]string, len(p.Tags))
		for i, t := range p.Tags {
			tags[i] = ui.AnsiYellow + "#" + t + ui.AnsiReset
		}
		b.WriteString(label("Tags") + "     " + strings.Join(tags, "  ") + "\n")
	}

	// Modules
	b.WriteString("\n" + ui.AnsiBold + ui.AnsiPurple + "Modules" + ui.AnsiReset + "\n")
	b.WriteString(ui.AnsiGray + strings.Repeat("─", 40) + ui.AnsiReset + "\n")
	for _, m := range p.Modules {
		describeModule(&b, m, 0)
	}

	b.WriteString("\n")
	fmt.Print(b.String())
}

func describeModule(b *strings.Builder, m module.Module, depth int) {
	indent := strings.Repeat("  ", depth)
	connector := "├─ "
	if depth == 0 {
		connector = ""
	}

	// Module line
	b.WriteString(indent + connector +
		ui.AnsiBold + ui.AnsiPurple + string(m.Type) + ui.AnsiReset +
		ui.AnsiGray + "  " + m.Name + ui.AnsiReset + "\n")

	// AbsPath
	b.WriteString(indent + "   " + ui.AnsiGray + m.AbsPath + ui.AnsiReset + "\n")

	// Subjects summary
	if len(m.Subjects) > 0 {
		b.WriteString(indent + "   " +
			ui.AnsiYellow + fmt.Sprintf("%d subject(s)", len(m.Subjects)) + ui.AnsiReset + "\n")
		for _, s := range m.Subjects {
			b.WriteString(indent + "   " +
				ui.AnsiGray + "├─ " + ui.AnsiReset +
				value(s.Name) +
				dim("  ["+string(s.Type)+"]") + "\n")
		}
	}

	// SubDirs
	if len(m.SubDirs) > 0 {
		for _, sd := range m.SubDirs {
			b.WriteString(indent + "   " +
				ui.AnsiGray + "├─ " + sd + ui.AnsiReset + "\n")
		}
	}

	// Recurse into sub-modules
	for _, sub := range m.Modules {
		describeModule(b, sub, depth+1)
	}
}
