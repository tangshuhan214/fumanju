package main

import (
	_ "fumanju/boot"
	_ "fumanju/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
