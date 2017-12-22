package rpc

import (
	"testing"
	"goim/im/tests"
	"goim/libs/proto"
	"github.com/stretchr/testify/assert"
	"goim/im/business"
)

func TestDefaultAccounter_Update(t *testing.T) {

	tearDown := tests.SetUpTest(t)

	defaultAccounter := NewDefaultAccounter()
	arg := &proto.NoArg{}
	replay := &proto.AccountReply{}

	err := defaultAccounter.Create(arg, replay)

	assert.NoError(t, err)
	assert.NotEmpty(t, replay.AccessKey)

	accountArg := &proto.AccountArg{
		Id:replay.Id,
		Status:business.DISABLE,
		Cols:[]string{"Status"},
	}

	err = defaultAccounter.Update(accountArg, replay)

	assert.NoError(t, err)
	assert.Equal(t, replay.Status, business.DISABLE)

	tearDown()
}