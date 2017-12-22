package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Init_20171219_173122 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Init_20171219_173122{}
	m.Created = "20171219_173122"
	migration.Register("Init_20171219_173122", m)
}

// Run the migrations
func (m *Init_20171219_173122) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `account` ( " +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`status` enum('enable','disable') NOT NULL DEFAULT 'enable' COMMENT '账号状态'," +
		"`accessKey` varchar(32) NOT NULL DEFAULT '' COMMENT 'accessKey'," +
		"`secretKey` varchar(32) NOT NULL DEFAULT '' COMMENT 'secretKey'," +
		"`createdTime` int(10) DEFAULT '0' COMMENT '创建时间'," +
		"`updatedTime` int(10) DEFAULT '0' COMMENT '更新时间'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;")

	m.SQL("CREATE TABLE `conversation` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`accountId` int(11) NOT NULL DEFAULT 0 COMMENT '账号Id'," +
		"`no` varchar(32) NOT NULL DEFAULT '' COMMENT '会话no'," +
		"`name` varchar(128) NOT NULL DEFAULT '' COMMENT '会话名称'," +
		"`type` enum('system','normal','transport','command') NOT NULL DEFAULT 'normal' COMMENT '会话类型'," +
		"`status` enum('ok','drop') NOT NULL DEFAULT 'ok' COMMENT '会话状态'," +
		"`attr` text NOT NULL COMMENT '额外属性'," +
		"`updatedTime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间'," +
		"`createdTime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间'," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `accountId` (`accountId`,`no`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8")

	m.SQL("CREATE TABLE `conversation_member` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`accountId` int(11) NOT NULL DEFAULT 0 COMMENT '账号Id'," +
		"`convNo` varchar(32) NOT NULL DEFAULT '' COMMENT '会话no'," +
		"`clientId` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端Id'," +
		"`clientName` varchar(128) NOT NULL DEFAULT '' COMMENT '客户端昵称'," +
		"`mute` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1静音/0非静音'," +
		"`forbidden` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1禁止/0非禁止发言'," +
		"`updatedTime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间'," +
		"`createdTime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间'," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `accountId` (`accountId`,`convNo`,`clientId`)," +
		"KEY `accountId_2` (`accountId`,`clientId`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *Init_20171219_173122) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `account`;")

	m.SQL("DROP TABLE `conversation`;")

	m.SQL("DROP TABLE `conversation_member`;")
}
