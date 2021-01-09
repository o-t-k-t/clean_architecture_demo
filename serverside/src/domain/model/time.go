package model

import "time"

// AppTime is time class for the application.
type AppTime time.Time

// Unixtime returns unixtime.
func (t AppTime) Unixtime() int64 {
	return time.Time(t).Unix()
}
