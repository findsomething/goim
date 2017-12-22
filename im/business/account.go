package business

import (
	"goim/im/models"
	"goim/libs/crypto/rand"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

const (
	ENABLE = "enable"
	DISABLE = "disable"
)

type AccountEntity struct {
}

func (a *AccountEntity) Create() (id int64, err error) {
	account := &models.Account{}
	account.AccessKey = rand.GetRandomString(32)
	account.SecreteKey = rand.GetRandomString(32)
	account.Status = ENABLE
	account.CreatedTime = time.Now().Unix()
	account.UpdatedTime = time.Now().Unix()

	o := orm.NewOrm()
	id, err = o.Insert(account)

	return
}

func (a *AccountEntity) Update(account *models.Account, cols []string) (id int64, err error) {
	if account.Id <= 0 {
		return 0, fmt.Errorf("参数缺失")
	}

	account.UpdatedTime = time.Now().Unix()
	o := orm.NewOrm()

	cols = append(cols, "UpdatedTime")

	id, err = o.Update(account, cols...)
	return
}

func (a *AccountEntity) Get(id int64) (account *models.Account, err error) {
	account = &models.Account{Id:id}
	o := orm.NewOrm()
	err = o.Read(account)
	return
}