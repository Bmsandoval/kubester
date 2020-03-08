package kube

import (
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/services/kube_svc"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
)

var ContextCmd = &cobra.Command{
	Use:   "context",
	Aliases: []string{"ctx"},
	Short: "kubectl config use-context",
	Long: ``,
	Run: Context,
}

func Context(_ *cobra.Command, _ []string) {
	release, err := kube_svc.SelectRelease()
	if err != nil {
		panic(err)
	}

	if err := utils.Exec(bash.HelmDelete(*release)); err != nil {
		panic(err)
	}
}

func init() {
	KubeCmds.AddCommand(ContextCmd)
}
