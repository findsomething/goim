package business

import (
	"goim/im/models"
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type ConversationMemberEntity struct {
}

func (c *ConversationMemberEntity) Create(conversationMember *models.ConversationMember) (id int64, err error) {
	if conversationMember.AccountId <= 0 || conversationMember.ConvNo == "" || conversationMember.ClientId <= 0 {
		return 0, fmt.Errorf("参数缺失")
	}
	conversationMember.Mute = 0
	conversationMember.Forbidden = 0
	conversationMember.UpdateTime = time.Now().Unix()
	conversationMember.CreatedTime = time.Now().Unix()

	o := orm.NewOrm()
	id, err = o.Insert(conversationMember)
	return
}

func (c *ConversationMemberEntity) Update(conversationMember *models.ConversationMember, cols []string) (id int64, err error) {
	if conversationMember.Id <= 0 {
		return 0, fmt.Errorf("参数缺失")
	}
	conversationMember.UpdateTime = time.Now().Unix()
	o := orm.NewOrm()

	id, err = o.Update(conversationMember, cols...)
	return
}

func (c *ConversationMemberEntity) Get(id int64) (conversationMember *models.ConversationMember, err error) {
	conversationMember = &models.ConversationMember{Id:id}
	o := orm.NewOrm()
	err = o.Read(conversationMember)
	return
}

