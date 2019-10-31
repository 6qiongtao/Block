package config

import (
	"encoding/json"
	"fmt"

	"vtoken_digiccy_go/route/tools"

	consulApi "github.com/hashicorp/consul/api"
	log "github.com/micro/go-micro/util/log"
)

var GlobalMap Global

//从本地json读取 consul报警邮件配置
type ConSulConfig struct {
	//告警商户
	EmailMerchant string
	//告警内容
	AlertEmail    string
	//consul 地址
	ConsulAddr    string

}

var ConsulKV *consulApi.KV

func init()  {
	consulApiInit()

	//初始化consul配置从本地json文件中
	data, err := tools.ReadFile("conf/" + RunMode + ".config.json")
	if err != nil {
		panic(err)
	}
	//从本地json读取配置
	err = json.Unmarshal(data, &ConsulCfg)
	//fmt.Println("init consul 配置:", string(data))
	if err != nil {
		panic(err)
	}

	//从consul web读取json k/v
	ConsulUnmarshal("vtoken_digiccy_go/route/global", &GlobalMap)
	fmt.Printf("%+v \n", GlobalMap)
}

//init consul
func consulApiInit() *consulApi.KV {
	conf := consulApi.DefaultConfig()
	conf.Address = ConsulCfg.ConsulAddr
	client, err := consulApi.NewClient(conf)
	if err != nil {
		log.Info("consulApi NewClient err:", err)
	}
	ConsulKV = client.KV()
	return ConsulKV
}

//Alert Email
type Global struct {
	AlertEmailToUser   string
	AlertEmailUrl      string
	AlertEmailMerchant string

}

func GetGlobalMap() Global {
	return GlobalMap
}

//从consul web读取json k/v
func ConsulUnmarshal(path string, config interface{}) (err error) {
	return ;
	pair, _, err := ConsulKV.Get("cfg", nil)
	if pair == nil || err != nil {
		log.Log("ConsulUnmarshal err:", err)
		//cache.ErrLog("docker-compose environment AppLogLevel")
		//merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)
		return err
	}
	err = json.Unmarshal(pair.Value, &GlobalMap)

	if err != nil {
		log.Log("ConsulUnmarshal err:", err)
		//cache.ErrLog("docker-compose environment AppLogLevel")
		//merror.Log(merror.LogLevel, "URI:", r.Host, r.URL, "err:",  err.Error(), "code:",500)

		return err
	}
	return
}
