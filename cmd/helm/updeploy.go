package helm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/config"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var DeployCmd = &cobra.Command{
	Use:   "updeploy",
	Aliases: []string{"de"},
	Short: "deploy environment",
	Long: ``,
	Run: Deploy,
}

type deploymentInformation struct {
	Name string
	FilePath string
}


func Deploy(_ *cobra.Command, _ []string) {
	ctxt := "dev"
	auxiliaryReleasables, err := GetReleasableAuxiliaryDeployments(ctxt)
	if err != nil { panic(err) }

	for _, r := range auxiliaryReleasables {
		if err := utils.Exec(bash.HelmInstallAuxiliary(ctxt, r.Name, r.FilePath)); err != nil {
			panic(err)
		}
	}

	primaryReleasables, err := GetReleasablePrimaryDeployments()
	if err != nil { panic(err) }

	for _, r := range primaryReleasables {
		if err := utils.Exec(bash.HelmInstall(ctxt, r.Name, r.FilePath)); err != nil {
			panic(err)
		}
	}
}

func GetReleasablePrimaryDeployments() ([]deploymentInformation, error) {
	var configs = config.GetConfigFromViper()
	deploymentPath := fmt.Sprintf("%s/deployments/submodules",configs.KubesterConfig.ProjectFilePath)

	// List all releasable items
	files, err := ioutil.ReadDir(deploymentPath)
	if err != nil {
		return nil, err
	}

	var releasables []deploymentInformation
	for _, f := range files {
		releasables = append(releasables, deploymentInformation{
			Name:     fmt.Sprintf("%s", f.Name()),
			FilePath: fmt.Sprintf("%s/%s", deploymentPath, f.Name()),
		})
	}

	return releasables, nil
}

func GetReleasableAuxiliaryDeployments(ctxt string) ([]deploymentInformation, error) {
	// List all released items
	err, out, errout := utils.ExecGetOutput(bash.HelmList())
	if err != nil { return nil, err }
	if errout != "" { return nil, errors.New(errout) }
	// Get releases in a well defined format
	var releases []bash.HelmListObj
	json.Unmarshal([]byte(out), &releases)

	// List all releasable items
	var configs = config.GetConfigFromViper()
	releasePath := fmt.Sprintf("%s/deployments",configs.KubesterConfig.ProjectFilePath)

	files, err := ioutil.ReadDir(releasePath)
	if err != nil {
		return nil, err
	}

	var releasables []deploymentInformation
	for _, f := range files {
		// skip submodules directory
		if f.Name() == "submodules" {
			continue
		}
		// collect any releases that aren't already released
		for _, r := range releases {
			releasableName := fmt.Sprintf("%s-%s", ctxt, f.Name())
			if r.Name ==  releasableName {
				// if auxiliary already released, ignore it
				continue
			}
			releasables = append(releasables, deploymentInformation{
				Name:     f.Name(),
				FilePath: fmt.Sprintf("%s/%s", releasePath, f.Name()),
			})
		}
	}
	return releasables, nil
}

func init() {
	HelmCmds.AddCommand(DeployCmd)
}
