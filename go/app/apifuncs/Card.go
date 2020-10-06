package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type resCardPost struct {
	Status string `json:"status"`
}

//CardResponce is /card no post ni taisuru func
func CardResponce(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	var uid string

	if len(q["uid"]) > 0 {
		uid = q["uid"][0]
	} else {
		fmt.Println("Can't catch uid")
	}

	if r.Method == http.MethodPost {
		status := uid //ありえない代入ですが、uidをどこかで使わないといけないので…
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
