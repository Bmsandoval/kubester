package git

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var FetchPullCmd = &cobra.Command{
	Use:   "fetch_pull",
	Aliases: []string{"fp"},
	Short: "fetch and pull current branch and submodules",
	Long: `git fetch --recurse-submodules && git pull --recurse-submodules`,
	Run: FetchPull,
}


func FetchPull(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitFetchRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}

	command = bash.GitPullRecurse()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmds.AddCommand(FetchPullCmd)
}
