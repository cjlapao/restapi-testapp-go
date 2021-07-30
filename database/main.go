package database

import (
	"ittech24/rest/apidemo/executioncontext"
	"ittech24/rest/apidemo/log"
)

var ctx = executioncontext.Get()
var logger = log.Get()
