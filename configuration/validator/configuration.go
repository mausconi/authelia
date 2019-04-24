package validator

import (
	"fmt"

	"github.com/clems4ever/authelia/configuration/schema"
)

var defaultPort = 8080
var defaultLogsLevel = "info"

// Validate and adapt the configuration read from file.
func Validate(configuration *schema.Configuration, validator *Validator) {
	if configuration.Port == 0 {
		configuration.Port = defaultPort
	}

	if configuration.LogsLevel == "" {
		configuration.LogsLevel = defaultLogsLevel
	}

	if configuration.Secret == "" {
		validator.Push(fmt.Errorf("Provide a secret using `secret` key"))
	}

	ValidateAuthenticationBackend(&configuration.AuthenticationBackend, validator)
	ValidateSession(&configuration.Session, validator)

	if configuration.TOTP == nil {
		configuration.TOTP = &schema.TOTPConfiguration{}
		ValidateTOTP(configuration.TOTP, validator)
	}
}
