package rpc

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"goim/libs/business"
)

func TestDefaultAuther_Auth(t *testing.T) {
	var (
		accountId int64 = 1
		clientId int64 = 100
	)
	auther := NewDefaultAuther("testKey")

	token := auther.CreateToken(accountId, clientId)

	preferAccountId, _, err := auther.Auth(token)

	assert.NoError(t, err)

	assert.Equal(t, preferAccountId, business.UniqueUserId(accountId, clientId))
}