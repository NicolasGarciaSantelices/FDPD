package services

import (
	"database/sql"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// postgreSQLService postgreSQLService
type PostgreSQLService struct {
	connectionString string
	logger           *logrus.Entry
	db               *sql.DB
}

// NewpostgreSQLService returns a service instance.
func NewPostgreSQLService(cop string) *PostgreSQLService {
	return &PostgreSQLService{
		connectionString: cop,
	}
}

// Health Health
func (service *PostgreSQLService) Health() bool {
	return true
}

// InjectServices InjectServices
func (service *PostgreSQLService) InjectServices(logger *logrus.Entry, otherServices []Service) {
	service.logger = logger
}

// Init Init this service
func (service *PostgreSQLService) Init() error {
	service.logger.Info("[postgreSQLService] Initializing...")

	service.logger.Info("[postgreSQLService] Using connection string for COPILOTO: " + service.connectionString)

	var err error

	service.db, err = sql.Open("postgres", service.connectionString)
	if err != nil {
		service.logger.Fatal("[postgreSQLService] Failed connecting to database: " + err.Error())
		return err
	} else {
		service.logger.Info("[postgreSQLService] Connected to database COPILOTO successfully")
	}

	return nil
}

// Execute Execute this service
func (service *PostgreSQLService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[postgreSQLService] Executing...")

	defer service.db.Close()

	waitGroup.Wait()

	return nil
}
