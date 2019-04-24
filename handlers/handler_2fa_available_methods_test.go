package handlers

import (
	"testing"

	"github.com/clems4ever/authelia/mocks"

	"github.com/clems4ever/authelia/configuration/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SecondFactorAvailableMethodsFixture struct {
	suite.Suite

	mock *mocks.MockAutheliaCtx
}

func (s *SecondFactorAvailableMethodsFixture) SetupTest() {
	s.mock = mocks.NewMockAutheliaCtx(s.T())
}

func (s *SecondFactorAvailableMethodsFixture) TearDownTest() {
	s.mock.Close()
}

func (s *SecondFactorAvailableMethodsFixture) TestShouldServeDefaultMethods() {
	SecondFactorAvailableMethodsGet(s.mock.Ctx)

	assert.Equal(s.T(), "[\"totp\",\"u2f\"]", string(s.mock.Ctx.Response.Body()))
}

func (s *SecondFactorAvailableMethodsFixture) TestShouldServeDefaultMethodsAndDuo() {
	s.mock.Ctx.Configuration = schema.Configuration{
		DuoAPI: &schema.DuoAPIConfiguration{},
	}
	SecondFactorAvailableMethodsGet(s.mock.Ctx)

	assert.Equal(s.T(), "[\"totp\",\"u2f\",\"duo_push\"]", string(s.mock.Ctx.Response.Body()))
}

func TestRunSuite(t *testing.T) {
	s := new(SecondFactorAvailableMethodsFixture)
	suite.Run(t, s)
}
