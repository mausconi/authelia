package handlers

import (
	"fmt"

	"github.com/clems4ever/authelia/middlewares"
	"github.com/tstranex/u2f"
)

// SecondFactorU2FRegister handler validating the client has successfully validated the challenge
// to complete the U2F registration.
func SecondFactorU2FRegister(ctx *middlewares.AutheliaCtx) {
	responseBody := u2f.RegisterResponse{}
	err := ctx.ParseBody(&responseBody)

	userSession := ctx.GetSession()

	if userSession.U2FChallenge == nil {
		ctx.Error(fmt.Errorf("U2F registration has not been initiated yet"), unableToRegisterSecurityKeyMessage)
		return
	}
	// Ensure the challenge is cleared if anything goes wrong.
	defer func() {
		userSession.U2FChallenge = nil
		ctx.SaveSession(userSession)
	}()

	registration, err := u2f.Register(responseBody, *userSession.U2FChallenge, u2fConfig)

	if err != nil {
		ctx.Error(fmt.Errorf("Unable to verify U2F registration: %s", err), unableToRegisterSecurityKeyMessage)
		return
	}

	b, err := registration.MarshalBinary()
	if err != nil {
		ctx.Error(fmt.Errorf("Unable to marshal U2F registration data: %s", err), unableToRegisterSecurityKeyMessage)
		return
	}

	ctx.Providers.StorageProvider.SaveU2FRegistration(userSession.Username, b)

	ctx.ReplyOK()
}
