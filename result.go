package main

type Result struct {
	email string
	err   error
}

func (r Result) Print() string {

	result := "OK"

	if r.err != nil {
		result = r.err.Error()
	}

	return r.email + "," + result + "\n"
}
