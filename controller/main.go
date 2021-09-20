package controller

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"net/http"

	"github.com/cjlapao/common-go/helper"
	commonLogger "github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/security"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/restapi-testapp-go/database"
	"github.com/cjlapao/restapi-testapp-go/entities"
	"github.com/cjlapao/restapi-testapp-go/repositories"
	"github.com/cjlapao/restapi-testapp-go/startup"
	"github.com/gorilla/mux"
)

var logger = commonLogger.Get()
var versionSvc = version.Get()

// var router mux.Router
var databaseContext database.MongoFactory
var repo repositories.Repository
var serviceProvider = startup.CreateProvider()

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

	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/login", controller.Login).Methods("POST")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/validate", controller.Validate).Methods("GET")

	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/article", security.AuthenticateMiddleware(controller.GetAllArticles)).Methods("GET")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/article", security.AuthenticateMiddleware(controller.PostArticle)).Methods("POST")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/article/{id}", security.AuthenticateMiddleware(controller.GetArticle)).Methods("GET")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/article/{id}", security.AuthenticateMiddleware(controller.PutArticle)).Methods("PUT")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/article/{id}", security.AuthenticateMiddleware(controller.DeleteArticle)).Methods("DELETE")

	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator", controller.Generator).Methods("GET")

	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/hello", controller.Hello).Methods("GET")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/hello/{name}", controller.Hello).Methods("GET")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/banana/{name}", controller.Hello).Methods("GET")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/status/{name}", controller.Hello).Methods("GET")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/auth/hello", security.AuthenticateMiddleware(controller.Hello)).Methods("GET")
	controller.Router.Handle(serviceProvider.Context.ApiPrefix+"/auth/hello/{name}", security.AuthenticateMiddleware(controller.Hello)).Methods("GET")

	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/debug", controller.Debug).Methods("GET")

	globalController = &controller
	return globalController
}

func RestApiModuleProcessor() {
	logger.Notice("Starting Go Rest API v%v", versionSvc.String())
	if serviceProvider.Context.BackendEnabled {
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
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc(serviceProvider.Context.ApiPrefix+"/", homePage)
	_ = NewAPIController(router, repo)

	// Creating and starting the http server
	srv := &http.Server{
		Addr:    ":" + serviceProvider.Context.ApiPort,
		Handler: router,
	}

	go func() {
		logger.Info("Api listening on http://localhost:" + serviceProvider.Context.ApiPort + serviceProvider.Context.ApiPrefix)
		logger.Success("Finished Initiating http server")
		if err := srv.ListenAndServe(); err != nil {
			logger.FatalError(err, "There was an error")
		}
	}()

	if serviceProvider.Context.TLS {
		cert, err := tls.X509KeyPair([]byte(serviceProvider.Context.TLSCertificate), []byte(serviceProvider.Context.TLSPrivateKey))
		if err == nil {
			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{cert},
			}

			sslSrv := &http.Server{
				Addr:      ":" + serviceProvider.Context.TLSApiPort,
				TLSConfig: tlsConfig,
				Handler:   router,
			}

			go func() {
				logger.Info("Api listening on https://localhost:" + serviceProvider.Context.TLSApiPort + serviceProvider.Context.ApiPrefix)
				logger.Success("Finished Initiating https server")
				if err := sslSrv.ListenAndServeTLS("", ""); err != nil {
					logger.FatalError(err, "There was an error")
				}
			}()
		} else {
			logger.Error("There was an error reading the certificates to enable HTTPS")
		}
	}

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	logger.Info("Server shutdown")
	// http.ListenAndServeTLS(":10001", "./ssl/local-cluster.internal.crt", "./ssl/local-cluster.internal.key", router)
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
