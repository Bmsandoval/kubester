package cmd

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Aliases: []string{"s"},
	Short: "stop minikube",
	Long: ``,
	Run: StopMinikube,
}


func StopMinikube(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("minikube"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.MinikubeStop()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(StopCmd)
}
