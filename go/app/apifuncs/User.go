package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

//UserResponse is /user ni taisuru func
func UserResponse(w http.ResponseWriter, r *http.Request) {
	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodPost {
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("Can't catch UserData(io error)")
			return
		}

		var rec dbctl.RecUserPostData

		if err := json.Unmarshal(jsonBytes, &rec); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("Can't catch UserData(JSON Unmarshal error)")
			return
		}

		var status string

		checkResult, err := dbctl.CheckEmailAndStudentNumberRegistered(rec.StudentNumber, rec.Email)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("database error(CheckEmailAndStudentNumberRegistered)")
			return
		}

		if checkResult {
			status = `{"status":"unabailable"}`
			fmt.Println("ERROR: The user already exists.(Duplicate student number or Email)")
		} else {
			if err := dbctl.Register(rec); err != nil {
				status = `{"status":"unabailable"}`
			} else {
				status = `{"status":"available"}`
			}
		}

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintln(w, status)

	} else if r.Method == http.MethodGet {
		var students []dbctl.EntryPersonInfo
		students, err := dbctl.GetCurrentEntryMembers()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

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
