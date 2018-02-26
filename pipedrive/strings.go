package pipedrive

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func Stringify(message interface{}) string {
	var buf bytes.Buffer

	value := reflect.ValueOf(message)

	if value.Kind() == reflect.Ptr && value.IsNil() {
		buf.Write([]byte("<nil>"))
	}

	v := reflect.Indirect(value)

	switch v.Kind() {
	case reflect.Struct:
		data, err := json.Marshal(message)

		if err != nil {
			break
		}

		if v.Type() == reflect.TypeOf(Timestamp{}) {
			break
		}

		buf.Write([]byte(string(data)))
	}

	return buf.String()
}
