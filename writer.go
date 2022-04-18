package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	lock = sync.RWMutex{}
)

type Writer struct {
	lock sync.RWMutex
	file *os.File
}

func NewWriter() (w *Writer, err error) {

	filename := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "-")
	file, err := os.OpenFile(filename+".csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	return &Writer{
		file: file,
		lock: sync.RWMutex{},
	}, err
}

func (w *Writer) Write(p []byte) (n int, err error) {
	lock.Lock()
	defer lock.Unlock()
	return w.file.Write(p)
}

func (w *Writer) Close() {
	if err := w.file.Close(); err != nil {
		log.Printf("can not close file: %v", err)
	}
}
