package rpc

import (
	"goim/libs/proto"
	"goim/im/models"
	"goim/im/business"
)

type Conversationer interface {
	Create(arg *proto.ConversationArg, replay *proto.ConversationReply) (err error)
	Update(arg *proto.ConversationArg, replay *proto.ConversationReply) (err error)
}

type DefaultConversationer struct {

}

func NewDefaultConversationer() *DefaultConversationer {
	return &DefaultConversationer{}
}

func (c *DefaultConversationer) Create(arg *proto.ConversationArg, replay *proto.ConversationReply) (err error) {
	var (
		id int64
		conversation *models.Conversation
	)

	conversationEntity := business.ConversationEntity{}
	conversation = &models.Conversation{
		AccountId:arg.AccountId,
		Name:arg.Name,
		Type:arg.Type,
		Status:arg.Status,
		Attr:arg.Attr,
	}
	if id, err = conversationEntity.Create(conversation); err != nil {
		return
	}

	conversation = &models.Conversation{
		Id:id,
	}

	if conversation, err = conversationEntity.Get(id); err != nil {
		return
	}

	replay.Id = conversation.Id
	replay.Status = conversation.Status
	replay.AccountId = conversation.AccountId
	replay.Attr = conversation.Attr
	replay.Name = conversation.Name
	replay.No = conversation.No
	replay.Type = conversation.Type

	return
}

func (c *DefaultConversationer) Update(arg *proto.ConversationArg, replay *proto.ConversationReply) (err error) {
	var (
		id int64
		conversation *models.Conversation
	)

	conversation = &models.Conversation{
		Id:arg.Id,
		AccountId:arg.AccountId,
		Name:arg.Name,
		Type:arg.Type,
		Status:arg.Status,
		Attr:arg.Attr,
	}
	conversationEntity := &business.ConversationEntity{}
	if id, err = conversationEntity.Update(conversation, arg.Cols); err != nil {
		return
	}

	if conversation, err = conversationEntity.Get(id); err != nil {
		return
	}

	replay.Id = conversation.Id
	replay.AccountId = conversation.AccountId
	replay.No = conversation.No
	replay.Name = conversation.Name
	replay.Type = conversation.Type
	replay.Status = conversation.Status
	replay.Attr = conversation.Attr

	return
}