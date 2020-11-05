package dbctl

import (
	"errors"
	"log"
	"runtime"
	"time"
)

//ResetSettingData is one struct in a reset setting.
type ResetSettingData struct {
	Email  string `json:"Email"`
	Day    string `json:"Day"`
	IsOnce bool   `json:"IsOnce"`
}

//ResetSettingResponse is used by GetResetSettings.
type ResetSettingResponse struct {
	Day    string `json:"Day"`
	IsOnce bool   `json:"IsOnce"`
}

//SingleEmail is only Email struct.
type SingleEmail struct {
	Email string `json:"Email"`
}

//GetResetSettings is a function to get list of resetSettings.
func GetResetSettings(email string) ([]ResetSettingResponse, error) {
	var resetSettingDatas []ResetSettingResponse
	rows, err := db.Query("select day, isOnce from week, users, emails where week.users_id = users.id and users.email_id = emails.id and email = ?", email)
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
		rows.Scan(&day, &isOnce)

		var isOnceBool bool
		if isOnce == 0 {
			isOnceBool = false
		} else {
			isOnceBool = true
		}

		resetSettingDatas = append(resetSettingDatas, ResetSettingResponse{Day: day, IsOnce: isOnceBool})
	}

	return resetSettingDatas, nil
}

//InsertResetSetting is a function to add a resetSetting to Week.
func InsertResetSetting(email string, day string, isOnce bool) error {
	if err := isDayAndEmail(email, day); err != nil {
		return err
	}

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
func DeleteResetSetting(email string, day string, isOnce bool) error {
	if err := isDayAndEmail(email, day); err != nil {
		return err
	}

	var isOnceInt int
	if isOnce {
		isOnceInt = 1
	} else {
		isOnceInt = 0
	}
	_, err := db.Exec("delete from week where day = ? and isOnce = ? and users_id in (select users.id from users, emails where users.email_id = emails.id and email = ?)", day, isOnceInt, email)
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

func isDayAndEmail(email string, day string) error {
	checkFlag, err := isEmailExist(email)
	if err != nil {
		return err
	}

	if !checkFlag {
		log.Printf("The email does not exist.")
		log.Printf("Please check the email you sent.")
		return errors.New("ERROR: Email does not exist")
	}

	if !isDayFormat(day) {
		log.Printf("Day is illegal.")
		log.Printf("Please check the day you sent.")
		return errors.New("ERROR: Illegal day")
	}

	return nil
}

func isDayFormat(day string) bool {
	wdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	var i int
	for i = 0; i < 7; i++ {
		if day == wdays[i] {
			break
		}
	}
	if i != 7 {
		return true
	}
	return false
}

func isEmailExist(email string) (bool, error) {
	rows, err := db.Query("select count(*) from emails where email = ?", email)
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
