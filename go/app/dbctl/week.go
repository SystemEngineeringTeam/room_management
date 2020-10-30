package dbctl

import (
	"log"
	"runtime"
	"time"
)

//ResetSettingData is one struct in a reset setting.
type ResetSettingData struct {
	Email  string `json:"email"`
	Day    string `json:"day"`
	IsOnce bool   `json:"isOnce"`
}

//GetResetSettings is a function to get list of resetSettings.
func GetResetSettings() ([]ResetSettingData, error) {
	var resetSettingDatas []ResetSettingData
	rows, err := db.Query("select day, isOnce, email from week, users, emails where week.users_id = users.id and users.email_id = emails.id")
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var day string
		var isOnce int
		var email string
		rows.Scan(&day, &isOnce, &email)

		var isOnceBool bool
		if isOnce == 0 {
			isOnceBool = false
		} else {
			isOnceBool = true
		}

		resetSettingDatas = append(resetSettingDatas, ResetSettingData{Email: email, Day: day, IsOnce: isOnceBool})
	}

	return resetSettingDatas, nil
}

//InsertResetSetting is a function to add a resetSetting to Week.
func InsertResetSetting(email string, day string, isOnce bool) error {
	var isOnceInt int
	if isOnce {
		isOnceInt = 1
	} else {
		isOnceInt = 0
	}
	_, err := db.Exec("insert into week(day, isOnce, users_id) select ?, ?, id from users where exists (select * from emails where emails.id = users.email_id and email = ?)", day, isOnceInt, email)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}

//DeleteResetSetting is a function to delete a resetSetting from Week.
func DeleteResetSetting(email string, day string) error {
	_, err := db.Exec("delete from week where day = ? and users_id in (select users.id from users, emails where users.email_id = emails.id and email = ?)", day, email)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}

//ResetEntryFlag do reset.
func ResetEntryFlag() error {
	wdays := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	day := wdays[(time.Now().Weekday()+6)%7]

	_, err := db.Exec("insert into logs(cards_id, card_read_datetime, isEntry) select min(cards.id), Now(), 0 from users, cards where users.id = cards.user_id and isEntry = 1 and not exists (select * from week where users.id = week.users_id and week.day = ?) group by users.id", day)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	_, err = db.Exec("update users set isEntry = 0 where isEntry = 1 and not exists (select * from week where users.id = week.users_id and week.day = ?)", day)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	if err = deleteOnceResetSetting(day); err != nil {
		return err
	}

	return nil
}

func deleteOnceResetSetting(day string) error {
	_, err := db.Exec("delete from week where isOnce = 1 and day = ?", day)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	return nil
}
