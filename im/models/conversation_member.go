package models

type ConversationMember struct {
	Id int64
	AccountId int64 `orm:"column(accountId)"`
	ConvNo string `orm:"column(convNo)"`
	ClientId int64 `orm:"column(clientId)"`
	ClientName string `orm:"column(clientName)"`
	Mute int32 `orm:"column(mute)"`
	Forbidden int32 `orm:"column(forbidden)"`
	CreatedTime int64 `orm:"column(createdTime)"`
	UpdateTime  int64 `orm:"column(updatedTime)"`
}

func (c *ConversationMember) TableName() string {
	return "conversation_member"
}