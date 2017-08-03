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

import "github.com/skywalkerd/amay/fileutil"

const (
  imageField = "image"
  boxUrlField = "boxUrl"
  folderHostField = "folderHost"
  folderVmField = "folderVm"
  provisionFilePathField = "provisionFilePath"
  domainField = "domain"
  subDomainField = "subDomain"
  cpuCountField = "cpuCount"
  ramAmountField = "ramAmount"
)

// Implementation of the Vagrant provider
type VagrantProvider struct {
  Provider
  fileutil.Template

  image, boxUrl, folderHost, folderVm, provisionFilePath, domain, subDomain string
  cpuCount, ramAmount int
}

// Gets the vagrant's image name
func (v VagrantProvider) Image() string {
  return v.image
}

// Sets the vagrant's image name
func (v VagrantProvider) SetImage(image string) {
  v.image = image
}

func (v VagrantProvider) BoxUrl() string {
  return v.boxUrl
}

func (v VagrantProvider) SetBoxUrl(boxUrl string) {
  v.boxUrl = boxUrl
}

func (v VagrantProvider) FolderHost() string {
  return v.folderHost
}

func (v VagrantProvider) SetFolderHost(folderHost string) {
  v.folderHost = folderHost
}

func (v VagrantProvider) FolderVm() string {
  return v.folderVm
}

func (v VagrantProvider) SetFolderVm(folderVm string) {
  v.folderVm = folderVm
}

func (v VagrantProvider) ProvisionFilePath() string {
  return v.provisionFilePath
}

func (v VagrantProvider) SetProvisionFilePath(provisionfilePath string) {
  v.provisionFilePath = provisionfilePath
}

func (v VagrantProvider) Domain() string {
  return v.domain
}

func (v VagrantProvider) SetDomain(domain string) {
  v.domain = domain
}

func (v VagrantProvider) SubDomain() string {
  return v.subDomain
}

func (v VagrantProvider) SetSubDomain(subDomain string) {
  v.subDomain = subDomain
}

// Gets the vagrant's cpu(s) count
func (v VagrantProvider) CpuCount() int {
  return v.cpuCount
}

// Sets the vagrant's cpu(s) count
func (v VagrantProvider) SetCpuCount(cpuCount int) {
  v.cpuCount = cpuCount
}

// Gets the amount of RAM memory
func (v VagrantProvider) RamAmount() int {
  return v.ramAmount
}

// Sets the amount of RAM memory
func (v VagrantProvider) SetRamAmount(ramAmount int) {
  v.ramAmount = ramAmount
}

func (v VagrantProvider) Apply() {
  v.AddPattern(imageField, v.Image())
  v.AddPattern(boxUrlField, v.BoxUrl())
  v.AddPattern(folderHostField, v.FolderHost())
  v.AddPattern(folderVmField, v.FolderVm())
  v.AddPattern(provisionFilePathField, v.ProvisionFilePath())
  v.AddPattern(domainField, v.Domain())
  v.AddPattern(subDomainField, v.SubDomain())
  v.AddPattern(cpuCountField, string(v.CpuCount()))
  v.AddPattern(ramAmountField, string(v.RamAmount()))
}