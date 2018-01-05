package rpc

import (
	"goim/libs/proto"
	"goim/im/models"
	"goim/im/business"
)

type Conversationer interface {
	Create(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error)
	Update(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error)

	CreateMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error)
	UpdateMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error)
}

type DefaultConversationer struct {

}

func NewDefaultConversationer() *DefaultConversationer {
	return &DefaultConversationer{}
}

func (c *DefaultConversationer) Create(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error) {
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

	reply.Id = conversation.Id
	reply.Status = conversation.Status
	reply.AccountId = conversation.AccountId
	reply.Attr = conversation.Attr
	reply.Name = conversation.Name
	reply.No = conversation.No
	reply.Type = conversation.Type

	return
}

func (c *DefaultConversationer) Update(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error) {
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

	reply.Id = conversation.Id
	reply.AccountId = conversation.AccountId
	reply.No = conversation.No
	reply.Name = conversation.Name
	reply.Type = conversation.Type
	reply.Status = conversation.Status
	reply.Attr = conversation.Attr

	return
}

func (c *DefaultConversationer) CreateMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error) {
	var (
		id int64
		conversationMember *models.ConversationMember
	)

	conversationMemberEntity := &business.ConversationMemberEntity{}
	conversationMember = &models.ConversationMember{
		AccountId:arg.AccountId,
		ConvNo:arg.ConvNo,
		ClientId:arg.ClientId,
		ClientName:arg.ClientName,
		Mute:arg.Mute,
		Forbidden:arg.Forbidden,
	}

	if id, err = conversationMemberEntity.Create(conversationMember); err != nil {
		return
	}

	if conversationMember, err = conversationMemberEntity.Get(id); err != nil {
		return
	}

	reply.Id = id
	reply.AccountId = conversationMember.AccountId
	reply.ConvNo = conversationMember.ConvNo
	reply.ClientId = conversationMember.ClientId
	reply.ClientName = conversationMember.ClientName
	reply.Mute = conversationMember.Mute
	reply.Forbidden = conversationMember.Forbidden

	return
}

func (c *DefaultConversationer) UpdateMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error) {
	var (
		id int64
		conversationMember *models.ConversationMember
	)

	conversationMemberEntity := &business.ConversationMemberEntity{}
	conversationMember = &models.ConversationMember{
		Id:arg.Id,
		AccountId:arg.AccountId,
		ConvNo:arg.ConvNo,
		ClientId:arg.ClientId,
		ClientName:arg.ClientName,
		Mute:arg.Mute,
		Forbidden:arg.Forbidden,
	}

	if id, err = conversationMemberEntity.Update(conversationMember, arg.Cols); err != nil {
		return
	}

	if conversationMember, err = conversationMemberEntity.Get(id); err != nil {
		return
	}

	reply.Id = conversationMember.Id
	reply.AccountId = conversationMember.AccountId
	reply.ConvNo = conversationMember.ConvNo
	reply.ClientId = conversationMember.ClientId
	reply.ClientName = conversationMember.ClientName
	reply.Mute = conversationMember.Mute
	reply.Forbidden = conversationMember.Forbidden

	return
}