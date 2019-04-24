package handlers

import (
	"fmt"

	"github.com/clems4ever/authelia/middlewares"
	"github.com/tstranex/u2f"
)

// SecondFactorU2FSignGet handler for initiating a signing request.
func SecondFactorU2FSignGet(ctx *middlewares.AutheliaCtx) {
	userSession := ctx.GetSession()

	registrationBin, err := ctx.Providers.StorageProvider.LoadU2FRegistration(userSession.Username)

	if err != nil {
		ctx.Error(fmt.Errorf("Unable to retrieve U2F device handle: %s", err), mfaValidationFailedMessage)
		return
	}

	appID := fmt.Sprintf("%s://%s", ctx.XForwardedProto(), ctx.XForwardedHost())
	var trustedFacets = []string{appID}
	challenge, err := u2f.NewChallenge(appID, trustedFacets)

	if err != nil {
		ctx.Error(fmt.Errorf("Unable to create U2F challenge: %s", err), mfaValidationFailedMessage)
		return
	}

	var registration u2f.Registration
	err = registration.UnmarshalBinary(registrationBin)
	if err != nil {
		ctx.Error(fmt.Errorf("Unable to unmarshal U2F device handle: %s", err), mfaValidationFailedMessage)
		return
	}

	// Save the challenge and registration for use in next request
	userSession.U2FRegistration = &registration
	userSession.U2FChallenge = challenge
	err = ctx.SaveSession(userSession)

	if err != nil {
		ctx.Error(fmt.Errorf("Unable to save U2F challenge and registration in session: %s", err), mfaValidationFailedMessage)
		return
	}

	signRequest := challenge.SignRequest([]u2f.Registration{registration})
	err = ctx.SetJSONBody(signRequest)

	if err != nil {
		ctx.Error(fmt.Errorf("Unable to set sign request in body: %s", err), mfaValidationFailedMessage)
		return
	}
}
