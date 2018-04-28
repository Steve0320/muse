package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"time"
	"muse2/tv"
)

// Define flags
var logPath = flag.String("logFile", "stdout", "location of logfile")
var bindAddr = flag.String("bindAddr", "127.0.0.1", "address to serve web interface on")
var port = flag.Int("port", 3000, "port to run web interface on")

// Define loggers
var setupLogger *log.Logger
var fatalLogger *log.Logger
var httpLogger *log.Logger

// Database
var db *gorm.DB

func main() {

	flag.Parse()

	// Declare often-used variables
	var logFile *os.File
	var err error

	// Open logging file
	if *logPath == "stdout"{
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

	// Open database
	db, err = gorm.Open("sqlite3", "muse.db")
	test(err, setupLogger)
	db.LogMode(true)

	// Perform initializations
	initTables()
	router := initRouter()

	// Define and start server
	server := &http.Server {
		Handler: loggingHandler(httpLogger, router),
		Addr: fmt.Sprintf("%s:%d", *bindAddr, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	httpLogger.Printf("Starting API server on %s\n", server.Addr)
	httpLogger.Fatal(server.ListenAndServe())

}

func test(e error, logger *log.Logger) {
	if e != nil {
		logger.Fatal(e)
	}
}

//*********************** API Request Handlers ***********************//

// Handle basic logging using native logger
func loggingHandler(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s request to %s\n", r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Convenience method to initialize all tables
func initTables() {
	if err := tv.InitTables(db); err != nil {
		fatalLogger.Fatal(err)
	}
}

// Convenience method for initializing API routes
func initRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// Index routes for TV shows
	router.Handle("/api/v1/tv", tv.HandleFullIndex(db, httpLogger))
	router.Handle("/api/v1/tv/shows", tv.HandleShowsIndex(db, httpLogger))
	router.Handle("/api/v1/tv/seasons", tv.HandleSeasonsIndex(db, httpLogger))
	router.Handle("/api/v1/tv/episodes", tv.HandleEpisodesIndex(db, httpLogger))

	// TODO
	// /api/v1/tv/shows/{id}/seasons/{id}/episodes/{id}
	// /api/v1/tv/seasons/{id}
	// /api/v1/tv/episodes/{id}

	return router

}