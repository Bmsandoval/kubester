package git

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	kube_svc "github.com/bmsandoval/kubester/services/input_svc"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var ResetAllCmd = &cobra.Command{
	Use:     "reset_all",
	Aliases: []string{"ra"},
	Short:   "reset current branch and submodules",
	Long:    `Resets all tracked files in base repo and all sub repos to last commit`,
	Run:     ResetAll,
}

func ResetAll(_ *cobra.Command, _ []string) {
	kube_svc.UserConfirms("Are you sure you want to do this? This will reset all tracked files in project. Recommend first stashing changes")
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitResetHard()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	command = bash.GitResetHardRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmds.AddCommand(ResetAllCmd)
}
