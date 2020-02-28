package bash

import "fmt"

func HelmInstall(app string) string { return fmt.Sprintf("helm upgrade --install dev-%s", app) }
func HelmDelete(app string) string  { return fmt.Sprintf("helm delete dev-%s", app) }
func HelmList() string              { return "helm ls -a -o json" }

type HelmListObj struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Revision   string `json:"revision"`
	Updated    string `json:"updated"`
	Status     string `json:"status"`
	Chart      string `json:"chart"`
	AppVersion string `json:"app_version"`
}
