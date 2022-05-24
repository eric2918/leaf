package gate

import (
	"net"
	"github.com/eric2918/leaf/module"
	"github.com/eric2918/leaf/chanrpc"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	Skeleton() *module.Skeleton
	ChanRPC()  *chanrpc.Server
}
