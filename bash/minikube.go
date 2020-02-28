package bash

import (
	"fmt"
	"github.com/bmsandoval/kubester/config"
)


func MinikubeStart() string {
	var configs = config.GetConfigFromViper()
	return fmt.Sprintf("minikube start --mount-string %s:/data --mount --cpus 4 --memory 8192", configs.KubesterConfig.ProjectFilePath)
}
func MinikubeStop() string { return "minikube stop" }
