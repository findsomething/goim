package rpc

import (
	"strings"
	"fmt"
	"goim/libs/define"
	"strconv"
	"time"
	"goim/libs/crypto/md5"
	"goim/libs/business"
)

const AUTH_EXPIRE = 86400

type Auther interface {
	Auth(token string) (userId int64, roomId int32, err error)
	CreateToken(accountId, clientId int64) (token string)
}

type TokenEntity struct {
	expire    int64
	sign      string
	accountId int64
	clientId  int64
	server    string
}

type DefaultAuther struct {
	secretKey string
}

func NewDefaultAuther(secretKey string) *DefaultAuther {
	return &DefaultAuther{secretKey: secretKey}
}

// sign:accountId:clientId:expire:server
func (a *DefaultAuther) Auth(token string) (userId int64, roomId int64, err error) {
	var (
		parseToken *TokenEntity
	)
	if parseToken, err = a.checkToken(token); err != nil {
		return
	}
	userId = business.UniqueUserId(parseToken.accountId, parseToken.clientId)
	roomId = define.NoRoom
	return
}

func (a *DefaultAuther) CreateToken(accountId, clientId int64) (token string) {
	expire := time.Now().Unix() + AUTH_EXPIRE
	sign := fmt.Sprintf("%s:%d:%d:%d:%s", a.secretKey, accountId, clientId, expire,
		define.IM_SERVER)
	token = fmt.Sprintf("%s:%d:%d:%d:%s", md5.Encode(sign), accountId, clientId, expire, define.IM_SERVER)
	return
}

func (a *DefaultAuther) checkToken(token string) (tokenEntity *TokenEntity, err error) {
	var (
		expire    int64
		sign      string
		accountId int64
		clientId  int64
		server    string
	)
	auths := strings.Split(token, ":")
	if len(auths) != 5 {
		err = fmt.Errorf(define.TOKEN_FORMAT_ERROR)
		return
	}
	if expire, err = strconv.ParseInt(auths[3], 10, 64); err != nil || time.Now().Unix() > expire {
		err = fmt.Errorf(define.TOKEN_EXPIRE_ERROR)
		return
	}
	sign = auths[0]
	accountId, _ = strconv.ParseInt(auths[1], 10, 64)
	clientId, _ = strconv.ParseInt(auths[2], 10, 64)
	if server = auths[4]; server != define.IM_SERVER {
		err = fmt.Errorf(define.IM_SERVER_ERROR)
		return
	}
	if !md5.Check(fmt.Sprintf("%s:%d:%d:%d:%s", a.secretKey, accountId, clientId, expire, server), sign) {
		err = fmt.Errorf(define.TOKEN_EXPIRE_ERROR)
		return
	}
	tokenEntity = &TokenEntity{
		expire:    expire,
		sign:      sign,
		accountId: accountId,
		clientId:  clientId,
		server:    server,
	}
	return
}
