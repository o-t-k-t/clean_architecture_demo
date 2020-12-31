package model

import "time"

// AppTime is time class for the application.
type AppTime time.Time

// CurrentAppTime creates AppTime of now.
func CurrentAppTime() AppTime {
	return AppTime(time.Now())
}

// Unixtime returns unixtime.
func (t AppTime) Unixtime() int64 {
	return t.Unixtime()
}
