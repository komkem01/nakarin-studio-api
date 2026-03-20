package config

import (
	"nakarin-studio/app/modules/example"
	exampletwo "nakarin-studio/app/modules/example-two"
	"nakarin-studio/app/modules/sentry"
	"nakarin-studio/app/modules/specs"
	"nakarin-studio/internal/kafka"
	"nakarin-studio/internal/log"
	"nakarin-studio/internal/otel/collector"
)

// Config is a struct that contains all the configuration of the application.
type Config struct {
	Database Database

	AppName     string
	AppKey      string
	Environment string
	Specs       specs.Config
	Debug       bool

	Port           int
	HttpJsonNaming string

	SslCaPath      string
	SslPrivatePath string
	SslCertPath    string

	Otel   collector.Config
	Sentry sentry.Config

	Kafka kafka.Config
	Log   log.Option

	Example example.Config

	ExampleTwo exampletwo.Config
}

var App = Config{
	Specs: specs.Config{
		Version: "v1",
	},
	Database: database,
	Kafka:    kafkaConf,

	AppName: "go_app",
	Port:    8080,
	AppKey:  "secret",
	Debug:   false,

	HttpJsonNaming: "snake_case",

	SslCaPath:      "nakarin-studio/cert/ca.pem",
	SslPrivatePath: "nakarin-studio/cert/server.pem",
	SslCertPath:    "nakarin-studio/cert/server-key.pem",

	Otel: collector.Config{
		CollectorEndpoint: "",
		LogMode:           "noop",
		TraceMode:         "noop",
		MetricMode:        "noop",
		TraceRatio:        0.01,
	},
}
