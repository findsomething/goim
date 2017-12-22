package tests

import (
	"testing"
	"goim/im/tool"
	"goim/im/models"
	"github.com/astaxie/beego/orm"
)

var (
	Conf *tool.Config
	err error
)

func SetUpTest(t *testing.T) func() {

	fileName := "/Users/lihan/mygo/src/goim/im/im-dev.conf"

	Conf, err = tool.InitConfig(fileName)

	models.InitMysql(Conf.RPCConf)
	clearDb()

	return func() {
	}
}

func clearDb() {
	var (
		tables []string
		rawSeter orm.RawSeter
	)
	o := orm.NewOrm()
	sql := "SELECT table_name FROM information_schema.tables WHERE table_schema = ?"
	rawSeter = o.Raw(sql, Conf.RPCConf.MysqlName)
	rawSeter.QueryRows(&tables)
	for _, table := range tables {
		if table != "migrations" {
			o.Raw("TRUNCATE "+table).Exec()
		}
	}
}