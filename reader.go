package main

import (
	"bufio"
	"errors"
	"github.com/samber/lo"
	"log"
	"os"
	"strings"
)

func ReadInput(path string) (emails []string, err error) {

	file, err := os.Open(path)

	if err != nil {
		return
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			log.Printf("can not close input file: %v", closeErr)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		emails = append(emails, strings.ToLower(strings.TrimSpace(scanner.Text())))
	}

	if err = scanner.Err(); err != nil {
		return
	}

	if len(emails) == 0 {
		err = errors.New("Input file is empty")
	}

	emails = lo.Uniq[string](emails)

	return
}
