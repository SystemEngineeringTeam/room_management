package dbctl

import (
	"database/sql"
	"encoding/json"
	"log"
	"runtime"
)

// UserData is user's login information.
type userData struct {
	Name          sql.NullString `json:"Name"`
	StudentNumber sql.NullString `json:"StudentNumber"`
	Email         sql.NullString `json:"Email"`
}

// Login is a func to login as a user.
func Login(email, password string) (string, error) {
	rows, err := db.Query("select name,student_number,email from users,emails where users.email_id=emails.id and email=? and password=?", email, passwordToSha256(password))
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

		return "", err
	}
	user := userData{}
	for rows.Next() {
		rows.Scan(&user.Name, &user.StudentNumber, &user.Email)
	}

	jsonBytes, err := json.Marshal(struct {
		Name          string `json:"Name"`
		StudentNumber string `json:"StudentNumber"`
		Email         string `json:"Email"`
	}{
		Name: changeNullStringToString(user.Name), StudentNumber: changeNullStringToString(user.StudentNumber), Email: changeNullStringToString(user.Email),
	})
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

		return "", err
	}

	json := string(jsonBytes)

	return json, nil
}
