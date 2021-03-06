package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameConf_Regular(t *testing.T) {
	var (
		confAppTmp = TConfigFile{
			PathDirConf:  ".",
			NameFileConf: "sample_config",
			NameTypeConf: "json",
		}
		expect = "sample_config.json"
		actual = confAppTmp.GetNameConf()
	)

	assert.Equal(t, expect, actual, "Returned unexpected value from the conf.")
}

func TestGetNameConf_WithExtInName(t *testing.T) {
	var (
		confApp TConfigFile
		expect  string
		actual  string
	)

	// User defined conf file path
	confApp = TConfigFile{
		PathFileConf: "./foo/bar.json", // Value from the command flag
		PathDirConf:  ".",              // This should be ignored
		NameFileConf: "sample_config",  // This should be ignored
		NameTypeConf: "json",           // This should be ignored
	}
	expect = "bar.json"
	actual = confApp.GetNameConf()
	assert.Equal(t, expect, actual, "If the conf name has an extension it should return as is.")

	// App defined conf file
	confApp = TConfigFile{
		PathFileConf: "",
		PathDirConf:  ".",
		NameFileConf: "foo.yaml",
		NameTypeConf: "json", // This should be ignored
	}
	expect = "foo.yaml"
	actual = confApp.GetNameConf()
	assert.Equal(t, expect, actual, "If the conf name has an extension it should return as is.")
}

func Test_hasExtInName_Failure(t *testing.T) {
	var (
		nameExt string = "sample_conf.html"
		expect  bool   = false
		actual  bool   = hasExtInName(nameExt)
	)

	assert.Equal(t, expect, actual, "Un-available extension for Viper should return false.")
}

func Test_hasExtInName_Regular(t *testing.T) {
	var (
		nameFile string = "sample_conf.json"
		expect   bool   = true
		actual   bool   = hasExtInName(nameFile)
	)

	assert.Equal(t, expect, actual, "Available extension for Viper should return true.")
}

func TestLoadFile_Failure(t *testing.T) {
	var (
		ConfApp = TConfigFile{
			PathFileConf: "./foobar.json",
			PathDirConf:  "..",
			NameFileConf: "sample_config",
			NameTypeConf: "json",
		}

		ConfUser = struct {
			NameToGreet string
		}{
			NameToGreet: "bar",
		}
	)

	err := LoadFile(ConfApp, ConfUser)
	assert.Error(t, err, "When un-existing path was provided it should return an error.")
}

func TestLoadFile_Regular(t *testing.T) {
	var (
		confApp  TConfigFile
		confUser struct {
			NameToGreet string `mapstructure:"name_to_greet"`
		}
		err    error
		expect string
		actual string
	)

	// User defined config file
	confApp = TConfigFile{
		PathFileConf: "../sample_config.json",
	}
	confUser.NameToGreet = "bar"

	if err = LoadFile(confApp, &confUser); err != nil {
		t.Fatalf("Failed to read conf file for test.\nError msg: %v", err)
	}

	expect = "Cobra"
	actual = confUser.NameToGreet
	assert.Equal(t, expect, actual, "Returned unexpected value from the conf.")

	// App defined config file
	confApp = TConfigFile{
		PathFileConf: "",
		PathDirConf:  "..",
		NameFileConf: "sample_config",
		NameTypeConf: "json",
	}
	confUser.NameToGreet = "bar"

	if err = LoadFile(confApp, &confUser); err != nil {
		t.Fatalf("Failed to read conf file for test.\nError msg: %v", err)
	}

	expect = "Cobra"
	actual = confUser.NameToGreet
	assert.Equal(t, expect, actual, "Returned unexpected value from the conf.")
}
