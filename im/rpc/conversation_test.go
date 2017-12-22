package rpc

import (
	"testing"
	"goim/im/tests"
	"goim/libs/proto"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConversationer_Update(t *testing.T) {
	tearDown := tests.SetUpTest(t)

	defaultConversationer := NewDefaultConversationer()
	arg := &proto.ConversationArg{
		AccountId:1,
		Name:"testConversation",
		Type:"normal",
	}

	replay := &proto.ConversationReply{}

	err := defaultConversationer.Create(arg, replay)

	assert.NoError(t, err)
	assert.Equal(t, replay.Name, "testConversation")

	arg = &proto.ConversationArg{
		Id:replay.Id,
		Name:"testConversation1",
		Cols:[]string{"Name"},
	}
	replay = &proto.ConversationReply{}

	err = defaultConversationer.Update(arg, replay)

	assert.NoError(t, err)
	assert.Equal(t, replay.Name, "testConversation1")

	tearDown()
}