package utils

import (
	"os"

	"gopkg.in/ini.v1"
)

func GetIni(section, key, defaultvalue string) string {
	cfg, err := ini.Load("config.ini")
	if (err  != nil) {
		Errorf("Fail to read the file: ", err)
		os.Exit(1)
	}

	if value := cfg.Section(section).Key(key).String(); value != "" {
		return value
	}

	return defaultvalue
}