package main

import (
  "github.com/urfave/cli"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "amay"
  app.Usage = "Little CLI software to handle development environment"
  app.Version = "0.1.0-SNAPSHOT"

  app.Commands = []cli.Command {
    {
      Name: "init",
      Usage: "Initializes a new `amay` domain",
      Flags: []cli.Flag { DomainFlag, },
      Action: InitAction,
    },
  }

  app.Run(os.Args)
}