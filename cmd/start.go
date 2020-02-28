package cmd

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Aliases: []string{"s"},
	Short: "start minikube",
	Long: ``,
	Run: StartMinikube,
}


func StartMinikube(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("minikube"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.MinikubeStart()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(StartCmd)
}
