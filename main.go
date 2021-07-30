package main

import (
	"fmt"
	"ittech24/rest/apidemo/bootstrap"
	"ittech24/rest/apidemo/controller"
	"ittech24/rest/apidemo/helper"
	"ittech24/rest/apidemo/log"
	"ittech24/rest/apidemo/startup"
	"ittech24/rest/apidemo/version"
	"os"
	"strings"
)

var startupSvc = startup.CreateProvider()
var logger = log.Get()
var versionSvc = version.Get()

func main() {
	versionSvc.Minor = 1
	getVersion := helper.GetFlagSwitch("version", false)
	if getVersion {
		format := helper.GetFlagValue("o", "json")
		switch strings.ToLower(format) {
		case "json":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		case "yaml":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		default:
			fmt.Println("Please choose a valid format, this can be either json or yaml")
		}
		os.Exit(0)
	}

	versionSvc.PrintAnsiHeader()

	bootstrap.Start()

	defer func() {
		bootstrap.Exit(0)
	}()

	controller.RestApiModuleProcessor()

	bootstrap.Exit(0)
}
