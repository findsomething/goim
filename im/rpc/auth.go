package rpc

type Auther interface {
	Auth(token string) (userId int64, roomId int32, err error)
}

type DefaultAuther struct {

}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}
