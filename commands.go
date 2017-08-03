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
  "fmt"
  "github.com/skywalkerd/amay/providers"
  "github.com/skywalkerd/amay/fileutil"
  "gopkg.in/yaml.v2"
)

// Handles the initialization of a `amay` domain
func InitAction(c *cli.Context) error {
  fmt.Println("Initiating domain")

  if GetGeneralConfig(c) != "" {
    generalConfig, err := loadConfig(GetGeneralConfig(c))

    if err == nil {
      fmt.Println("Current working domain:", generalConfig.Domain)
    }
  }

  return nil
}

func loadConfig(path string) (providers.GeneralConfig, error) {
  fmt.Println("Loading config", path)

  bytes, err := fileutil.ReadFile(path)

  if err != nil {
    fmt.Println("Can't read file")
    return providers.GeneralConfig{}, err
  }

  result := providers.GeneralConfig{}

  err = yaml.Unmarshal(bytes, &result)

  fmt.Println("result:", result.Domain)

  if err != nil {
    fmt.Println("Can't process bytes to GeneralConfig struct")
    return result, err
  } else {
    fmt.Println("Config loaded")
    return result, err
  }

}