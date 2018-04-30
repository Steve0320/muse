package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"log"
	"net/http"
	"muse2/tv"
	"github.com/manyminds/api2go"
	"fmt"
	"time"
	"muse2/helpers"
)

// Define flags
var logPath = flag.String("logFile", "", "location of logfile (leave blank for stdout)")
var dbPath = flag.String("dbFile", "muse.db", "location of database")
var bindAddr = flag.String("bindAddr", "127.0.0.1", "address to serve web interface on")
var port = flag.Int("port", 3000, "port to run web interface on")

// Define loggers
var setupLogger *log.Logger
var fatalLogger *log.Logger
var httpLogger *log.Logger
var dbLogger *log.Logger

func main() {

	flag.Parse()

	// Declare often-used variables
	var logFile *os.File
	var err error

	// Open logging file
	if *logPath == "" {
		logFile = os.Stdout
	} else {
		logFile, err = os.OpenFile(*logPath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Log file %s could not be opened\n", *logPath)
		}
	}

	// Initialize loggers
	setupLogger = log.New(logFile, "SETUP: ", log.Ldate | log.Ltime | log.Lshortfile)
	fatalLogger = log.New(logFile, "FATAL: ", log.Ldate | log.Ltime | log.Lshortfile)
	httpLogger = log.New(logFile, "SERVER: ", log.Ldate | log.Ltime)
	dbLogger = log.New(logFile, "DATABASE: ", log.Ldate | log.Ltime)

	// Perform initializations
	db := initDB()
	initTables(db)
	server := initServer(db)

	httpLogger.Printf("Starting API server on %s\n", server.Addr)
	httpLogger.Fatal(server.ListenAndServe())

}

//*********************** Convenience Initializers ***********************//

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", *dbPath)
	helpers.CheckError(err, fatalLogger)
	db.SetLogger(dbLogger)
	db.LogMode(true)
	setupLogger.Print("Completed database initialization")
	return db
}

func initTables(db *gorm.DB) {
	helpers.CheckError(tv.InitTables(db), fatalLogger)
	setupLogger.Print("Completed table initialization")

}

func initServer(db *gorm.DB) *http.Server {

	// Setup API
	api := api2go.NewAPI("api/v1")

	api.AddResource(tv.Show{}, tv.NewShow(db))
	api.AddResource(tv.Season{}, tv.NewSeason(db))
	api.AddResource(tv.Episode{}, tv.NewEpisode(db))

	// Setup HTTP server
	server := &http.Server {
		Handler: helpers.LoggingHandler(httpLogger, api.Handler()),
		Addr: fmt.Sprintf("%s:%d", *bindAddr, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	setupLogger.Print("Completed server initialization")

	return server

}
