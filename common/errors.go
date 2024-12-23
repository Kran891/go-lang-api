package main

import "github.com/kran891/go-lang-api/logger"

const (
	WORKFLOW_NOT_FOUND = "Hello world %s"
)

func main() {
	logger.Errorf("Hello %s", "Kran")
}
