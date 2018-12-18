package main

import (
	"fmt"
	"github.com/lasher23/cheetah/pkg/config"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

type T struct {
	Foo string
}

func main() {
	app := cli.NewApp()
	app.Name = "Cheetah"
	app.Usage = "Management Tool for starting programs"
	app.Version = "0.0.1-SNAPSHOT"
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "runs a program of the configuration",
			Action: func(c *cli.Context) error {
				conf, e := config.GetConfig()
				if e == config.NoConfigError {
					e = config.CreateConfigFile()
					if e != nil {
						return e
					}
					fmt.Println("Configuration file created")
					return nil
				} else if e != nil {
					return e
				}
				key := c.Args().Get(0)
				var shortcut config.Shortcut
				var args []string
				for _, v := range conf.Shortcuts {
					if v.Name == key {
						shortcut = v
					}
				}
				if len(c.Args()) > 1 {
					args = c.Args()[1 : len(c.Args())-1]
				}
				command := exec.Command(shortcut.Command, args...)
				command.Dir = shortcut.ExecutionPath
				output, e := command.Output()
				fmt.Print(string(output))
				return e
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}
