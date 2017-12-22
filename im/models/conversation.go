package models

type Conversation struct {
	Id int64
	AccountId int64 `orm:"column(accountId)"`
	No string `orm:"column(no)"`
	Name string `orm:"column(name)"`
	Type string `orm:"column(type)"`
	Status string `orm:"column(status)"`
	Attr string `orm:"column(attr)"`
	CreatedTime int64 `orm:"column(createdTime)"`
	UpdatedTime  int64 `orm:"column(updatedTime)"`
}

func (c *Conversation) TableName() string {
	return "conversation"
}