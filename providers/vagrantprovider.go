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

package providers

// Implementation of the Vagrant provider
type VagrantProvider struct {
  Provider

  image string
  cpuCount int
  ramAmount int
}

// Gets the vagrant's image name
func (v VagrantProvider) GetImage() string {
  return v.image
}

// Sets the vagrant's image name
func (v VagrantProvider) SetImage(image string) {
  v.image = image
}

// Gets the vagrant's cpu(s) count
func (v VagrantProvider) GetCpuCount() int {
  return v.cpuCount
}

// Sets the vagrant's cpu(s) count
func (v VagrantProvider) SetCpuCount(cpuCount int) {
  v.cpuCount = cpuCount
}

// Gets the amount of RAM memory
func (v VagrantProvider) GetRamAmount() int {
  return v.ramAmount
}

// Sets the amount of RAM memory
func (v VagrantProvider) SetRamAmount(ramAmount int) {
  v.ramAmount = ramAmount
}