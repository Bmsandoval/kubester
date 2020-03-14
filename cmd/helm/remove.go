package helm

import (
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/services/kube_svc"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"re"},
	Short:   "remove environment",
	Long:    ``,
	Run:     Remove,
}

func Remove(_ *cobra.Command, _ []string) {
	release, err := kube_svc.SelectRelease()
	if err != nil {
		panic(err)
	}

	if err := utils.Exec(bash.HelmDelete(*release)); err != nil {
		panic(err)
	}
}

func init() {
	HelmCmds.AddCommand(RemoveCmd)
}
