package services

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"FDPD-BACKEND/src/middlewares"
	routers "FDPD-BACKEND/src/routers"
)

// APIService APIService
type APIService struct {
	port              string
	logger            *logrus.Entry
	engine            *gin.Engine
	PostgreSQLService *PostgreSQLService
}

// NewAPIService returns a service instance.
func NewAPIService(port string) *APIService {
	return &APIService{
		port: port,
	}
}

// Health Health
func (service *APIService) Health() bool {
	return true
}

// InjectServices InjectServices
func (service *APIService) InjectServices(logger *logrus.Entry, services []Service) {
	service.logger = logger
	for _, otherService := range services {
		if PostgreSQLService, ok := otherService.(*PostgreSQLService); ok {
			service.PostgreSQLService = PostgreSQLService
		}

	}
}

// Init Init this service
func (service *APIService) Init() error {
	service.logger.Info("[APIService] Initializing...")
	service.logger.Info("[APIService] Using port: " + service.port)

	service.engine = gin.Default()
	service.engine.Use(middlewares.CORSMiddleware())
	service.engine.Use(gin.Recovery())
	return nil
}

// Execute Execute this service
func (service *APIService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[APIService] Executing...")

	routers.EndpointGroup(service.engine, service.PostgreSQLService.db)

	err := service.engine.Run(":" + service.port)
	if err != nil {
		service.logger.Fatal("[APIService] Failed running api server: " + err.Error())
		return err
	} else {
		service.logger.Info("[APIService] Running api server: " + service.port)
	}

	return nil
}
