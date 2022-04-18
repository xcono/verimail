package main

import (
	"errors"
	emailverifier "github.com/AfterShip/email-verifier"
	"log"
)

func Verify(email string) (err error) {

	log.Printf("verifying: %s", email)

	ret, err := VerifyResult(email)

	if err != nil {
		return
	}

	if ret.SMTP.FullInbox {
		return errors.New("full inbox")
	}

	if !(ret.SMTP.CatchAll || ret.SMTP.Deliverable) {
		return errors.New("not deliverable")
	}

	if !ret.SMTP.HostExists {
		return errors.New("host does not exists")
	}

	return
}

func VerifyResult(email string) (res *emailverifier.Result, err error) {

	res, err = verifier.Verify(email)

	if err != nil {
		return
	}

	if !res.Syntax.Valid {
		return res, errors.New("syntax is invalid")
	}

	resp, err := verifier.CheckSMTP(res.Syntax.Domain, res.Syntax.Username)
	res.SMTP = resp

	return
}
