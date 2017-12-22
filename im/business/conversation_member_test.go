package business

import (
	"testing"
	"goim/im/tests"
	"goim/im/models"
	"goim/libs/crypto/rand"
	"github.com/stretchr/testify/assert"
)

func TestConversationMemberEntity_Update(t *testing.T) {
	var (
		conversationMember *models.ConversationMember
		id int64
		err error
	)
	tearDown := tests.SetUpTest(t)

	conversationMemberEntity := &ConversationMemberEntity{}
	conversationMember = &models.ConversationMember{
		AccountId:1,
		ConvNo:rand.GetRandomString(32),
		ClientId:1,
		ClientName:"testClientName",
	}

	id, err = conversationMemberEntity.Create(conversationMember)
	assert.NoError(t, err)

	conversationMember = &models.ConversationMember{
		Id:id,
		ClientName:"testClientName1",
	}

	id, err = conversationMemberEntity.Update(conversationMember, []string{"ClientName"})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	conversationMember, err = conversationMemberEntity.Get(id)

	assert.Equal(t, conversationMember.ClientName, "testClientName1")

	tearDown()
}
