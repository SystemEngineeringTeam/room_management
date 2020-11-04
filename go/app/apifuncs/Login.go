package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

// Login is
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		userData := struct {
			Email    string `json:"Email"`
			Password string `json:"Password"`
		}{}

		err = json.Unmarshal(jsonBytes, &userData)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("json unmarshal error")
			return
		}

		ret, err := dbctl.Login(userData.Email, userData.Password)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("database error(Login)")
			return
		}

		w.WriteHeader(http.StatusOK)

		r.Header.Set("Content-Type", "application/json")
		fmt.Println(ret)

		fmt.Fprintln(w, ret)

	}
}
