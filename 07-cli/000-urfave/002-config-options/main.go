package main

import (
	"context"
	"log"
	"os"

	"github.com/jllovet/learning-go/07-cli/000-urfave/002-config-options/greeting"
	altsrc "github.com/urfave/cli-altsrc/v3"
	clijson "github.com/urfave/cli-altsrc/v3/json"
	"github.com/urfave/cli/v3"
)

func main() {
	var filename string

	cmd := &cli.Command{
		Version: "v0.0",
		Name:    "Contrive",
		Commands: []*cli.Command{
			{
				Name: "hello",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						Value:       "config-en.json",
						Usage:       "filename for configurations",
						Destination: &filename,
					},
					&cli.StringFlag{
						Name:    "lang",
						Aliases: []string{"l"},
						Value:   "en",
						Usage:   "language for the greeting",
						Sources: cli.NewValueSourceChain(
							clijson.JSON("lang", altsrc.NewStringPtrSourcer(&filename)),
							clijson.JSON("lang", altsrc.StringSourcer("config.json")),
							cli.EnvVar("APP_LANG")),
					},
					cli.VersionFlag,
				},
				Action: greeting.Greet,
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}
