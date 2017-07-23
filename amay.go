// amay
// Copyright (C) 2017+ skywalkerd and the project contributors
// Written by skywalkerd <andre.jacques@protonmail.com> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
  app.Copyright = "Copyright (c) 2017+ skywalkerd and the project contributors"

  app.Commands = []cli.Command {
    {
      Name: "init",
      Usage: "Initializes a new `amay` domain",
      Flags: []cli.Flag { ProviderFlag, DomainFlag, },
      Action: InitAction,
    },
  }

  app.Run(os.Args)
}