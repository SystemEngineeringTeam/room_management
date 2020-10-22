package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

type recUserCardData struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
}

//UserCardResponce is UserCardResponce
func UserCardResponce(w http.ResponseWriter, r *http.Request) {
	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	r.Header.Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch data(io error)", err)
			return
		}

		var rec recUserCardData

		if err := json.Unmarshal(jsonBytes, &rec); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch data(JSON Unmarshal error)", err)
			return
		}

		if err := dbctl.LinkUserAndCard(rec.UID, rec.Email); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(LinkUserAndCard)", err)
			return
		}

		if err := dbctl.OrganizeLog(rec.UID, rec.Email); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(OrganizeLog)", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"available"}`)
	}
}
