// Copyright 2020 Simen A. W. Olsen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"

	"github.com/cobraz/jira-to-tripletex/pkg/jtt"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Jira-To-Tripletex",
		Description: "Transport Jira issues to Tripletex",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "get:config",
				Aliases: []string{"gc"},
				Action:  jtt.GetConfig,
			},
			&cli.Command{
				Name:    "set:config",
				Aliases: []string{"cnf"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Required: true,
						Usage:    "Jira server host",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "jiraEmail",
						Required: true,
						Usage:    "Jira Email",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "jiraToken",
						Required: true,
						Usage:    "Jira API token",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "consumerToken",
						Required: true,
						Usage:    "Tripletex ConsumerToken",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "employeeToken",
						Required: true,
						Usage:    "Tripletex EmployeeToken",
						// Destination: &language,
					},
					&cli.StringFlag{
						Name:     "activityCode",
						Required: true,
						Usage:    "Tripletex Activity Id",
						// Destination: &language,
					},
				},
				Action: jtt.SetConfig,
			},
			&cli.Command{
				Name:        "list:activities",
				Description: "Lists all activities in Tripletex",
				Action: func(c *cli.Context) error {
					return jtt.GetActivities()
				},
			},
			&cli.Command{
				Name:        "copy",
				Description: "Copy project to Tripletex",
				Aliases:     []string{"c"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "projectId",
						Required: true,
						Usage:    "Tripletex Project ID",
					},
					&cli.StringFlag{
						Name:     "key",
						Required: true,
						Usage:    "Jira Project Key",
					},
				},
				Action: func(c *cli.Context) error {
					return jtt.CopyToTripletex(c.String("key"), int32(c.Int("projectId")))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
