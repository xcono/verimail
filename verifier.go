package main

import (
	"errors"
	emailverifier "github.com/AfterShip/email-verifier"
	"log"
)

func Verify(email string, smtpCheck bool) (err error) {

	log.Printf("verifying: %s", email)

	result, err := VerifyResult(email, smtpCheck)

	if err != nil {
		return
	}

	if result.SMTP == nil {
		return
	}

	if result.SMTP.FullInbox {
		return errors.New("full inbox")
	}

	if !(result.SMTP.CatchAll || result.SMTP.Deliverable) {
		return errors.New("not deliverable")
	}

	if !result.SMTP.HostExists {
		return errors.New("host does not exists")
	}

	return
}

func VerifyResult(email string, smtpCheck bool) (res *emailverifier.Result, err error) {

	res, err = verifier.Verify(email)

	if err != nil {
		return
	}

	if !res.Syntax.Valid {
		return res, errors.New("syntax is invalid")
	}

	if smtpCheck {
		var resp *emailverifier.SMTP
		resp, err = verifier.CheckSMTP(res.Syntax.Domain, res.Syntax.Username)
		res.SMTP = resp
	}

	return
}
