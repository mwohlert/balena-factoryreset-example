package utils

import (
	"balena-factorreset-example/constants"
	"bytes"
	"encoding/json"
	"net/http"
)

// ForceRequest : Body of a request using force
type ForceRequest struct {
	Force bool `json:"force"`
}

var appID = GetEnv(constants.BalenaAppIDEnv, constants.BalenaAppID)

// PurgeData purges application Data for a specific application
func PurgeData() error {
	supervisorURI, err := GetSupervisorURI()
	if err != nil {
		return err
	}

	supervisorAPIKey, err := GetSupervisorAPIKey()
	if err != nil {
		return err
	}

	purgeRequest := ForceRequest{Force: true}

	jsonBody, err := json.Marshal(purgeRequest)
	if err != nil {
		return err
	}

	go func() {
		_, err = http.Post(supervisorURI+"/v2/applications/"+appID+"/purge?apikey="+supervisorAPIKey,
			"application/json", bytes.NewBuffer(jsonBody))
	}()

	return nil
}
