package config

import (
	"fmt"
	"os"
)

var AppLogLevel string

func init() {
	AppLogLevel = os.Getenv("APP_LOG_LEVEL")
	fmt.Println(AppLogLevel)
}
