package internal

import (
	"fmt"
	"github.com/jackiiilong/leafserver/componet/game"
	"github.com/jackiiilong/leafserver/componet/msg"
	conf "github.com/jackiiilong/leafserver/conf"
	"github.com/jackiiilong/leafserver/gate"
)

type Gate struct {
	*gate.Gate
}

func (m *Gate) OnInit() {
	fmt.Printf("Gate.OnInit:   \n" )

	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
