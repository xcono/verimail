package main

import (
	"io"
	"log"
	"sync"
)

var counter int

type Verifier func(string) error

type Pool struct {
	jobs           chan string
	results        chan Result
	maxWorkers     int
	maxConnections int
	verify         Verifier
	writer         io.Writer
}

func NewPool(verify Verifier, writer io.Writer, maxWorkers, maxConnections int) Pool {

	return Pool{
		verify:         verify,
		writer:         writer,
		maxWorkers:     maxWorkers,
		maxConnections: maxConnections,
	}
}

func (p Pool) Start(emails []string) {

	p.jobs = make(chan string, p.maxConnections)
	p.results = make(chan Result, p.maxConnections)

	go p.allocate(emails)

	done := make(chan bool)
	go p.write(done)

	p.pool()
	<-done
}

func (p Pool) allocate(emails []string) {

	for _, email := range emails {
		p.jobs <- email
	}

	close(p.jobs)
}

func (p Pool) write(done chan bool) {
	for result := range p.results {
		_, err := p.writer.Write([]byte(result.Print()))
		if err != nil {
			log.Printf("pool writer error %v", err)
		}
	}
	done <- true
}

func (p Pool) pool() {

	var wg sync.WaitGroup
	for i := 0; i < p.maxWorkers; i++ {
		wg.Add(1)
		go p.worker(&wg)
	}
	wg.Wait()
	close(p.results)
}

func (p Pool) worker(wg *sync.WaitGroup) {
	for email := range p.jobs {
		result := Result{email, p.verify(email)}
		counter++
		p.results <- result
	}
	wg.Done()
}
