package models

type Account struct {
	Id          int64
	Status      string `orm:"column(status)"`
	AccessKey   string `orm:"column(accessKey)"`
	SecreteKey  string `orm:"column(secretKey)"`
	CreatedTime int64 `orm:"column(createdTime)"`
	UpdatedTime  int64 `orm:"column(updatedTime)"`
	Base `orm:"-"`
}

func (a *Account) TableName() string {
	return "account"
}