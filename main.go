package main

import (
	"woocoin/cli"
	"woocoin/db"
)

func main() {
	defer db.Close() //defer main 함수종료시 실행됨
	cli.Start()
}
