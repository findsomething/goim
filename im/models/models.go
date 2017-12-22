package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goim/im/tool"
)

func InitMysql(rpcConf tool.RpcConfig) {

	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s/%s?charset=utf8", rpcConf.MysqlAddr, rpcConf.MysqlName),
		int(rpcConf.MysqlMaxIdleConns), int(rpcConf.MysqlMaxOpenConns))

	orm.RegisterModel(new(Account))
	orm.RegisterModel(new(Conversation))
	orm.RegisterModel(new(ConversationMember))
}
