package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type logInfo struct {
	StudentNumber    string `json:"StudentNumber"`
	Name             string `json:"Name"`
	CardReadDatetime string `json:"CardReadDatetime"`
}

//LogResponse is /log ni taisuru func
func LogResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		var logInfos []logInfo

		// test process
		logInfos = append(logInfos, logInfo{StudentNumber: "k19092", Name: "bbb", CardReadDatetime: "2030/25/19"})
		logInfos = append(logInfos, logInfo{StudentNumber: "k30001", Name: "abc", CardReadDatetime: "2100/12/31"})

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
