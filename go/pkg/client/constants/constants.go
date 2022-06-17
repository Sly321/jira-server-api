package constants

import "jira-server-api/main/pkg/util/env"

var (
	JIRA_HOST           = env.Get("JIRA_HOST")
	JIRA_BASIC_AUTH     = env.Get("JIRA_BASIC_AUTH")
	JIRA_REST_URL       = JIRA_HOST + "/rest/api/2"
	JIRA_REST_AGILE_URL = JIRA_HOST + "/rest/agile/1.0"
)
