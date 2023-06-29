package response

import "time"

type Template struct {
	Status  int
	Data    interface{}
	Message string
	Date    time.Time
}
