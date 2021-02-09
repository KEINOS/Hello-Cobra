/*
Package conf config.go defines functions to read file configuration using Viper.

It was separated to a different package from `cmd` to ease testing and re-use.

- This package was very much taught by the below articles:
  - "[Backend #12] Load config from file & environment variables in Golang with Viper"
	- https://youtu.be/n5p8HkO6bnE by TECH SCHOOL @ YouTube
  - "cobra-viper-example"
  	- https://github.com/nirasan/cobra-viper-example @ GitHub

*/
package conf

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// TypeConfigApp defines the data structure to store basic config of the app.
type TypeConfigApp struct {
	PathFileConf string // File path of config. If set, will have priority than NameFileConf and PathDirConf.
	PathDirConf  string // Dir path of config to search.
	NameFileConf string // File name of config file. May or may not have an extension.
	// File extension of the config file. REQUIRED if the conf file does not have the extension in the name.
	NameTypeConf string

	IsUsingDefaultConf bool // Flag to determine if the app is using the default value or conf file value.
}

// GetNameConf is a method of TypeConfigApp that returns the config file name.
func (c TypeConfigApp) GetNameConf() string {
	if "" != c.PathFileConf {
		return filepath.Base(c.PathFileConf)
	}

	if hasExtInName(c.NameFileConf) {
		return c.NameFileConf
	}

	return c.NameFileConf + "." + c.NameTypeConf
}

// LoadConfig stores values from the config file or env variables to userConfig.
//
// @args appConfig  TypeConfigApp  : Basic application config to read the conf file.
//
// @args userConfig struct         : An object to be stored the values from conf file.
//
// @return err      error          : nil if success. Error msg if fails to read/store values from conf file.
func LoadConfig(appConfig TypeConfigApp, userConfig interface{}) (err error) {
	// Reset current stored values in viper
	viper.Reset()

	// Read values from the ENV variable if set
	viper.AutomaticEnv()

	// Set file path to search
	if "" != appConfig.PathFileConf {
		// Set one-liner file path
		viper.SetConfigFile(appConfig.PathFileConf)
	} else {
		// Set inidividual file path info
		viper.AddConfigPath(appConfig.PathDirConf)
		viper.SetConfigName(appConfig.NameFileConf)
		viper.SetConfigType(appConfig.NameTypeConf)
	}

	// Search and read values from the config file.
	err = viper.ReadInConfig()
	if err == nil {
		// Map the read config values into "userConfig" object
		err = viper.Unmarshal(&userConfig)
	}

	return err // return error if viper fails to read or map the values
}

// hasExtInName returns true if the nameFile contains a file extension which viper can detect.
func hasExtInName(nameFile string) bool {
	var extWithNoDot string = strings.TrimLeft(filepath.Ext(nameFile), ".")

	return hasStringInSlice(extWithNoDot, viper.SupportedExts)
}

// hasStringInSlice returns true if a string is in a list of slice.
func hasStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
