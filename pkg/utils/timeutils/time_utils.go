package timeutils

import "time"

const (
	strDate = "2006-01-02"
)

func GenerateStrDate() string {
	return time.Now().Format(strDate)
}
