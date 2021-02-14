# Package conf

Directory of file-configuration package using [Viper](https://github.com/spf13/viper).

- Note that this package may be used as JSON/YAML file reader but it's not designed to load MEGA sized file.

## Sample usage in other package

Pretend the user config file (`userConfig.json`) was as below:

```json
{
    "my_value": "Cobra"
}
```

Then the Go sample to load it would be as below:

```go
package cmd

import "github.com/KEINOS/Hello-Cobra/conf"

// TDataUser defines the data structure to be stored from the conf file read.
type TDataUser struct {
    // MyValue is the variable to store the value of "my_value" key in the conf file.
    MyValue string `mapstructure:"my_value"`
}

var (
    configFile = conf.TConfigApp{
        PathDirConf:        ".",
        NameFileConf:       "userConfig",
        NameTypeConf:       "json",
    }

    userValues = TDataUser{
        MyValue: "default value",
    }
)

// Load value from "./userConfig.json" to `userValues`
if err := conf.LoadConfig(*configFile, &userValues); err != nil {
    // do something with the error
}

// Use loaded value from the conf file
myValue := userValues.MyValue

// `myValue` expects to be "Cobra"

```
