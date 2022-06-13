package gocq

import (
	"time"

	"github.com/Mrs4s/go-cqhttp/global/terminal"
	"github.com/Mrs4s/go-cqhttp/internal/base"
	"github.com/Mrs4s/go-cqhttp/server"
	log "github.com/sirupsen/logrus"
)

// InitBase 解析 flags 与配置文件到 base
func InitBase() {
	base.Parse()
	if !base.FastStart && terminal.RunningByDoubleClick() {
		err := terminal.NoMoreDoubleClick()
		if err != nil {
			log.Errorf("遇到错误: %v", err)
			time.Sleep(time.Second * 5)
		}
		return
	}
	switch {
	case base.LittleH:
		base.Help()
	case base.LittleD:
		server.Daemon()
	case base.LittleWD != "":
		base.ResetWorkingDir()
	}
	base.Init()
}
