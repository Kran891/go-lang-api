package main

import logger "github.com/kran891/go-log"

const (
	WORKFLOW_NOT_FOUND = "Hello world %s"
)

func main() {
	logger.Errorf("Hello %s", "Kran")
}
