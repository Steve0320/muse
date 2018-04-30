package helpers

import "log"

func CheckError(e error, logger *log.Logger) {
	if e != nil {
		logger.Fatal(e)
	}
}
