package rest

import (
	"jira-server-api/main/pkg/client/constants"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Basic "+constants.JIRA_BASIC_AUTH)
	return client.Do(req)
}
