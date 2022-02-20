package errors

import (
	"os"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
)

func Handler(title string, e error) bool {
	if e != nil {
		// if we're on debugging mode
		// log the error
		if os.Getenv("APP_DEBUG") == "true" {
			logger.Error(title, e)
		}
		return true
	}
	return false
}
