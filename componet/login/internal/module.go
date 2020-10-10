package internal

import (
	"fmt"
	"github.com/jackiiilong/leafserver/componet/base"
	"github.com/jackiiilong/leafserver/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Login struct {
	*module.Skeleton
}

func (m *Login) OnInit() {
	fmt.Printf("Login.OnInit:   \n" )

	m.Skeleton = skeleton
}

func (m *Login) OnDestroy() {

}
