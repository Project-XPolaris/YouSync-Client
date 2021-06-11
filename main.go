package main

import (
	srv "github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"yousyncclient/api"
	"yousyncclient/client"
	"yousyncclient/config"
)

var svcConfig *srv.Config
var Logger = logrus.WithField("scope", "main")

func initService() error {
	workPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	svcConfig = &srv.Config{
		Name:             "YouSyncClientService",
		DisplayName:      "YouSyncClient Service",
		WorkingDirectory: workPath,
		Arguments:        []string{"run"},
	}
	return nil
}
func Program() {
	err := config.ReadConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		client.DefaultSyncClient.Run()
	}()
	api.RunAPIService()
}

type program struct{}

func (p *program) Start(s srv.Service) error {
	go Program()
	return nil
}

func (p *program) Stop(s srv.Service) error {
	return nil
}

func InstallAsService() {
	prg := &program{}
	s, err := srv.New(prg, svcConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	s.Uninstall()

	err = s.Install()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("successful install service")
}

func UnInstall() {

	prg := &program{}
	s, err := srv.New(prg, svcConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	s.Uninstall()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("successful uninstall service")
}

func StartService() {
	prg := &program{}
	s, err := srv.New(prg, svcConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
func StopService() {
	prg := &program{}
	s, err := srv.New(prg, svcConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	err = s.Stop()
	if err != nil {
		logrus.Fatal(err)
	}
}
func RestartService() {
	prg := &program{}
	s, err := srv.New(prg, svcConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	err = s.Restart()
	if err != nil {
		logrus.Fatal(err)
	}
}
func RunApp() {
	app := &cli.App{
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "service",
				Usage: "service manager",
				Subcommands: []*cli.Command{
					{
						Name:  "install",
						Usage: "install service",
						Action: func(context *cli.Context) error {
							InstallAsService()
							return nil
						},
					},
					{
						Name:  "uninstall",
						Usage: "uninstall service",
						Action: func(context *cli.Context) error {
							UnInstall()
							return nil
						},
					},
					{
						Name:  "start",
						Usage: "start service",
						Action: func(context *cli.Context) error {
							StartService()
							return nil
						},
					},
					{
						Name:  "stop",
						Usage: "stop service",
						Action: func(context *cli.Context) error {
							StopService()
							return nil
						},
					},
					{
						Name:  "restart",
						Usage: "restart service",
						Action: func(context *cli.Context) error {
							RestartService()
							return nil
						},
					},
				},
				Description: "YouSyncClient service controller",
			},
			{
				Name:  "run",
				Usage: "run app",
				Action: func(context *cli.Context) error {
					Program()
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := initService()
	if err != nil {
		logrus.Fatal(err)
	}
	RunApp()
}
