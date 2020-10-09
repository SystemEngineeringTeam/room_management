package dbctl

import (
	"database/sql"
	"log"
	"runtime"
)

//LogInfo is Log no Info
type LogInfo struct {
	StudentNumber    string `json:"StudentNumber"`
	Name             string `json:"Name"`
	CardReadDatetime string `json:"CardReadDatetime"`
	EntryOrExit      int64  `json:"EntryOrExit"`
}

//GetLogInfos can get log no info
func GetLogInfos() ([]LogInfo, error) {
	rows, err := db.Query("select logs.id, student_number, name, card_read_datetime, logs.isEntry from logs, cards left outer join users on users.id=cards.user_id where logs.cards_id=cards.id order by card_read_datetime")
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return nil, err
	}
	defer rows.Close()

	var LogInfos []LogInfo

	for rows.Next() {
		var logsID int
		var StudentNumber sql.NullString = sql.NullString{}
		var Name sql.NullString = sql.NullString{}
		var CardReadDateTime sql.NullString = sql.NullString{}
		var EntryOrExit sql.NullInt64 = sql.NullInt64{}
		rows.Scan(&logsID, &StudentNumber, &Name, &CardReadDateTime, &EntryOrExit)
		if !CardReadDateTime.Valid {
			log.Printf("LogData read error(CardReadDateTime IS NULL)")
			log.Printf("Logs ID is %d", logsID)
			continue
		}
		if !EntryOrExit.Valid {
			log.Printf("LogData read error(isEntry IS NULL)")
			log.Printf("Logs ID is %d", logsID)
			continue
		}
		LogInfos = append(LogInfos, LogInfo{StudentNumber: StudentNumber.String, Name: Name.String, CardReadDatetime: CardReadDateTime.String, EntryOrExit: EntryOrExit.Int64})
	}

	return LogInfos, nil
}
