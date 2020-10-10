package module

import (
	"fmt"
	"github.com/jackiiilong/leafserver/conf"
	"github.com/jackiiilong/leafserver/log"
	"runtime"
	"sync"
)

type IModule interface {
	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}

type module struct {
	mi       IModule
	closeSig chan bool
	wg       sync.WaitGroup
}

var mods []*module

func Register(mi IModule) {
	fmt.Printf("Register IModule:  " )
	fmt.Println(mi)
	fmt.Println(mi.Run)

	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	mods = append(mods, m)
}

func Init() {
	for i := 0; i < len(mods); i++ {
		mods[i].mi.OnInit()
	}

	for i := 0; i < len(mods); i++ {
		m := mods[i]
		m.wg.Add(1)
		go run(m)
	}
}

func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}

func run(m *module) {
	fmt.Println("run go routine:   \n" )
	fmt.Println(m.mi.Run )

	m.mi.Run(m.closeSig)
	m.wg.Done()
}

func destroy(m *module) {
	defer func() {
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				log.Error("%v: %s", r, buf[:l])
			} else {
				log.Error("%v", r)
			}
		}
	}()

	m.mi.OnDestroy()
}
