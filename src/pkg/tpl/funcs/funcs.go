package funcs

import (
	"github.com/CarsonSlovoka/go-pkg/v2/tpl/funcs"
	"html/template"
	"log"
	"reflect"
	"strings"
	"time"
)

func GetUtilsFuncMap() map[string]any {
	funcMap := funcs.GetUtilsFuncMap()
	funcMap["safeHTML"] = func(val string) template.HTML { // 承諾此數值是安全的，不需要額外的跳脫字元來輔助
		return template.HTML(val)
	}
	funcMap["debug"] = func(a ...any) string {
		log.Printf("%+v", a)
		return "" // fmt.Sprintf("%+v", a) // 只把訊息顯示在console，避免放到html之中
	}
	funcMap["timeStr"] = func(t time.Time) string {
		// t.Format("2006-01-02 15:04") // 到分感覺沒有意義
		return t.Format("2006-01-02")
	}

	funcMap["hasSuffix"] = func(s, suffix string) bool {
		return strings.HasSuffix(s, suffix)
	}

	funcMap["time"] = func(value string) (time.Time, error) {
		return time.Parse("2006-01-02", value)
	}

	funcMap["set"] = func(obj any, key string, val any) (string, error) {
		ps := reflect.ValueOf(obj)
		s := ps.Elem()
		if s.Kind() != reflect.Struct {
			log.Printf("type error. 'Struct' expected\n")
			return "", nil
		}
		field := s.FieldByName(key)
		if !field.IsValid() {
			log.Printf("key not found: %s\n", key)
			return "", nil
		}

		if !field.CanSet() {
			log.Printf("The field[%s] is unchangeable. You can't change it.\n", key)
			return "", nil
		}
		field.Set(reflect.ValueOf(val))
		return "", nil
	}
	return funcMap
}
