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

var CleanAllCmd = &cobra.Command{
	Use:   "clean_all",
	Aliases: []string{"ca"},
	Short: "clean current branch and submodules",
	Long: `Removes untracked files and directories in base repo and all sub repos`,
	Run: CleanAll,
}


func CleanAll(_ *cobra.Command, _ []string) {
	kube_svc.UserConfirms("Are you sure you want to do this? This will remove all untracked files. Recommend stashing first")
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitClean()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	command = bash.GitCleanRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmds.AddCommand(CleanAllCmd)
}
