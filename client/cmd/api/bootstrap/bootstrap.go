package bootstrap

import (
	"fmt"
	"mail_service/internal/platform/server"
	"mail_service/internal/platform/storage/mysql"
	cliente_services "mail_service/internal/services/cliente"

	gormsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "pass"
	dbHost = "localhost"
	dbPort = "10101"
	dbName = "example"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(gormsql.Open(mysqlURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	clienteReposiroty := mysql.NewClienteRepository(db)
	clienteServices := cliente_services.NewClienteService(clienteReposiroty)

	srv := server.New(host, port, clienteServices)
	return srv.Run()
}
