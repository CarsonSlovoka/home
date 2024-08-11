package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var locTaipei *time.Location

type JsonTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (j *JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Time)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`) // 把左右兩邊的分號拿掉
	if strings.Index(s, "/") > 0 {    // 統一換成-號
		s = strings.ReplaceAll(s, "/", "-")
	}
	var (
		theTime time.Time
		err     error
	)

	switch len(s) {
	case 10:
		theTime, err = time.Parse("2006-01-02", s)
		// 轉換成UTC+8
		theTime = theTime.Add(-8 * time.Hour) // UTC-8
		theTime = theTime.In(locTaipei)       // 再換回台北時間，會在+8，這樣的時區就是我們要的

	case 16:
		theTime, err = time.Parse("2006-01-02 15:04", s)
		theTime = theTime.Add(-8 * time.Hour)
		theTime = theTime.In(locTaipei)
	default:
		// 有完整時區(loc)資訊，不用在轉換
		err = json.Unmarshal(b, &theTime)
		if err != nil {
			err = fmt.Errorf("格式必須是: 2006-01-02T15:04:05-07:00 | %w", err)
		}
	}
	if err != nil {
		return err
	}
	j.Time = theTime
	return nil
}
