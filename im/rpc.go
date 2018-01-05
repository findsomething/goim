package main

import (
	log "github.com/thinkboy/log4go"
	inet "goim/libs/net"
	"net"
	"net/rpc"
	"goim/im/models"
	"goim/libs/proto"
	"goim/im/tool"
	imrpc "goim/im/rpc"
)

func InitRpc(rpcConf tool.RpcConfig) (err error) {
	var (
		network, addr string
		c             = &RPC{secretKey: rpcConf.SecretKey}
	)
	models.InitMysql(rpcConf)
	rpc.Register(c)
	for i := 0; i < len(Conf.RPCAddrs); i++ {
		log.Info("start listen rpc addr: \"%s\"", Conf.RPCAddrs[i])
		if network, addr, err = inet.ParseNetwork(Conf.RPCAddrs[i]); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
			return
		}
		go rpcListen(network, addr)
	}
	return
}

func rpcListen(network, addr string) {
	l, err := net.Listen(network, addr)
	if err != nil {
		log.Error("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	// if process exit, then close the rpc bind
	defer func() {
		log.Info("rpc addr: \"%s\" close", addr)
		if err := l.Close(); err != nil {
			log.Error("listener.Close() error(%v)", err)
		}
	}()
	rpc.Accept(l)
}

type RPC struct {
	secretKey string
}

func (r *RPC) Ping(arg *proto.NoArg, reply *proto.NoReply) error {
	return nil
}

func (r *RPC) CreateAccount(arg *proto.NoArg, reply *proto.AccountReply) (err error) {
	defaultAccounter := imrpc.NewDefaultAccounter()
	err = defaultAccounter.Create(arg, reply)

	return
}

func (r *RPC) UpdateAccount(arg *proto.AccountArg, reply *proto.AccountReply) (err error) {
	defaultAccounter := imrpc.NewDefaultAccounter()
	err = defaultAccounter.Update(arg, reply)

	return
}

func (r *RPC) CreateConversation(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error) {
	defaultConversationer := imrpc.NewDefaultConversationer()
	err = defaultConversationer.Create(arg, reply)

	return
}

func (r *RPC) UpdateConversation(arg *proto.ConversationArg, reply *proto.ConversationReply) (err error) {
	defaultConversationer := imrpc.NewDefaultConversationer()
	err = defaultConversationer.Update(arg, reply)

	return
}

func (r *RPC) CreateConversationMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error) {
	defaultConversationer := imrpc.NewDefaultConversationer()
	err = defaultConversationer.CreateMember(arg, reply)

	return
}

func (r *RPC) UpdateConversationMember(arg *proto.ConversationMemberArg, reply *proto.ConversationMemberReply) (err error) {
	defaultConversationer := imrpc.NewDefaultConversationer()
	err = defaultConversationer.UpdateMember(arg, reply)

	return
}

func (r *RPC) Token(arg *proto.TokenArg, reply proto.TokenReply) (err error) {
	defaultDefaultAuther := imrpc.NewDefaultAuther(r.secretKey)
	reply.Token = defaultDefaultAuther.CreateToken(arg.AccountId, arg.ClientId)

	return
}

func (r *RPC) Verify(arg *proto.VerifyTokenArg, reply *proto.VerifyTokenReply) (err error) {
	defaultDefaultAuther := imrpc.NewDefaultAuther(r.secretKey)
	reply.UserId, reply.RoomId, err = defaultDefaultAuther.Auth(arg.Token)

	return
}
