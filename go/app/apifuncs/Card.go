package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type recCardPostData struct {
	UID string `json:"uid"`
}

//CardResponse is /card no post ni taisuru func
func CardResponse(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Can't catch uid(io error)")
		return
	}

	var rec recCardPostData

	if err := json.Unmarshal(jsonBytes, &rec); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Can't catch uid(JSON Unmarshal error)")
		return
	}

	if r.Method == http.MethodPost {
		var status string

		//do NANI

		if false {
			status = `{"status":"unabailable"}`
		} else {
			status = `{"status":"available"}`
		}

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintln(w, status)
	}
}
