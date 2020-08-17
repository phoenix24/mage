package configs

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

func ReadConfig(path string) ProxyConfig {
	viper.SetConfigType("yaml")

	var config ProxyConfig
	var content, _ = ioutil.ReadFile(path)
	var _ = viper.ReadConfig(bytes.NewBuffer(content))
	if  err := viper.Unmarshal(&config); err != nil {
		fmt.Println("failed to read service config: ", err)
	}
	return config
}
