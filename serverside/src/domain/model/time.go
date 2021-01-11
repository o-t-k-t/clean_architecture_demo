package model

import "time"

var overRappingNowTime *time.Time

// NowTime 現在時刻の取得
func NowTime() time.Time {
	if overRappingNowTime != nil {
		return *overRappingNowTime
	}
	return time.Now()
}

// OverrapNowTime テストスタブ用に現在時刻を上書き更新
func OverrapNowTime(year int, month time.Month, day, hour, min, sec, nsec int) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	t := time.Date(year, month, day, hour, min, sec, nsec, jst)
	overRappingNowTime = &t
}

// AppTime is time class for the application.
type AppTime time.Time

// Unixtime returns unixtime.
func (t AppTime) Unixtime() int64 {
	return time.Time(t).Unix()
}
