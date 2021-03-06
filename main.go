package main

import (
	_ "github.com/kykl/goapi/docs"
	_ "github.com/kykl/goapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
