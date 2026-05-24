package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:       "show [tracked | config]",
	Short:     "Show information about operatree",
	Long:      "Shows information about operatree",
	ValidArgs: []cobra.Completion{"tracked", "config"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:       show,
}

func show(cmd *cobra.Command, args []string) {
	// Load config
	c, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load config file. Use operatree init to initialize config"))
	}

	switch args[0] {
	case "tracked":
		c.ListProjects()
		return
	case "config":
		b, err := yaml.Marshal(c)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("%s\n", b)
	default:
		log.Fatal("unknow command")
	}

}
