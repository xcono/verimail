package main

import (
	"flag"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"time"
)

var (
	verifier = emailverifier.NewVerifier()
)

func main() {

	input := flag.String("i", "input.txt", "specify file path to source of emails")
	workers := flag.Int("w", 2, "specify max workers count")
	connections := flag.Int("c", 2, "specify max connections count")
	smtp := flag.Bool("smtp", false, "specify if smtp check required")

	if *smtp {
		verifier.EnableSMTPCheck()
	}

	// logger
	f, err := Logger()
	if err != nil {
		_ = fmt.Errorf("can not open log file: %v", err)
		os.Exit(1)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("can not close log file: %v", err)
		}
	}()

	// reader
	emails, err := ReadInput(*input)
	if err != nil {
		_ = fmt.Errorf("can not open input file: %v", err)
		os.Exit(1)
	}

	// writer
	w, err := NewWriter()
	if err != nil {
		_ = fmt.Errorf("can not open output file: %v", err)
		os.Exit(1)
	}
	defer w.Close()

	// progress bar
	bar := progressbar.Default(int64(len(emails)), "email verification")
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			_ = bar.Set(counter)
		}
	}()
	defer func() {
		ticker.Stop()
		err = bar.Finish()
		if err != nil {
			_ = fmt.Errorf("progress bar error: %v", err)
		}
	}()

	// run verification
	NewPool(func(s string) error {
		return Verify(s, *smtp)
	}, w, *workers, *connections).Start(emails)
}
