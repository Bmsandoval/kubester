package kube

import (
	"github.com/spf13/cobra"
)

var KubeCmds = &cobra.Command{
	Use:   "kube",
	Aliases: []string{"k"},
	Short: "kubectl ...",
	Run: Kube,
}


func Kube(_ *cobra.Command, _ []string) { }

//TODO: reset everything including database. AKA - kubewreck
// this is required only if you want to delete your database completely
//$ helm list --short | xargs -L1 helm delete
//$ kubectl delete pvc --all
//$ minikube stop
//$ minikube start
