package main

import (
	"encoding/json"
	"fmt"
	"jira-server-api/main/pkg/client/issue"
)

func main() {
	issue := issue.Get("D602449-206")
	fmt.Println(PrettyPrint(issue.Fields.Status))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
