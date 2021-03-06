package main

import (
	"fast-wire/app/config"
	"fast-wire/app/server"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type application struct {
	Server *server.Server
}

func newApplication(Server *server.Server) application {
	return application{Server: Server}
}

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", "../.env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}
	initLogging(config)

	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		fmt.Println(config.String())
	}

	app, err := InitializeApplication(config)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	g := errgroup.Group{}
	g.Go(func() error {
		logrus.WithFields(
			logrus.Fields{
				"Host": config.Server.Host,
			},
		).Info("starting the http server")
		return app.Server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logrus.WithError(err).Fatalln("program terminated")
	}
}

func initLogging(c config.Config) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if c.Logging.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}

	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: c.Logging.Pretty,
		})
	}
}
