package business

import (
	"strconv"
	"fmt"
)

const UNI_HEADER_LEN = 8
const UNI_HEADER = 10000000

func UniqueUserId(accountId, clientId int64) (userId int64) {
	uniClientIdStr := fmt.Sprintf("%d%d", UNI_HEADER+accountId, clientId)
	userId, _ = strconv.ParseInt(uniClientIdStr, 10, 64)
	return
}

func ParseUniqueUserId(uniClientId int64) (accountId, clientId int64) {
	uniClientIdStr := strconv.FormatInt(uniClientId, 10)
	accountIdStr := uniClientIdStr[:UNI_HEADER_LEN]
	clientIdStr := uniClientIdStr[UNI_HEADER_LEN:]
	accountId, _ = strconv.ParseInt(accountIdStr, 10, 64)
	accountId -= UNI_HEADER
	clientId, _ = strconv.ParseInt(clientIdStr, 10, 64)
	return
}
