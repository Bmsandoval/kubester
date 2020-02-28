package helm

import (
	"github.com/spf13/cobra"
)

var HelmCmds = &cobra.Command{
	Use:   "helm",
	Aliases: []string{"h"},
	Short: "helm command start",
	Run: Helm,
}


func Helm(_ *cobra.Command, _ []string) { }
