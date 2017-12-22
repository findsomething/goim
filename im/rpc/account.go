package rpc

import (
	"goim/libs/proto"
	"goim/im/business"
	"goim/im/models"
)

type Accounter interface {
	Create(arg *proto.NoArg, replay *proto.AccountReply) (err error)
	Update(arg *proto.AccountArg, replay *proto.AccountReply) (err error)
}

type DefaultAccounter struct {

}

func NewDefaultAccounter() *DefaultAccounter {
	return &DefaultAccounter{}
}

func (a *DefaultAccounter) Create(arg *proto.NoArg, replay *proto.AccountReply) (err error) {
	var (
		id int64
		account *models.Account
	)

	accountEntity := &business.AccountEntity{}
	id, err = accountEntity.Create()
	if err != nil {
		return
	}
	account, err = accountEntity.Get(id)

	replay.Id = account.Id
	replay.AccessKey = account.SecreteKey
	replay.SecretKey = account.SecreteKey
	replay.Status = account.Status

	return
}

func (a *DefaultAccounter) Update(arg *proto.AccountArg, replay *proto.AccountReply) (err error) {
	var (
		id int64
		account *models.Account
	)

	account = &models.Account{
		Id:arg.Id,
		Status:arg.Status,
	}
	accountEntity := &business.AccountEntity{}
	id, err = accountEntity.Update(account, arg.Cols)

	account, err = accountEntity.Get(id)

	replay.Id = account.Id
	replay.Status = account.Status

	return
}