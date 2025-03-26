package errors

import "log"

func LogError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}
