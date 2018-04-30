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
	//initTables()
	//router := initRouter()

	// Define and start server
	//server := &http.Server {
	//	Handler: loggingHandler(httpLogger, router),
	//	Addr: fmt.Sprintf("%s:%d", *bindAddr, *port),
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout: 15 * time.Second,
	//}

	//httpLogger.Printf("Starting API server on %s\n", server.Addr)
	//httpLogger.Fatal(server.ListenAndServe())

	show := tv.NewShowSource(db)
	api := api2go.NewAPI("v1")
	api.AddResource(tv.Show{}, &show)
	http.ListenAndServe(":8080", api.Handler())

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
/*
// Convenience method for initializing API routes
func initRouter() *mux.Router {

	//show = tv.NewShowSource()
	//api := api2go.NewAPI("v1")
	//api.AddResource(tv.Show{}, &tv.ShowSource{})
	//http.ListenAndServe(":8080", api.Handler())

	//router := mux.NewRouter().StrictSlash(true)
	//api := api2go.NewAPI("v1")

	// Index routes for TV
	//router.Handle("/api/v1/tv", tv.HandleFullIndex(db, httpLogger))
	//router.Handle("/api/v1/tv/shows", tv.HandleShowsIndex(db, httpLogger))
	//router.Handle("/api/v1/tv/seasons", tv.HandleSeasonsIndex(db, httpLogger))
	//router.Handle("/api/v1/tv/episodes", tv.HandleEpisodesIndex(db, httpLogger))

	// Item routes for TV
	//router.Handle("/api/v1/tv/shows/{id}", tv.HandleShow(db, httpLogger))
	//router.Handle("/api/v1/tv/seasons/{id}", tv.HandleSeason(db, httpLogger))
	//router.Handle("/api/v1/tv/episodes/{id}", tv.HandleEpisode(db, httpLogger))

	// TODO: These routes have a lot of repetition in the handlers, try to shorten up
	// TODO: Implement CRUD operations (only READ so far)
	// Routes for nested resources
	//router.Handle("/api/v1/tv/shows/{id}/seasons", tv.HandleShowSeasonsIndex(db, httpLogger))
	//router.Handle("/api/v1/tv/shows/{id}/seasons/{sid}", tv.HandleShowSeasons(db, httpLogger))
	//router.Handle("/api/v1/tv/shows/{id}/seasons/{sid}/episodes", tv.HandleShowSeasonsEpisodesIndex(db, httpLogger))
	//router.Handle("/api/v1/tv/shows/{id}/seasons/{sid}/episodes/{eid}", tv.HandleShowSeasonsEpisodes(db, httpLogger))

	// Search route for TV
	//router.Handle("/api/v1/tv/search", tv.HandleSearch(db, httpLogger))

	// TODO: /api/v1/tv/shows/search?q=...

	//return router

}*/