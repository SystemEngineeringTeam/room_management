package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

type recCardPostData struct {
	UID string `json:"uid"`
}

//CardResponse is /card no post ni taisuru func
func CardResponse(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, `{"status":"Unavailable"}`)
		fmt.Println("Can't catch uid(io error)", err)
		return
	}

	var rec recCardPostData

	if err := json.Unmarshal(jsonBytes, &rec); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, `{"status":"Unavailable"}`)
		fmt.Println("Can't catch uid(JSON Unmarshal error)", err)
		return
	}

	if r.Method == http.MethodPost {
		//card insert or user update
		isCardRegistered, err := dbctl.IsCardRegistered(rec.UID)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(IsCardRegistered)", err)
			return
		}

		if isCardRegistered {
			err = dbctl.ToggleEntryORExit(rec.UID)
		} else {
			err = dbctl.InsertCard(rec.UID)
		}
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(ToggleEntryORExit or InsertCard)", err)
			return
		}
		//log insert
		err = dbctl.InsertLog(rec.UID)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(InsertLog)", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"available"}`)
	}
}
