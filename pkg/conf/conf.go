package conf

import (
	"time"

	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigFile("./pkg/conf/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&App); err != nil {
		panic(err)
	}
}

var App goligoli

type goligoli struct {
	JWT          jwt
	UserService  service
	VideoService service
	Etcd         etcd
	Gateway      gateway
	Res          res
}

type jwt struct {
	Secret        string
	Issuer        string
	Exp           time.Duration
	SigningMethod string
}

type service struct {
	Name    string
	RpcPort string
	AppID   string
	AppKey  string
	MySQL   mysql
}

type mysql struct {
	Dsn string
}

type etcd struct {
	URL string
}

type gateway struct {
	Port string
}

type res struct {
	VideoDir string
	PhotoDir string
}
