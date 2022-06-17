package env

import (
	"jira-server-api/main/pkg/util/logging"
	"log"
	"os"
	"strings"
)

func Get(name string) string {
	data, err := os.ReadFile(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	content := string(data)

	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		keyValuePair := strings.SplitN(lines[i], "=", 2)
		if len(keyValuePair) == 2 {
			key := keyValuePair[0]
			value := keyValuePair[1]
			logging.D("env.Get - key: " + key + " value: " + value)

			if name == key {
				logging.D("env.Get - resolve '" + name + "' to the value: " + value)
				return value
			}
		}
	}

	logging.D("env.Get - could not resolve '" + name + "' to any value")
	return ""
}
