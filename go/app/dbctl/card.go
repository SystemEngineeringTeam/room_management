package dbctl

import (
	"database/sql"
	"errors"
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
	} else if count > 1 {
		log.Printf("ERROR: the same UID")
		log.Printf("The same UID card exists.")
		log.Printf("Please check the cards table.")
		return true, errors.New("ERROR: the same UID")
	}
	return false, nil
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
	rows, err := db.Query("select id, user_id from cards where uid=?", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	var cardID int
	var userID sql.NullString = sql.NullString{}
	rows.Next()
	rows.Scan(&cardID, &userID)

	if userID.Valid {
		rows, err = db.Query("select isEntry from logs, cards where logs.cards_id = cards.id and user_id in (select user_id from cards where uid = ?) order by card_read_datetime desc limit 1", uid)
	} else {
		rows, err = db.Query("select isEntry from logs where cards_id = ? order by card_read_datetime desc limit 1", cardID)
	}
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	var isEntry sql.NullInt64 = sql.NullInt64{}
	rows.Next()
	rows.Scan(&isEntry)

	var nextValue int64
	if isEntry.Valid {
		nextValue = 1 - isEntry.Int64
	} else {
		nextValue = 1
	}

	_, err = db.Exec("insert into logs(cards_id, card_read_datetime, isEntry) values(?, Now(), ?)", cardID, nextValue)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	toggleEntryORExit(nextValue, cardID)

	return nil
}

func toggleEntryORExit(nextValue int64, cardID int) error {
	_, err := db.Exec("update users set isEntry=? where exists(select * from cards where users.id=cards.user_id and cards.id=?)", nextValue, cardID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}
