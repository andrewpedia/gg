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

type Game struct {
	*module.Skeleton
}

func (m *Game) OnInit() {
	fmt.Printf("Game.OnInit:   \n" )

	m.Skeleton = skeleton
}

func (m *Game) OnDestroy() {

}
