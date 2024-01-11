package main

import (
	"admin-system/pkg/api"
	"admin-system/pkg/config"
	"admin-system/pkg/services"
	"admin-system/pkg/types"
	"fmt"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version     = "v1.0.0"
	ProgramName = "admin-system"
)

var (
	BuildTime = "2024-01-10"
	GitCommit = ""
)

const (
	CMD_NAME_RUN   = "run"
	CMD_NAME_START = "start"
)

const (
	CMD_FLAG_NAME_DSN    = "dsn"
	CMD_FLAG_NAME_DEBUG  = "debug"
	CMD_FLAG_NAME_STATIC = "static"
	CMD_FLAG_NAME_DOMAIN = "domain"
)

var manager api.ManagerApi

func init() {
	log.SetLevel("info")
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						if manager != nil {
							manager.Close()
						}
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func main() {

	grace()

	local := []*cli.Command{
		runCmd,
	}
	app := &cli.App{
		Name:     ProgramName,
		Version:  fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var runCmd = &cli.Command{
	Name:      CMD_NAME_RUN,
	Usage:     "run as a web service",
	ArgsUsage: "[listen address]",
	Aliases:   []string{CMD_NAME_START},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  CMD_FLAG_NAME_DEBUG,
			Usage: "open debug mode",
		},
		&cli.StringFlag{
			Name:  CMD_FLAG_NAME_DSN,
			Usage: "data source name of database",
			Value: types.DEFAULT_DATA_SOURCE_NAME,
		},
		&cli.StringFlag{
			Name:  CMD_FLAG_NAME_STATIC,
			Usage: "frontend static path",
			Value: types.DefaultStaticHome,
		},
		&cli.StringFlag{
			Name:  CMD_FLAG_NAME_DOMAIN,
			Usage: "http/websocket access domain",
			Value: types.DEFUALT_ACCESS_DOMAIN,
		},
	},
	Action: func(cctx *cli.Context) error {
		cfg := &config.Config{
			Version:  Version,
			HttpAddr: types.DEFAULT_HTTP_LISTEN_ADDR,
			DSN:      cctx.String(CMD_FLAG_NAME_DSN),
			Debug:    cctx.Bool(CMD_FLAG_NAME_DEBUG),
			Static:   cctx.String(CMD_FLAG_NAME_STATIC),
			Domain:   cctx.String(CMD_FLAG_NAME_DOMAIN),
		}

		cfg.Version = Version
		if cfg.Debug {
			log.SetLevel("debug")
		} else {
			log.SetLevel("info")
		}
		log.Json("configuration", cfg)
		if cctx.Args().First() != "" {
			cfg.HttpAddr = cctx.Args().First()
		}
		if err := cfg.Save(); err != nil {
			return err
		}
		//start up as a web server
		manager = services.NewManager(cfg)
		return manager.Run()
	},
}
