package main

import (
	"github.com/Mrs4s/go-cqhttp/cmd/gocq"

	_ "github.com/Mrs4s/go-cqhttp/db/leveldb"   // leveldb
	_ "github.com/Mrs4s/go-cqhttp/modules/mime" // mime检查模块
	_ "github.com/Mrs4s/go-cqhttp/modules/silk" // silk编码模块
	// 其他模块
	// _ "github.com/Mrs4s/go-cqhttp/db/mongodb"    // mongodb 数据库支持
	// _ "github.com/Mrs4s/go-cqhttp/modules/pprof" // pprof 性能分析
)

func main() {
	gocq.InitBase()
	gocq.InitLog()
	gocq.InitCache()
	gocq.InitDB()
	gocq.PrintBanner()
	gocq.LoadDevice()
	gocq.CheckKey(gocq.ParseCommand())
	gocq.Main()
}
