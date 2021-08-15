package executioncontext

import (
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
)

var globalContext *Context
var logger = log.Get()

// Context entity
type Context struct {
	MongoConnectionString string `json:"mongoConnectionString"`
	Database              string `json:"database"`
	ShowHelp              bool   `json:"help"`
	BackendEnabled        bool   `json:"backendEnabled"`
	ApiPrefix             string `json:"apiPrefix"`
	ApiPort               string `json:"apiPort"`
}

func Get() *Context {
	if globalContext != nil {
		return globalContext
	}

	logger.Debug("Creating Execution Context")
	globalContext = &Context{
		ShowHelp: helper.GetFlagSwitch("help", false),
	}

	globalContext.Getenv()

	return globalContext
}

// Getenv gets the environment variables for the entities
func (e *Context) Getenv() {

	e.Database = os.Getenv("RESTAPI_DATABASENAME")
	e.MongoConnectionString = os.Getenv("RESTAPI_MONGO_CONNECTION_STRING")
	e.ApiPrefix = os.Getenv("RESTAPI_API_PREFIX")
	e.ApiPort = os.Getenv("RESTAPI_API_PORT")

	if e.MongoConnectionString == "" {
		e.BackendEnabled = false
	} else {
		e.BackendEnabled = true
	}

	if e.Database == "" {
		e.Database = "restapi_demo_db"
	}

	if e.ApiPort == "" {
		e.ApiPort = "80"
	}

	if strings.HasSuffix(e.ApiPrefix, "/") {
		e.ApiPrefix = strings.TrimSuffix(e.ApiPrefix, "/")
	}
}
