package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"github.com/spf13/cobra"
)

// var subjectName string
// var subjectDate string
var validSubjects []cobra.Completion
var addCmd = &cobra.Command{}

var ns subject.Subject

func init() {
	// build completion slice from available subjects dynamically
	for k := range project.SubjectModuleMap {
		sn := strings.ToLower(string(k))
		validSubjects = append(validSubjects, sn)
	}

	// define command
	addCmd = &cobra.Command{
		Use:   fmt.Sprintf("add [%s]", strings.Join(validSubjects, " | ")),
		Short: "Add a new subject",
		Long: `Launch an interactive form to add a new subject to the project, or create one directly using flags for scripting.

The subject type is required and determines which fields are collected.
Providing --name skips the interactive form entirely and creates the subject
immediately with the supplied flag values. Every creation is appended to activity.log.

Interactive mode:
  operatree add event                    # fully interactive
  operatree add task                     # fully interactive

Non-interactive mode — provide at minimum --name:
  operatree add task --name "Prepare Report"
  operatree add event --name "Site Visit" --date 2026-06-01 --location Cairo

Examples:
  operatree add event
  operatree add event --name "Cairo Site Visit" --date 2026-06-01 --location Cairo --participants "Alex,Sara" --tags "site,inspection"
  operatree add task --name "Prepare Report" --owner Alex --status active --related-events "Cairo Site Visit"
  operatree add topic --name "Predictive Maintenance" --related-objective "Reduce Downtime" --tags "ml,iot"
  operatree add objective --name "Reduce Downtime" --status active --tags "maintenance,kpi"
	operatree add datasource --name "Sensor Readings 2025" --source "IoT Team" --source-link "https://kaggle.com/path/to/source" --source-datasize "2.4GB"
  operatree add task --name "Deploy" --date 2026-06-15 -d /path/to/project`,
		ValidArgs: validSubjects,
		Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
		Run:       newSubject,
	}

	addCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	addCmd.Flags().StringVar(&ns.Name, "name", "", "subject name")
	addCmd.Flags().StringVar(&ns.Date, "date", "", "subject date")
	addCmd.Flags().StringVar(&ns.Notes, "notes", "", "subject notes")
	addCmd.Flags().StringSliceVar(&ns.Tags, "tags", []string{}, "subject tags, comma delimited")
	addCmd.Flags().StringSliceVar(&ns.Participants, "participants", []string{}, "subject participants, comma delimited")
	addCmd.Flags().StringVar(&ns.Location, "location", "", "subject location, valid for Events only")
	addCmd.Flags().StringVar(&ns.Owner, "owner", "", "subject owner")
	addCmd.Flags().StringVar(&ns.Status, "status", "", "subject status")
	addCmd.Flags().StringVar(&ns.RelatedObjective, "related-objective", "", "subject related objective")
	addCmd.Flags().StringSliceVar(&ns.RelatedEvents, "related-events", []string{}, "subject related events, comma delimited")
	addCmd.Flags().StringSliceVar(&ns.Outputs, "outputs", []string{}, "subject outputs, comma delimited")
	addCmd.Flags().StringVar(&ns.Source, "source", "", "datasource origin (e.g. Kaggle, internal team, API)")
	addCmd.Flags().StringVar(&ns.SourceLink, "source-link", "", "datasource URL or path to source data")
	addCmd.Flags().StringVar(&ns.SourceObjective, "source-objective", "", "datasource related objective")
	addCmd.Flags().StringVar(&ns.SourceDataSize, "source-datasize", "", "datasource size or volume")

	addCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(addCmd)
}

func newSubject(cmd *cobra.Command, args []string) {
	a := args[0]

	// safety check -> dynamic loading prevents reach this error
	if a == "" {
		log.Fatal("unsupprted subject")
	}

	st := strings.ToUpper(a)

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := project.NewSubject(&p, ns, subject.SubjectType(st)); err != nil {
		log.Fatal(err)
	}
}
