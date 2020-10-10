package conf

import (
	"encoding/json"
	"fmt"
	"github.com/jackiiilong/leafserver/log"
	"io/ioutil"
	"os"
	"runtime/debug"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("...%v", err)
	}
	fmt.Println("current working dir: " + path)

	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		log.Fatal("...%v", err)
	}

	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
