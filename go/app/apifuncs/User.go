package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type recUserPostData struct {
	StudentNumber string `json:"StudentNumber"`
	Email         string `json:"Email"`
	Name          string `json:"Name"`
	UID           string `json:"UID"`
}

type entryPersonInfo struct {
	StudentNumber string `json:"StudentNumber"`
	Name          string `json:"Name"`
	EntryDatetime string `json:"EntryDatetime"`
}

//UserResponse is /user ni taisuru func
func UserResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("Can't catch UserData(io error)")
			return
		}

		var rec recUserPostData

		if err := json.Unmarshal(jsonBytes, &rec); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("Can't catch UserData(JSON Unmarshal error)")
			return
		}

		var status string

		//do HOGE

		if false {
			status = `{"status":"unabailable"}`
		} else {
			status = `{"status":"available"}`
		}

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintln(w, status)

	} else if r.Method == http.MethodGet {
		var students []entryPersonInfo

		// test process
		students = append(students, entryPersonInfo{StudentNumber: "k19092", Name: "bbb", EntryDatetime: "2030/25/19"})
		students = append(students, entryPersonInfo{StudentNumber: "k30001", Name: "abc", EntryDatetime: "2100/12/31"})

		jsonBytes, err := json.Marshal(students)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		if students == nil {
			jsonString = "[]"
		}

		fmt.Fprintln(w, jsonString)
	}
}
