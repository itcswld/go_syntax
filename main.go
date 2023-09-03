package main

import stmt "syntax/04_Statements"

//Variables "syntax/01_Variables"

func main() {
	stmt.GotoStmt()
}

//go run main.go

//生成執行檔
//go build
//--執行 執行檔
//./build

//生成命名執行檔
//go build exefile.go
//./exefile

//GOOS=Linux GOARCH=amd64 go build stmt.go
//CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build

//go install

//go get -d --just download dont install
//go get -fix	-- fix before install
//go get -insecure --download pkg which link not support https
//go get -u --update installed pkg
