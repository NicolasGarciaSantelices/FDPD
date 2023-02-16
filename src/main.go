package main

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	services "FDPD-BACKEND/src/services/api"
	"FDPD-BACKEND/src/utils"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&utils.LogFormat{})

	logger := logrus.WithFields(nil)
	logger.Info("Initializing api...")

	allServices := []services.Service{
		services.NewPostgreSQLService(
			os.Getenv("SQL_CONNECTION_STRING"),
		),
		services.NewAPIService("8080"),
	}

	for _, service := range allServices {
		service.InjectServices(logger, allServices)
	}

	for _, service := range allServices {
		service.Init()
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(allServices))

	for _, service := range allServices {
		go service.Execute(&waitGroup)
	}
	waitGroup.Wait()
}
