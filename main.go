package main

import (
	"github.com/honlyc/struct2all/web"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "struct2all"
	app.Usage = "struct2all工具集"

	app.Commands = []cli.Command{
		{
			Name:  "all",
			Usage: "generate sql from golang model struct",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "source file or dir, default: current dir",
				},
				cli.StringFlag{
					Name:  "struct, s",
					Usage: "struct name or pattern: https://golang.org/pkg/path/filepath/#Match",
				},
				cli.StringFlag{
					Name:  "out, o",
					Usage: "output file",
				},
			},
			Action: web.CreateHttp,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

/*func main_test() {
	app := cli.NewApp()
	app.Name = "hfast"
	app.Usage = "hfast工具集"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "o",
					Value:       "",
					Usage:       "project owner for create project",
					Destination: &p.Owner,
				},
				cli.StringFlag{
					Name:        "d",
					Value:       "",
					Usage:       "project directory for create project",
					Destination: &p.Path,
				},
				cli.BoolFlag{
					Name:        "proto",
					Usage:       "whether to use protobuf for create project",
					Destination: &p.WithGRPC,
				},
			},
			Action: runNew,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "hfast build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "hfast run",
			Action:  runAction,
		},
		{
			Name:    "new-page",
			Aliases: []string{"np"},
			Usage:   "hfast new page",
			Action:  runNewPage,
		},
		{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "hfast tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "hfast version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "hfast self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}*/
