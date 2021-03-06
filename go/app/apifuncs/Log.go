package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

//LogResponse is /log ni taisuru func
func LogResponse(w http.ResponseWriter, r *http.Request) {

	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodGet {

		var logInfos []dbctl.LogInfo

		logInfos, err := dbctl.GetLogInfos()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonBytes, err := json.Marshal(logInfos)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		if logInfos == nil {
			jsonString = "[]"
		}

		fmt.Fprintln(w, jsonString)
	}
}
