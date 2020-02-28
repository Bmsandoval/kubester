package helm

import (
	"encoding/json"
	"fmt"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/config"
	"github.com/bmsandoval/kubester/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var DeployCommand = &cobra.Command{
	Use:   "updeploy",
	Aliases: []string{"de"},
	Short: "deploy environment",
	Long: ``,
	Run: Deploy,
}


func Deploy(_ *cobra.Command, _ []string) {
	auxiliaryReleasables := GetReleasableAuxiliaryDeployments()
	if len(auxiliaryReleasables) < 1 { return }

	log.Printf("%+v", auxiliaryReleasables)
}


func GetReleasableAuxiliaryDeployments() []string {
	// List all released items
	err, out, errout := utils.ExecGetOutput(bash.HelmList())
	if err != nil { panic(err) }
	if errout != "" { panic(errout) }
	// Get releases in a well defined format
	var releases []bash.HelmListObj
	json.Unmarshal([]byte(out), &releases)

	// List all releasable items
	var configs = config.GetConfigFromViper()
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/deployments/",configs.KubesterConfig.ProjectFilePath))
	if err != nil {
		log.Fatal(err)
	}

	var releasables []string
	for _, f := range files {
		// skip submodules directory
		if f.Name() == "submodules" {
			continue
		}
		// collect any releases that aren't already released
		for _, r := range releases {
			releasableName := fmt.Sprintf("dev-%s", f.Name())
			if r.Name ==  releasableName {
				// if auxiliary already released, ignore it
				continue
			}
			releasables = append(releasables, releasableName)
		}
	}
	return releasables
}

func init() {
	HelmCmds.AddCommand(DeployCommand)
}
