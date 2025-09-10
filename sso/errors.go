package sso

import "errors"

var (
	noLoginError           = errors.New("no login error")
	sessionIDNotFoundError = errors.New("session ID not found")
)
