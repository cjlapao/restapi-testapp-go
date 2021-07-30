package executioncontext

import (
	"ittech24/rest/apidemo/helper"
	"ittech24/rest/apidemo/log"
	"os"
)

var globalContext *Context
var logger = log.Get()

// Context entity
type Context struct {
	Module                string `json:"module"`
	Operation             string `json:"operation"`
	TenantID              string `json:"tenantId"`
	MongoConnectionString string `json:"mongoConnectionString"`
	Database              string `json:"database"`
	ShowHelp              bool   `json:"help"`
}

func Get() *Context {
	if globalContext != nil {
		return globalContext
	}

	logger.Debug("Creating Execution Context")
	globalContext = &Context{
		Operation: helper.GetFlagValue("operation", "api"),
		ShowHelp:  helper.GetFlagSwitch("help", false),
	}

	globalContext.Getenv()

	return globalContext
}

// Getenv gets the environment variables for the entities
func (e *Context) Getenv() {

	op := os.Getenv("DO_OPERATION")
	module := os.Getenv("DT_MODULE")

	if len(op) > 0 {
		e.Operation = op
	}

	if len(module) > 0 {
		e.Module = module
	}
}
