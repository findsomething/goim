package proto

type AccountArg struct {
	Id     int64
	Status string
	Cols   []string
}

type AccountReply struct {
	Id        int64
	Status    string
	AccessKey string
	SecretKey string
}

type ConversationArg struct {
	Id        int64
	AccountId int64
	Name      string
	Type      string
	Status    string
	Attr      string
	Cols      []string
}

type ConversationReply struct {
	Id        int64
	AccountId int64
	No        string
	Name      string
	Type      string
	Status    string
	Attr      string
}

type ConversationMemberArg struct {
	Id         int64
	AccountId  int64
	ConvNo     string
	ClientId   int64
	ClientName string
	Mute       int32
	Forbidden  int32
	Cols       []string
}

type ConversationMemberReply struct {
	Id         int64
	AccountId  int64
	ConvNo     string
	ClientId   int64
	ClientName string
	Mute       int32
	Forbidden  int32
}

type TokenArg struct {
	AccountId int64
	ClientId  int64
}

type TokenReply struct {
	Token string
}

type VerifyTokenArg struct {
	Token string
}

type VerifyTokenReply struct {
	UserId int64
	RoomId int64
	err    error
}
