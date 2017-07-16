package main

import (
  "github.com/urfave/cli"
  "fmt"
)

func InitAction(_c *cli.Context) error {
  fmt.Println("Initiating domain")

  if _c.String("domain") != "" {
    fmt.Println("Domain is", _c.String("domain"))
  }
  return nil
}
