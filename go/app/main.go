package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"set1.ie.aitech.ac.jp/room_management/apifuncs"
)

func main() {
	log.Print("hello world\n")
	http.HandleFunc("/card", apifuncs.CardResponce)
	http.ListenAndServe(":80", nil)
}
