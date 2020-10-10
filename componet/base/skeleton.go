package base

import (
	"github.com/jackiiilong/leafserver/chanrpc"
	"github.com/jackiiilong/leafserver/conf"
	"github.com/jackiiilong/leafserver/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
