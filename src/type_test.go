// go test carson.io -v

package main

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_marshalJSON(t *testing.T) {
	tUTC0, _ := time.Parse("2006-01-02", "2024-08-10")
	jT := JsonTime{tUTC0}
	bs, err := json.Marshal(jT)
	if err != nil {
		t.Fatal(err)
	}
	if s := string(bs); s != `"2024-08-10T00:00:00Z"` {
		t.Fatal(s)
	}
}

func Test_unmarshalJSON(t *testing.T) {
	var jT JsonTime
	for i, d := range []struct {
		data                 string
		layout               string
		expectedFormatString any
	}{
		{`"2024-08-10"`, "2006-01-02", "2024-08-10"},
		{`"2024-08-10 20:48"`, "2006-01-02 15:04", "2024-08-10 20:48"},
		{`"2024-08-10T20:48:00+08:00"`, "2006-01-02 15:04:05 -07:00", "2024-08-10 20:48:00 +08:00"},
	} {
		err := json.Unmarshal([]byte(d.data), &jT)
		if err != nil {
			t.Fatal(err)
		}
		if jT.Format(d.layout) != d.expectedFormatString {
			t.Fatal(i)
		}
	}
}
