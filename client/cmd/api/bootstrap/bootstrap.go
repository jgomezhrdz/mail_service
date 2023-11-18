package bootstrap

import (
	"context"
	"fmt"
	mailing "mail_service/internal"
	"mail_service/internal/platform/bus/inmemory"
	"mail_service/internal/platform/server"
	"mail_service/internal/platform/storage/mysql"
	cliente_services "mail_service/internal/services/cliente"
	"time"

	gormsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "root"
	dbPass    = "pass"
	dbHost    = "localhost"
	dbPort    = "10101"
	dbName    = "example"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(gormsql.Open(mysqlURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	var (
		eventBus = inmemory.NewEventBus()
	)

	clienteReposiroty := mysql.NewClienteRepository(db)

	clienteServices := cliente_services.NewClienteService(clienteReposiroty, eventBus)
	increasingCourseCounterService := cliente_services.NewCourseCounterService()

	eventBus.Subscribe(
		mailing.ClienteCreatedEventType,
		cliente_services.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, clienteServices)
	return srv.Run(ctx)
}
