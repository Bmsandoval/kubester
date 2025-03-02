package cmd

import (
	"fmt"
	"github.com/bmsandoval/kubester/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var ExperimentCmd = &cobra.Command{
	Use:     "experiment",
	Aliases: []string{"exp"},
	Short:   "a command endpoint to experiment with stuff",
	Long:    ``,
	Run:     Experiment,
}

func Experiment(_ *cobra.Command, _ []string) {
	var configs = config.GetConfigFromViper()

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/deployments/submodules/", configs.KubesterConfig.ProjectFilePath))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func init() {
	rootCmd.AddCommand(ExperimentCmd)
}
