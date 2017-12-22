package business

import (
	"testing"
	"goim/im/tests"
	"goim/im/models"
	"github.com/stretchr/testify/assert"
)

func TestConversationEntity_Update(t *testing.T) {
	var (
		conversation *models.Conversation
		id int64
		err error
	)

	tearDown := tests.SetUpTest(t)

	account := &AccountEntity{}
	accountId, _ := account.Create()

	conversationEntity := &ConversationEntity{}

	conversation = &models.Conversation{
		AccountId:accountId,
		Name:"testConversation",
		Type:"normal",
	}

	id, err = conversationEntity.Create(conversation)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	conversation = &models.Conversation{
		Id:id,
		Name:"testConversation1",
	}

	id, err = conversationEntity.Update(conversation, []string{"Name"})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	conversation, err = conversationEntity.Get(id)

	assert.Equal(t, conversation.Name, "testConversation1")

	tearDown()
}