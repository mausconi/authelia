package main

import (
	"errors"
	"flag"
	"os"

	"github.com/clems4ever/authelia/regulation"

	"github.com/clems4ever/authelia/session"

	"github.com/clems4ever/authelia/authentication"
	"github.com/clems4ever/authelia/authorization"
	"github.com/clems4ever/authelia/configuration"
	"github.com/clems4ever/authelia/logging"
	"github.com/clems4ever/authelia/middlewares"
	"github.com/clems4ever/authelia/notification"
	"github.com/clems4ever/authelia/server"
	"github.com/clems4ever/authelia/storage"
	"github.com/sirupsen/logrus"
)

func tryExtractConfigPath() (string, error) {
	configPtr := flag.String("config", "", "The path to a configuration file.")
	flag.Parse()

	if *configPtr == "" {
		return "", errors.New("No config file path provided")
	}

	return *configPtr, nil
}

func main() {
	if os.Getenv("ENVIRONMENT") == "dev" {
		logging.Logger().Info("===> Authelia is running in development mode. <===")
	}

	configPath, err := tryExtractConfigPath()
	if err != nil {
		logging.Logger().Error(err)
	}

	config, errs := configuration.Read(configPath)

	if len(errs) > 0 {
		for _, err = range errs {
			logging.Logger().Error(err)
		}
		panic(errors.New("Some errors have been reported"))
	}

	switch config.LogsLevel {
	case "info":
		logging.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		logging.SetLevel(logrus.TraceLevel)
	}

	fileUserProvider := authentication.NewFileUserProvider(
		"./test/suites/basic/users_database.test.yml")
	storageProvider := storage.NewSQLiteProvider(config.Storage.Local.Path)
	notifier := notification.NewSMTPNotifier(config.Notifier.SMTP)
	authorizer := authorization.NewAuthorizer(*config.AccessControl)
	sessionProvider := session.NewProvider(config.Session)
	regulator := regulation.NewRegulator(config.Regulation, storageProvider)

	providers := middlewares.Providers{
		Authorizer:      authorizer,
		UserProvider:    &fileUserProvider,
		Regulator:       regulator,
		StorageProvider: storageProvider,
		Notifier:        notifier,
		SessionProvider: sessionProvider,
	}
	server.StartServer(*config, providers)
}
