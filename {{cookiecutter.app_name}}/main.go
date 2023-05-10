package main

import (
	"context"
	"github.com/spf13/viper"
	"midigator-portfolios/cookiecutter-golang/runner"

	//"flag"
	//"fmt"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/instance"
	"sync"

	//"midigator-portfolios/cookiecutter-golang/version"
	"os"
	"os/signal"
	"syscall"

	//{% if cookiecutter.use_cobra_cmd == "n" %}
	//	"flag"
	//	"fmt"
	//	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
	//{% endif %}
	//{% if cookiecutter.use_cobra_cmd == "y" %}
	//	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/cmd"
	//{% endif %}
	//"flag"

	"midigator-portfolios/cookiecutter-golang/cmd"
	//"midigator-portfolios/cookiecutter-golang/version"
	"github.com/urfave/cli"
)

func main() {
	v := viper.New()
	//c := config.Init(v)
	//i := instance.Init(c)
	config := config.Init(v)
	instance := instance.Init(config)
	//cmd.Config = c
	//cmd.Instance = i
	//defer i.Destroy()
	defer instance.Destroy()

	var shutDownChannel chan *bool
	//{% if cookiecutter.use_cobra_cmd == "y" %}
	cmd.Execute()
	//{% else %}
	//	versionFlag := flag.Bool("version", false, "Version")
	//	flag.Parse()
	//
	//	if *versionFlag {
	//		fmt.Println("Build Date:", version.BuildDate)
	//	   fmt.Println("Git Commit:", version.GitCommit)
	//	   fmt.Println("Version:", version.Version)
	//	   fmt.Println("Go Version:", version.GoVersion)
	//	   fmt.Println("OS / Arch:", version.OsArch)
	//		return
	//	}
	//{% endif %}

	clientApp := cli.NewApp()
	clientApp.Name = "go-cookiecutter"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup

				wg.Add(1)
				go runner.NewAPI(config, instance).Go(ctx, &wg)

				wg.Add(1)
				go runner.NewGRPC(config, instance).Go(ctx, &wg)

				wg.Wait()
				return nil

			},
		},
		{
			Name:  "start_workers",
			Usage: "Start the workers",
			Action: func(c *cli.Context) error {
				ctx := context.Background()

				var wg sync.WaitGroup

				wg.Add(1)
				go runner.NewWorker(config, instance).Go(ctx, shutDownChannel, &wg)
				wg.Wait()

				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		<-signalChannel
		shutDown(shutDownChannel)
	}()
}

func shutDown(shutDownChannel chan *bool) {
	shutDown := true
	shutDownChannel <- &shutDown
	close(shutDownChannel)
}
