package git

import (
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
)

var OnboardCmd = &cobra.Command{
	Use:     "onboard",
	Aliases: []string{"on"},
	Short:   "Onboard git stuff",
	Long:    `basically just inits the submodules`,
	Run:     Onboard,
}

func Onboard(_ *cobra.Command, _ []string) {
	command := bash.GitOnboard()
	if err := utils.Exec(command); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	GitCmds.AddCommand(OnboardCmd)
}
