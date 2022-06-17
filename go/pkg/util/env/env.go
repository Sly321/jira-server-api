package env

import (
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
		keyValuePair := strings.Split(lines[i], "=")
		if len(keyValuePair) == 2 {
			key := keyValuePair[0]
			value := keyValuePair[1]

			if name == key {
				return value
			}
		}
	}

	return ""
}
