package cache

import (

	mylog "vtoken_digiccy_go/test/tools/log"

)

var (
	ApiLog *mylog.Log
	ErrLog *mylog.Log
)
func init() {
	ApiLog = mylog.Init("20060102.api")
	ErrLog = mylog.Init("20060102.err")
}
