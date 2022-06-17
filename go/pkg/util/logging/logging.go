package logging

import (
	"fmt"
)

var LEVEL = "DEBUG"

func D(msg string) error {
	if LEVEL == "DEBUG" {
		fmt.Println("DEBUG: " + msg)
	}
	return nil
}
