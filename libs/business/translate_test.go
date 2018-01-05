package business

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseUniqueUserId(t *testing.T) {
	var (
		accountId int64 = 2
		clientId int64 = 1001
	)
	uniClientId := UniqueUserId(accountId, clientId)

	parseAccountId, parseClientId := ParseUniqueUserId(uniClientId)

	assert.Equal(t, accountId, parseAccountId)
	assert.Equal(t, clientId, parseClientId)
}
