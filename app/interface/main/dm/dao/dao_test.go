package dao

import (
	"flag"
	"os"
	"testing"

	"go-common/app/interface/main/dm/conf"
	"go-common/library/log"
)

var testDao *Dao

func TestMain(m *testing.M) {
	flag.Set("app_id", "main.community.dm")
	flag.Set("conf_token", "4d343dc5d8c33460156baec8913eb64a")
	flag.Set("tree_id", "2692")
	flag.Set("conf_version", "docker-1")
	flag.Set("deploy_env", "uat")
	flag.Set("conf_host", "config.bilibili.co")
	flag.Set("conf_path", "/tmp")
	flag.Set("region", "sh")
	flag.Set("zone", "sh001")
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Xlog)
	defer log.Close()
	testDao = New(conf.Conf)
	os.Exit(m.Run())
}
