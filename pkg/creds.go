package stargoat

import (
	"fmt"
	"os"
)

type CredsProvider interface {
	Get() (string, string, error)
}

type EnvVarCredsProvider struct {
	UsernameEnv string
	PasswordEnv string
}

func (e *EnvVarCredsProvider) Get() (string, string, error) {
	user, userOk := os.LookupEnv(e.UsernameEnv)
	pass, passOk := os.LookupEnv(e.PasswordEnv)
	if userOk && passOk {
		return user, pass, nil
	}
	if !userOk && !passOk {
		return "", "", fmt.Errorf("error getting credentials, %s and %s env vars missing", e.UsernameEnv, e.PasswordEnv)
	}
	if !userOk {
		return "", "", fmt.Errorf("error getting credentials, %s env var missing", e.UsernameEnv)
	}
	// remaining case is that passwordEnv is missing
	return "", "", fmt.Errorf("error getting credentials, %s env var missing", e.PasswordEnv)
}
