package main

import "github.com/urfave/cli"

var (
  DomainFlag = cli.StringFlag{
    Name: "domain, d",
    Usage: "Specifies the targeted domain",
  }
)
