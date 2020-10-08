package dbctl

import (
	"database/sql"
	"log"
	"runtime"
)

//IsCardRegistered is uidからカードが登録済かどうか判定する func
func IsCardRegistered(uid string) (bool, error) {
	rows, err := db.Query("select count(*) from cards where uid=?", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return false, err
	}
	defer rows.Close()

	var count int
	rows.Next()
	rows.Scan(&count)

	if count == 1 {
		return true, nil
	}
	return false, nil
}

//ToggleEntryORExit is 入退室を切り替える func
func ToggleEntryORExit(uid string) error {
	rows, err := db.Query("select isEntry from users where exists(select * from cards where users.id=cards.user_id and uid=?)", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	var isEntry sql.NullInt64 = sql.NullInt64{}
	rows.Next()
	rows.Scan(&isEntry)

	var nextValue int64
	if isEntry.Valid {
		nextValue = 1 - isEntry.Int64
	} else {
		log.Printf("isEntry is NULL")
		return nil
	}

	_, err = db.Exec("update users set isEntry=? where exists(select * from cards where users.id=cards.user_id and uid=?)", nextValue, uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}

//InsertCard is Card ni insert suru func.
func InsertCard(uid string) error {
	_, err := db.Exec("insert into cards(uid, user_id) values(?, NULL)", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}

//InsertLog is Log ni insert suru func.
func InsertLog(uid string) error {
	rows, err := db.Query("select id from cards where uid=?", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	var cardID int
	rows.Next()
	rows.Scan(&cardID)

	_, err = db.Exec("insert into logs(cards_id, card_read_datetime) values(?, Now())", cardID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	return nil
}
