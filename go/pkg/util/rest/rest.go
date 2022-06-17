package rest

import (
	"jira-server-api/main/pkg/client/constants"
	"jira-server-api/main/pkg/util/logging"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Basic "+constants.JIRA_BASIC_AUTH)
	logging.D("rest.Get - requesting: " + url + " with header: Authorization: Basic " + constants.JIRA_BASIC_AUTH)
	return client.Do(req)
}
