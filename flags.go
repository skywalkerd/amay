// amay
// Copyright (C) 2017+ skywalkerd and the project contributors
// Written by A. Jacques <andre.jacques@protonmail.com> and the project contributors
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

import "github.com/urfave/cli"

var (
  ProviderFlag = cli.StringFlag {
    Name: "provider, p",
    Usage: "Specifies the targeted provider",
  }

  DomainFlag = cli.StringFlag {
    Name: "domain, d",
    Usage: "Specifies the targeted domain",
  }
)

// Gets the provider flag value
func GetProvider(c *cli.Context) string {
  if c.String("provider") != "" {
    return c.String("provider")
  }

  return ""
}

// Gets the domain flag value
func GetDomain(c *cli.Context) string {
  if c.String("domain") != "" {
    return c.String("domain")
  }

  return ""
}