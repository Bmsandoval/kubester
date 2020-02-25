package git

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var StashAllCmd = &cobra.Command{
	Use:   "stash_all",
	Aliases: []string{"sa"},
	Short: "stash current branch and submodules",
	Long: ``,
	Run: StashAll,
}


func StashAll(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitStash()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	command = bash.GitStashRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmd.AddCommand(StashAllCmd)
}
