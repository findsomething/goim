package business

import (
	"goim/im/models"
	"fmt"
	"goim/libs/crypto/rand"
	"time"
	"github.com/astaxie/beego/orm"
)

type ConversationEntity struct {
}

func (c *ConversationEntity) Create(conversation *models.Conversation) (id int64, err error) {
	if conversation.AccountId <= 0 {
		return 0, fmt.Errorf("缺失账号Id")
	}
	conversation.No = rand.GetRandomString(32)
	conversation.Status = "ok"
	conversation.CreatedTime = time.Now().Unix()
	conversation.UpdatedTime = time.Now().Unix()

	o := orm.NewOrm()
	id, err = o.Insert(conversation)
	return
}

func (c *ConversationEntity) Update(conversation *models.Conversation, cols []string) (id int64, err error) {
	if conversation.Id <= 0 {
		return 0, fmt.Errorf("参数缺失")
	}

	conversation.UpdatedTime = time.Now().Unix()
	o := orm.NewOrm()

	cols = append(cols, "UpdatedTime")

	id, err = o.Update(conversation, cols...)
	return
}

func (c *ConversationEntity) Get(id int64) (conversation *models.Conversation, err error) {
	conversation = &models.Conversation{Id:id}
	o := orm.NewOrm()
	err = o.Read(conversation)
	return
}