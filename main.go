package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"cross-sell/selldb"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

var (
	clientDBServer = os.Getenv("MYSQL_SERVER") //"localhost"
	clientDBUser   = os.Getenv("MYSQL_USER")   //"username"
	clientDBPass   = os.Getenv("MYSQL_PASS")   //"password"
	clientDBPort   = os.Getenv("MYSQL_PORT")   //3306
)

func connectDB()  {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		clientDBUser,
		clientDBPass,
		clientDBServer,
		clientDBPort,
	)

	// open mysql service with mysql driver
	dbx, err := sqlx.Open("mysql", dsn)
	if err != nil {
		logrus.Errorf("Error connecting to mysql through go-sql-driver : %s", err.Error())
	}

	// check if successfully connect
	if err := dbx.Ping(); err != nil {
		logrus.Warnf("Error pinging mysql through go-sql-driver %s", err)
	} else {
		fmt.Printf("Successfully connected to mysql through go-sql-driver: %s:%s\n", clientDBServer, clientDBPort)
	}

	// pass db variable to selldb
	if err = selldb.SetDB(dbx); err != nil {
		logrus.Warnln("Error setting mysql DB for databaseKB", err)
	}
}

func init() {

	// connect to database
	connectDB()

	// perform other stuff here
}

func main() {

	// handlers
	router := NewRouter()

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	// Launch server with CORS validations
	fmt.Printf("Starting server listning at 8080 port...\n")
	if err := http.ListenAndServe(":8080" , handlers.CORS(allowedOrigins, allowedMethods)(router)); err != nil {
		log.Fatal(err)
	}
}

