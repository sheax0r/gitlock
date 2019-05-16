package main

import (
	"log"
	"os"

	"github.com/sheax0r/gitlock/gh"

	"github.com/joeshaw/envdecode"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {

	var cfg struct {
		GithubToken string `env:"GITHUB_TOKEN,required"`
		GithubUser  string `env:"GITHUB_USER,required"`
		GithubRepo  string `env:"GITHUB_REPO,required"`
	}
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "lock",
			Usage: "lock the file.",
			Action: func(c *cli.Context) error {
				return gh.NewClient(cfg.GithubUser, cfg.GithubToken, cfg.GithubRepo).Lock()
			},
		},
		{
			Name:  "unlock",
			Usage: "Unlock the file.",
			Action: func(c *cli.Context) error {
				return gh.NewClient(cfg.GithubUser, cfg.GithubToken, cfg.GithubRepo).Unlock()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
