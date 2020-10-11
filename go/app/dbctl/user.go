package dbctl

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"runtime"
)

// RecUserPostData はユーザの情報を受け取るための構造体
type RecUserPostData struct {
	StudentNumber string `json:"StudentNumber"`
	Email         string `json:"Email"`
	Name          string `json:"Name"`
	Password      string `json:"Password"`
	UID           string `json:"UID"`
}

// EntryPersonInfo は今いる人の一覧を返すための構造体
type EntryPersonInfo struct {
	StudentNumber string `json:"StudentNumber"`
	Name          string `json:"Name"`
	EntryDatetime string `json:"EntryDatetime"`
}

func passwordToSha256(s string) string {
	b := sha256.Sum256([]byte(s))
	pass := hex.EncodeToString(b[:])
	return pass
}

// Register はユーザーを登録する関数
func Register(u RecUserPostData) error {
	hashedPassword := passwordToSha256(u.Password)
	_, err := db.Exec("insert into emails(email,password) values(?,?);", u.Email, hashedPassword)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	_, err = db.Exec("insert into users(student_number, name, isEntry, email_id) select ?, ?, logs.isEntry, emails.id from emails, cards, logs where cards.id=logs.cards_id and email=? and uid=? order by card_read_datetime desc limit 1;", u.StudentNumber, u.Name, u.Email, u.UID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	_, err = db.Exec("update cards set user_id=(select id from users where student_number=?) where uid=?;", u.StudentNumber, u.UID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}

// GetCurrentEntryMembers は今いる人の一覧を返す関数
func GetCurrentEntryMembers() ([]EntryPersonInfo, error) {
	rows, err := db.Query("select student_number, name, max(card_read_datetime) from users, cards, logs where users.id = cards.user_id and logs.cards_id = cards.id and users.isEntry = 1 group by student_number, name")
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer rows.Close()

	users := make([]EntryPersonInfo, 0)

	for rows.Next() {
		tmp := EntryPersonInfo{}
		rows.Scan(&tmp.StudentNumber, &tmp.Name, &tmp.EntryDatetime)
		if err != nil {
			pc, file, line, _ := runtime.Caller(0)
			f := runtime.FuncForPC(pc)
			log.Printf(errFormat, err, f.Name(), file, line)
			return nil, err
		}

		users = append(users, tmp)
	}

	return users, nil
}
