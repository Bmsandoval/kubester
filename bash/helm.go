package bash

import "fmt"

func HelmInstallAuxiliary(context string, appName string, appPath string) string { return fmt.Sprintf("helm upgrade --install %s-%s %s/", context, appName, appPath) }
func HelmInstall(context string, appName string, appPath string) string { return fmt.Sprintf("helm upgrade --install %s-%s %s/chart/%s", context, appName, appPath, appName) }
func HelmDelete(app string) string  { return fmt.Sprintf("helm delete %s", app) }
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
