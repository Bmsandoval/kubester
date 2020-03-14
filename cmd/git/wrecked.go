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

var WreckCmd = &cobra.Command{
	Use:     "wreck",
	Aliases: []string{"w"},
	Short:   "runs the `Reset` and `Clean` command",
	Long:    `Resets all tracked files and removes all untracked files`,
	Run:     Wreck,
}

func Wreck(_ *cobra.Command, _ []string) {
	kube_svc.UserConfirms("Are you sure you want to do this? This will reset all tracked files in project. Recommend first stashing changes")
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	// RESET
	command := bash.GitResetHard()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
	command = bash.GitResetHardRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	// CLEAN
	command = bash.GitClean()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
	command = bash.GitCleanRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	// SWITCH TO MASTER
	command = bash.GitBranch()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
	command = bash.GitBranchRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	// FETCH AND PULL
	command = bash.GitFetchRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
	command = bash.GitPullRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmds.AddCommand(WreckCmd)
}
