package kube_svc

import (
	"encoding/json"
	"errors"
	"github.com/bmsandoval/kubester/bash"
	"github.com/bmsandoval/kubester/utils"
	"github.com/ktr0731/go-fuzzyfinder"
)

func SelectRelease() (*string, error) {
	// List all released items
	err, out, errout := utils.ExecGetOutput(bash.HelmList())
	if err != nil {
		return nil, err
	}
	if errout != "" {
		return nil, errors.New(errout)
	}
	// Get releases in a well defined format
	var releases []bash.HelmListObj
	json.Unmarshal([]byte(out), &releases)

	var releaseNames []string
	for _, r := range releases {
		releaseNames = append(releaseNames, r.Name)
	}

	selected, err := fuzzyfinder.Find(releaseNames,
		func(i int) string {
			return releaseNames[i]
		})
	if err != nil {
		return nil, errors.New("no release selected, aborting")
	}
	return &releaseNames[selected], nil
}
