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

// An IProvider manage virtual machine
type IProvider interface {
  Create() error
  Destroy() error
  Launch() error
  Halt() error
  Status() (string, error)
}

// A provider for `amay` must have a unique full domain, i.e. subdomain + domain should be unique in an environment
type Provider struct {
  name string
  domain string
  subDomain string
}

// Gets the provider's name
func (p Provider) GetName() string {
  return p.name
}

// Sets the provider's name
func (p Provider) SetName(name string) {
  p.name = name
}

// Gets the provider's domain
func (p Provider) GetDomain() string {
  return p.domain
}

// Sets the provider's domain
func (p Provider) SetDomain(domain string) {
  p.domain = domain
}

// Gets the provider's sub domain
func (p Provider) GetSubDomain() string {
  return p.subDomain
}

// Sets the provider's sub domain
func (p Provider) SetSubDomain(subDomain string) {
  p.subDomain = subDomain
}