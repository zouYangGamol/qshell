package qshell

import (
	"qiniu/api.v6/conf"
)

type ZoneConfig struct {
	UpHost    string
	RsHost    string
	RsfHost   string
	IovipHost string
	ApiHost   string
}

var (
	DEFAULT_API_HOST = ZoneNBConfig.ApiHost
)

const (
	ZoneNB  = "nb"
	ZoneBC  = "bc"
	ZoneAWS = "aws"
	ZoneNA0 = "na0"
)

//zone all defaults to the service source site

var ZoneNBConfig = ZoneConfig{
	UpHost:    "http://up.qiniu.com",
	RsHost:    "http://rs.qiniu.com",
	RsfHost:   "http://rsf.qbox.me",
	IovipHost: "http://iovip.qbox.me",
	ApiHost:   "http://api.qiniu.com",
}

var ZoneBCConfig = ZoneConfig{
	UpHost:    "http://up-z1.qiniu.com",
	RsHost:    "http://rs.qiniu.com",
	RsfHost:   "http://rsf-z1.qbox.me",
	IovipHost: "http://iovip-z1.qbox.me",
	ApiHost:   "http://api-z1.qiniu.com",
}

var ZoneAWSConfig = ZoneConfig{
	UpHost:    "http://up.gdipper.com",
	RsHost:    "http://rs.gdipper.com",
	RsfHost:   "http://rsf.gdipper.com",
	IovipHost: "http://io.gdipper.com",
	ApiHost:   "http://api.gdipper.com",
}

var ZoneNA0Config = ZoneConfig{
	UpHost:    "http://upload-na0.qiniu.com",
	RsHost:    "http://rs-na0.qbox.me",
	RsfHost:   "http://rsf-na0.qbox.me",
	IovipHost: "http://iovip-na0.qbox.me",
	ApiHost:   "http://api-na0.qiniu.com",
}

func SetZone(zoneConfig ZoneConfig) {
	conf.UP_HOST = zoneConfig.UpHost
	conf.RS_HOST = zoneConfig.RsHost
	conf.RSF_HOST = zoneConfig.RsfHost
	conf.IO_HOST = zoneConfig.IovipHost
	DEFAULT_API_HOST = zoneConfig.ApiHost
}

func IsValidZone(zone string) (valid bool) {
	switch zone {
	case ZoneNB, ZoneBC, ZoneAWS, ZoneNA0:
		valid = true
	default:
		valid = false
	}
	return
}
