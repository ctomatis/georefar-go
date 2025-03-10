package internal

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Map map[string]any

func ToMap(v any) (m Map) {
	jbytes, _ := json.Marshal(v)
	json.Unmarshal(jbytes, &m)
	return
}

func toDashCase(s string) string {
	buf := new(strings.Builder)

	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteString("-")
		}
		buf.WriteRune(v)
	}
	return strings.ToLower(buf.String())
}

func Resource(v any) string {
	elm := reflect.TypeOf(v).Elem()

	if !strings.HasPrefix(elm.String(), "resources.") {
		return ""
	}
	return toDashCase(elm.Name())
}
