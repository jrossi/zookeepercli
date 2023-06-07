/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-zkcli/zkcli/output"
	"github.com/go-zkcli/zkcli/zk"
	"github.com/outbrain/golib/log"
	"github.com/urfave/cli"
)

var (
	version = "1.1.0"
)

func main() {

	var format string
	app := cli.NewApp()
	app.Version = version
	app.Name = "zkcli"
	app.Usage = "zkcli is a non-interactive command line client for ZooKeeper"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "servers",
			Value:  "localhost:2181",
			Usage:  "ZK Server list in format: srv1[:port1][,srv2[:port2]...]",
			EnvVar: "ZKC_SERVERS",
		},
		cli.StringFlag{
			Name:   "format",
			Value:  "txt",
			Usage:  "Output format",
			EnvVar: "ZKC_FORMAT",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalString("servers") == "" {
			log.Fatal("Expected comma delimited list of servers via --servers|$ZKC_SERVERS")
		}
		serversArray := strings.Split(c.GlobalString("servers"), ",")
		if len(serversArray) == 0 {
			log.Fatal("Expected comma delimited list of servers via --servers")
		}
		zk.SetServers(serversArray)

		switch c.GlobalString("format") {
		case "json", "txt":
			format = c.GlobalString("format")
		default:
			log.Fatalf("Format %s is not in list: json,text", c.GlobalString("format"))

		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:        "exists",
			Usage:       "zkcli excists <path>",
			Description: "Does znode at <path> exist.  Exit with error code 0 if present code 1 if not.",
			Action: func(c *cli.Context) {
				if exists, err := zk.Exists(c.Args().First()); err == nil && exists {
					output.PrintString([]byte("true"), format)
				} else {
					log.Fatale(err)
				}
			},
		},
		{
			Name:        "get",
			Usage:       "zkcli get <path>",
			Description: "Get value on znode at <path>.",
			Action: func(c *cli.Context) {
				if result, err := zk.Get(c.Args().First()); err == nil {
					output.PrintString(result, format)
				} else {
					log.Fatale(err)
				}

			},
		},
		{
			Name:        "set",
			Usage:       "zkcli set <path> [data]",
			Description: "Set value on znode at <path>.  If data not present on cli it's read from stdin",
			Action: func(c *cli.Context) {
				var info []byte
				if len(c.Args()) > 1 {
					info = []byte(c.Args().Get(1))
				} else {
					var err error
					info, err = ioutil.ReadAll(os.Stdin)
					if err != nil {
						log.Fatale(err)
					}
				}
				if result, err := zk.Set(c.Args().Get(0), info); err == nil {
					log.Infof("Set %+v", result)
				} else {
					log.Fatale(err)
				}

			},
		},
		{
			Name:        "create",
			Usage:       "zkcli create [command options] <path> <data>",
			Description: "Create data in znode at path",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force",
					Usage: "Force create",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					log.Fatal("Expected data argument")
				}
				if result, err := zk.Create(c.Args().First(), []byte(c.Args().Get(1)), c.Bool("force")); err == nil {
					log.Infof("Created %+v", result)
				} else {
					log.Fatale(err)
				}
			},
		},
		{
			Name:        "list",
			Aliases:     []string{"ls"},
			Usage:       "zkcli list [command options] [path]",
			Description: "list znode at [path]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "recusive, r",
					Usage: "Recusive list",
				},
			},
			Action: func(c *cli.Context) {
				x := zk.Children
				if c.Bool("recusive") {
					x = zk.ChildrenRecursive
				}
				if result, err := x(c.Args().First()); err == nil {
					output.PrintStringArray(result, format)
				} else {
					log.Fatale(err)
				}

			},
		},
		{
			Name:        "delete",
			Aliases:     []string{"del", "rm", "remove"},
			Usage:       "zkcli delete [command options] [path]",
			Description: "delete znode at [path]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force",
					Usage: "Force delete",
				},
			},
			Action: func(c *cli.Context) {
				if err := zk.Delete(c.Args().First()); err != nil {
					log.Fatale(err)
				}
			},
		},
		{
			Name:        "deleteall",
			Aliases:     []string{"rmr"},
			Usage:       "zkcli deleteall [command options] [path]",
			Description: "delete znode at [path] with all its descendants",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force",
					Usage: "Force delete",
				},
			},
			Action: func(c *cli.Context) {
				if err := zk.DeleteAll(c.Args().First()); err != nil {
					log.Fatale(err)
				}
			},
		},
	}
	app.Run(os.Args)
}
