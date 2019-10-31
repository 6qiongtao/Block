package merror

import (
	"time"

	"v_tools/email"

	"github.com/astaxie/beego"

	"vtoken_digiccy_go/test/config"
)

func SendEmail(title, content, touser string) {
	//TODO:
	runMode := beego.BConfig.RunMode
	go func() {
		for _, toUser := range config.EmailList {
			time.Sleep(time.Duration(200) * time.Millisecond)
			email.Send(runMode+title, content, toUser, config.GetGlobalMap().AlertEmailMerchant)
		}
	}()
}
