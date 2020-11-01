package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

//WeekResponce is a function with regards to /week.
func WeekResponce(w http.ResponseWriter, r *http.Request) {
	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	r.Header.Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {

		var resetSettings []dbctl.ResetSettingData

		resetSettings, err := dbctl.GetResetSettings()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonBytes, err := json.Marshal(resetSettings)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)

		if resetSettings == nil {
			jsonString = "[]"
		}

		fmt.Fprintln(w, jsonString)

	} else if r.Method == http.MethodPost {

		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch Setting Data(io error)", err)
			return
		}

		var rec dbctl.ResetSettingData

		if err := json.Unmarshal(jsonBytes, &rec); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch Setting Data(JSON Unmarshal error)", err)
			return
		}

		if err := dbctl.InsertResetSetting(rec.Email, rec.Day, rec.IsOnce); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(InsertResetSetting)", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"available"}`)

	} else if r.Method == http.MethodPut {

		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch Setting Data(io error)", err)
			return
		}

		var rec dbctl.ResetSettingData

		if err := json.Unmarshal(jsonBytes, &rec); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch Setting Data(JSON Unmarshal error)", err)
			return
		}

		if err := dbctl.DeleteResetSetting(rec.Email, rec.Day, rec.IsOnce); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("database error(DeleteResetSetting)", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"available"}`)
	}
}

//ResetEntryFlagTest debugs ResetEntryFlag.
func ResetEntryFlagTest(w http.ResponseWriter, r *http.Request) {
	dbctl.ResetEntryFlag()
}
