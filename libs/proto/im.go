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