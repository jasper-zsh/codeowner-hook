package utils

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	MAX_UNIX_TIMESTAMP = 2147483647
)

func parseNumberTime[T int | float32 | float64](val T) time.Time {
	if val > MAX_UNIX_TIMESTAMP {
		return time.Unix(int64(val/1000), 0)
	} else {
		return time.Unix(int64(val), 0)
	}
}

func decodeHook(from, to reflect.Type, data any) (result any, err error) {
	funcs := []mapstructure.DecodeHookFuncType{
		func(t1, t2 reflect.Type, i interface{}) (interface{}, error) {
			if t2 != reflect.TypeOf(time.Time{}) {
				return data, nil
			}
			switch t1.Kind() {
			case reflect.Float32:
				return parseNumberTime[float32](data.(float32)), nil
			case reflect.Float64:
				return parseNumberTime[float64](data.(float64)), nil
			case reflect.String:
				return time.Parse(time.RFC3339, data.(string))
			default:
				return data, nil
			}
		},
	}
	for _, f := range funcs {
		result, err = f(from, to, data)
		if err != nil {
			return
		}
	}
	return
}

func DecodeMap(src map[string]any, dst interface{}) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: decodeHook,
		Result:     dst,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(src)
}
