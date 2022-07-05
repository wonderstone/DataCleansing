package datacleaner

import (
	"fmt"
	"github.com/spf13/viper"
)

// ReadData function has a string parameter, which is the path to the data file.

func ReadData(dir string) string {
	viper.SetConfigName("BackTest")
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, World!")
	return "Hello, World!"
}
