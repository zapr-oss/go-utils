package common_utils

import (
	"io"
	"log"
)

func CloseStream(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Println("Error while closing object: ", err)
	}
}
