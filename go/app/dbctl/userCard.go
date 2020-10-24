package dbctl

import (
	"database/sql"
	"errors"
	"log"
	"runtime"
)

//LinkUserAndCard はユーザとカードをリンクさせる関数
func LinkUserAndCard(uid string, email string) error {
	//カードに既にユーザが登録されていないか確認する
	if err := checkLinkedUserAndCard(uid, email); err != nil {
		return err
	}

	//update
	res, err := db.Exec("update cards set user_id = (select users.id from users, emails where users.email_id = emails.id and email = ?) where uid = ?", email, uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	//updateによって影響を受けた行の数を取得
	var affectedRows int64
	if affectedRows, err = res.RowsAffected(); err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	//更新時にエラーが発生しなかったが、紐付けできなかった場合。
	if affectedRows == 0 {
		log.Printf("ERROR: Linkage failure")
		log.Printf("It's possible that the email or uid did not exist.")
		log.Printf("Please check the uid and email you entered.")
		return errors.New("ERROR: Linkage failure")
	}

	return nil
}

//OrganizeLog はLogの矛盾を解消する関数
func OrganizeLog(uid string, email string) error {
	//LinkUserAndCardで紐付けしたカードの最も古い読み取り時間を取得
	rows, err := db.Query("select card_read_datetime, user_id from logs, cards where logs.cards_id = cards.id and uid=? order by card_read_datetime asc limit 1", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	var oldCardReadDatetime string
	var userID string
	rows.Next()
	rows.Scan(&oldCardReadDatetime, &userID)

	//ユーザの紐付けしたカードを初めて使う直前の入退室情報を取得
	rows, err = db.Query("select isEntry from logs, cards where logs.cards_id = cards.id and card_read_datetime < ? and user_id = ? order by card_read_datetime desc", oldCardReadDatetime, userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	var isEntry sql.NullInt64 = sql.NullInt64{}
	var nextValue int64
	rows.Next()
	rows.Scan(&isEntry)
	if !isEntry.Valid {
		nextValue = 1
	} else {
		nextValue = 1 - isEntry.Int64
	}

	_, err = db.Exec("update logs inner join (select id, row_number() over (order by id asc) RowNum from (select logs.id from logs inner join cards on logs.cards_id = cards.id where card_read_datetime >= ? and user_id = ? order by card_read_datetime asc) l) t on logs.id = t.id set isEntry = case when RowNum%2=1 then ? else ? end", oldCardReadDatetime, userID, nextValue, 1-nextValue)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	if err := syncIsEntry(userID); err != nil {
		return err
	}

	return nil
}

//checkLinkedUserAndCard is カードに既にユーザが登録されているか確認する関数
//登録されていた場合、エラーになる。
func checkLinkedUserAndCard(uid string, email string) error {
	rows, err := db.Query("select user_id from cards where uid = ?", uid)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	defer rows.Close()

	var userID sql.NullString = sql.NullString{}
	rows.Next()
	rows.Scan(&userID)
	if userID.Valid {
		log.Printf("ERROR: Already registered user")
		log.Printf("A user is already registered on the card.")
		log.Printf("Please check uid.")
		return errors.New("ERROR: Already registered user")
	}

	return nil
}

//syncIsEntry is usersのisEntryを更新する関数
func syncIsEntry(userID string) error {
	_, err := db.Exec("update users set isEntry = (select isEntry from logs where exists (select * from cards where logs.cards_id = cards.id and users.id = ?) order by card_read_datetime desc limit 1) where id = ?", userID, userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}
