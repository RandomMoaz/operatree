package cmd

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/charmbracelet/glamour"
	"github.com/hanymamdouh82/operatree/internal/help"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(explainCmd)
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Prints operatree dir philosophy",
	Long:  "Explains full documentation for OperaTree dir usage philosophy",
	Run:   explain,
}

func explain(cmd *cobra.Command, args []string) {

	data, _ := help.FS.ReadFile("dir_struct.md")

	out, err := glamour.Render(string(data), "dark")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}
