package main

import (
	_ "MAYMOE/routers"

	"github.com/astaxie/beego"
)

func init() {

}

func main() {
	beego.SetLogger("file", `{"filename":"logs/aliyun.log"}`)
	beego.Run()
}
