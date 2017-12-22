package business

import (
	"testing"
	"goim/im/tests"
	"github.com/stretchr/testify/assert"
	"goim/im/models"
)

func TestCreate(t *testing.T) {
	tearDown := tests.SetUpTest(t)

	account := &AccountEntity{}

	id, err := account.Create()

	if id <= 0 || err != nil {
		t.Error("创建账号失败")
	}

	tearDown()
}

func TestUpdate(t *testing.T) {
	var err error
	tearDown := tests.SetUpTest(t)

	account := &AccountEntity{}

	id, _ := account.Create()

	mAccount := &models.Account{Id:id, Status:"disable"}

	id, err = account.Update(mAccount, []string{"Status"})

	assert.Nil(t, err, "更新教室失败", err)

	updateAccount, err := account.Get(id)

	assert.Equal(t, "disable", updateAccount.Status)

	tearDown()
}