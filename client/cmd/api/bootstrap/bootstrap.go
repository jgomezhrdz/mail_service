package bootstrap

import (
	"database/sql"
	"fmt"
	"mail_service/internal/platform/server"
	"mail_service/internal/platform/storage/mysql"

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
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	clienteReposiroty := mysql.NewClienteRepository(db)

	srv := server.New(host, port, clienteReposiroty)
	return srv.Run()
}
