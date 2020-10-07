package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type uidJSON struct {
	UID string `json:"uid"`
}

type resCardPost struct {
	Status string `json:"status"`
}

//CardResponse is /card no post ni taisuru func
func CardResponse(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Can't catch uid(io error)")
		return
	}

	var rec uidJSON

	if err := json.Unmarshal(jsonBytes, &rec); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Can't catch uid(JSON Unmarshal error)")
		return
	}

	if r.Method == http.MethodPost {
		status := rec.UID //ありえない代入ですが、uidをどこかで使わないといけないので…
		//do NANI
		if false {
			status = "unabailable"
		} else {
			status = "available"
		}

		var res resCardPost
		res.Status = status

		jsonBytes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintln(w, jsonString)
	}
}
