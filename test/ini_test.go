package test

import (
	"testing"

	"github.com/DmitryVesenniy/goconfig/ini"
)

func TestGetConfig(t *testing.T) {
	config := &Config{}
	getConfig := ini.Get(config, "settings.txt")

	config = (getConfig()).(*Config)

	if config.PathList != "/path/to/file" {
		t.Error("[!] error config.PathList")
	}

	if !config.SkipExist {
		t.Error("[!] error config.SkipExist")
	}
}

type Config struct {
	PathList  string `ini:"PATH"`
	SkipExist bool   `ini:"SKIP"`
}
