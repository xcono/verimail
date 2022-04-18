package main

import (
	"log"
	"os"
)

func Logger() (f *os.File, err error) {

	f, err = os.OpenFile("errors.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	log.SetOutput(f)

	return
}
