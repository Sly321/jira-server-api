package main

import (
	"encoding/json"
	"fmt"
	"jira-server-api/main/pkg/client/issue"
	"jira-server-api/main/pkg/util/env"
)

func main() {
	issue := issue.Get("D02449-206")
	fmt.Println(env.Get("JIRA_HOST"))
	fmt.Println(PrettyPrint(issue.Fields.Status))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
