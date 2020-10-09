package dbctl

//LogInfo is Log no Info
type LogInfo struct {
	StudentNumber    string `json:"StudentNumber"`
	Name             string `json:"Name"`
	CardReadDatetime string `json:"CardReadDatetime"`
}

//GetLogInfos can get log no info
func GetLogInfos() []LogInfo {
	rows, err := db.Query("select student_number, name, card_read_datetime, isEntry from logs, cards left outer join users on users.id=cards.user_id where logs.cards_id=cards.id order by card_read_datetime")
	return nil
}
