package startup

import (
	"ittech24/rest/apidemo/executioncontext"
	"ittech24/rest/apidemo/log"
	"ittech24/rest/apidemo/version"
)

type ServiceProvider struct {
	Context *executioncontext.Context
	Version *version.Version
	Logger  *log.Logger
}

var globalProvider *ServiceProvider

func CreateProvider() *ServiceProvider {
	if globalProvider != nil {
		return globalProvider
	}

	globalProvider = &ServiceProvider{}
	globalProvider.Context = executioncontext.Get()
	globalProvider.Logger = log.Get()
	globalProvider.Version = version.Get()

	return globalProvider
}

func GetServiceProvider() *ServiceProvider {
	return globalProvider
}
