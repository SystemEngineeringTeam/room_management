package main

import (
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/resettimer"

	_ "github.com/go-sql-driver/mysql"
	"set1.ie.aitech.ac.jp/room_management/apifuncs"
)

func main() {
	log.Print("hello world\n")
	http.HandleFunc("/card", apifuncs.CardResponse)
	http.HandleFunc("/user", apifuncs.UserResponse)
	http.HandleFunc("/user/card", apifuncs.UserCardResponce)
	http.HandleFunc("/log", apifuncs.LogResponse)
	http.HandleFunc("/reset", apifuncs.WeekResponce)
	// http.HandleFunc("/resetTest", apifuncs.ResetEntryFlagTest)

	resettimer.ResetTimer()

	http.ListenAndServe(":80", nil)
}
