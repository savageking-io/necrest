package main

import (
	"os"

	"github.com/savageking-io/necconf"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type RESTConfig struct {
	Port uint16 `yaml:"port"`
}

func main() {
	app := cli.NewApp()
	app.Name = "necrest"
	app.Version = AppVersion
	app.Description = "Smart backend service for smart game developers"
	app.Usage = "REST Microservice of NoErrorCode ecosystem"

	app.Authors = []cli.Author{
		{
			Name:  "savageking.io",
			Email: "i@savageking.io",
		},
		{
			Name:  "Mike Savochkin (crioto)",
			Email: "mike@crioto.com",
		},
	}

	app.Copyright = "2025 (c) savageking.io. All Rights Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start REST",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ConfigFilepath,
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       LogLevel,
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
	}

	_ = app.Run(os.Args)
}

func Serve(c *cli.Context) error {
	SetLogLevel(LogLevel)

	dir, file, err := necconf.ExtractDirectoryAndFilenameFromPath(ConfigFilepath)
	if err != nil {
		log.Error("Bad configuration file: %s", err.Error())
		return err
	}

	config := new(necconf.Config)
	if err := config.Init(dir); err != nil {
		log.Errorf("Unrecoverable error: %s", err.Error())
		return err
	}

	conf := new(RESTConfig)
	fs := os.DirFS(dir)
	if err := config.ReadConfig(fs, file, &conf); err != nil {
		log.Errorf("Failed to read configuration: %s", err.Error())
		return err
	}

	return nil
}

func SetLogLevel(level string) {
	switch level {
	case "trace":
		log.SetLevel(log.TraceLevel)
		break
	case "debug":
		log.SetLevel(log.DebugLevel)
		break
	case "info":
		log.SetLevel(log.InfoLevel)
		break
	case "warn":
		log.SetLevel(log.WarnLevel)
		break
	case "error":
		log.SetLevel(log.ErrorLevel)
		break
	default:
		log.SetLevel(log.InfoLevel)
	}
}
