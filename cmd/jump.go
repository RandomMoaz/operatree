package cmd

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/runner"
	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(jumpCmd)
}

var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "Opens project dir in default file manager",
	Long:  "Opens a tracked project dir using fuzzy-finder into default file manager",
	Run:   jump,
}

func jump(cmd *cobra.Command, args []string) {

	c, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load config file. Use operatree init to initialize config"))
	}

	cp, err := config.Find(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := runner.OpenFileManager(cp.AbsPath); err != nil {
		log.Fatal(err)
	}
}
