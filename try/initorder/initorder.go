package main

import (
	"github.com/cao7113/hellogolang/try/initorder/subp"
	"log"
	"time"
)

// https://raw.githubusercontent.com/yangwenmai/maiyang.me/master/blog/init.png

var b = time.Now().UnixNano()
var c int64

const a = 123

func init() {
	c = time.Now().UnixNano() // init after a, b, c
}

func main() {
	subp.Fa()
	subp.Fb()
	log.Println(a, b, c)
}
