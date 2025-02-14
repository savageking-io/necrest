package main

const (
	ConfigurationDirectory = "/etc/noerrorcode"
)

var (
	// AppVersion is a version of the application. Must be set during build
	AppVersion     = "0.0.0"
	ConfigFilepath = "/etc/noerrorcode/rest.yaml"
	LogLevel       = "info"
)
