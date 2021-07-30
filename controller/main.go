package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ittech24/rest/apidemo/database"
	"ittech24/rest/apidemo/entities"
	"ittech24/rest/apidemo/helper"
	ittechlog "ittech24/rest/apidemo/log"
	"ittech24/rest/apidemo/security"
	"ittech24/rest/apidemo/version"
	"log"
	"os"

	"ittech24/rest/apidemo/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

var logger = ittechlog.Get()
var versionSvc = version.Get()
var router mux.Router
var databaseContext database.MongoFactory
var repo repositories.Repository

type article = entities.Article

// Controller Controller structure
type Controller struct {
	Router     *mux.Router
	Repository *repositories.Repository
}

var globalController *Controller

// NewAPIController  Creates a new controller
func NewAPIController(router *mux.Router, repo repositories.Repository) *Controller {
	if globalController != nil {
		return globalController
	}

	controller := Controller{
		Router:     router,
		Repository: &repo,
	}

	controller.Router.Handle("/article", security.AuthenticateMiddleware(controller.GetAllArticles)).Methods("GET")
	controller.Router.Handle("/article", security.AuthenticateMiddleware(controller.PostArticle)).Methods("POST")
	controller.Router.Handle("/article/{id}", security.AuthenticateMiddleware(controller.GetArticle)).Methods("GET")
	controller.Router.Handle("/article/{id}", security.AuthenticateMiddleware(controller.PutArticle)).Methods("PUT")
	controller.Router.Handle("/article/{id}", security.AuthenticateMiddleware(controller.DeleteArticle)).Methods("DELETE")

	controller.Router.HandleFunc("/login", controller.Login).Methods("POST")
	controller.Router.HandleFunc("/validate", controller.Validate).Methods("GET")

	controller.Router.HandleFunc("/hello", controller.Hello).Methods("GET")
	controller.Router.HandleFunc("/hello/{name}", controller.Hello).Methods("GET")
	controller.Router.Handle("/auth/hello", security.AuthenticateMiddleware(controller.Hello)).Methods("GET")
	controller.Router.Handle("/auth/hello/{name}", security.AuthenticateMiddleware(controller.Hello)).Methods("GET")

	globalController = &controller
	return globalController
}

func RestApiModuleProcessor() {
	logger.Notice("Starting Go Rest API v%v", versionSvc.String())
	cs := os.Getenv("mongoConnectionString")
	if cs != "" {
		logger.Info("Found MongoDB connection, enabling MongoDb backend...")
		databaseContext = database.NewFactory()
		repo = repositories.NewRepo(&databaseContext)
		pushTestData()
	}

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("endpoint Hit: homepage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", homePage)
	_ = NewAPIController(router, repo)
	logger.Success("Finished Init")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func pushTestData() {
	importData := entities.ImportData{}
	if !helper.FileExists("demo.json") {
		logger.Error("File demo.json not found, not importing data")
	}

	data, err := ioutil.ReadFile("demo.json")

	if err != nil {
		logger.LogError(err)
		return
	}

	json.Unmarshal(data, &importData)
	if len(importData.Articles) > 0 {
		logger.Info("Importing Demo articles")
		repo.UpsertManyArticles(importData.Articles)
	}
	if len(importData.People) > 0 {
		logger.Info("Importing Demo people")
		repo.UpsertManyPersons(importData.People)
	}
}
