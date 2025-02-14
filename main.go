package main

import (
	"os"

	"github.com/savageking-io/necconf"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type RESTConfig struct {
	Port uint16
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
					Value:       "rest.yaml",
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       "info",
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
	}

	_ = app.Run(os.Args)
}

func Serve(c *cli.Context) error {
	config := new(necconf.Config)
	err := config.Init(ConfigurationDirectory)
	if err != nil {
		log.Errorf("Unrecoverable error: %s", err.Error())
		return err
	}

	conf := new(RESTConfig)
	dir := os.DirFS(ConfigurationDirectory)
	err = config.ReadConfig(dir, ConfigFilepath, &conf)
	if err != nil {
		log.Errorf("Failed to read configuration: %s", err.Error())
		return err
	}

	return nil
}
