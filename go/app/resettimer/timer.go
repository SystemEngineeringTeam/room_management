package resettimer

import (
	"time"

	"set1.ie.aitech.ac.jp/room_management/dbctl"
)

// タイムゾーンの定義
var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

// ResetTimer はリセットを呼び続ける関数
func ResetTimer() {

	t := time.Now().In(jst)        // 現在時刻を取得
	nextDate := getNextDate(t)     // 次の日を取得
	callResetFunc(nextDate.Sub(t)) // リセットし続けるための関数を呼ぶ

}

func getNextDate(t time.Time) time.Time {
	nextDate := t.Add(24 * time.Hour)              // 24時間後のデータを取得
	y, m, d := nextDate.Date()                     // 年月日を取得
	nextDate = time.Date(y, m, d, 0, 0, 0, 0, jst) // 次の日の00:00:00のデータを生成
	return nextDate
}

func callResetFunc(d time.Duration) {
	timer := time.NewTimer(d)

	for t := range timer.C { // timer.Cはチャンネルから送られた現在時刻
		// フラグをリセットする関数
		dbctl.ResetEntryFlag()
		timer.Reset(getNextDate(t).Sub(t)) // 次の日の00:00:00に発火するように設定
	}
}
