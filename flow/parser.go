package flow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/webitel/flow_manager/model"
)

/*
\d := map[string]interface{}{
		"terminator": "ddsada",
		"files": []interface{}{
			map[string]interface{}{
				"name": 123,
				"id":   "${123}",
			},
		},
		"getDigits": map[string]interface{}{
			"setVar":    "getIvrDigit",
			"min":       "3",
			"max":       4,
			"tries":     1,
			"timeout":   2000,
			"flushDTMF": true,
		},
	}
*/

func (f *Flow) Decode(in interface{}, out interface{}) *model.AppError {
	var hook mapstructure.DecodeHookFuncType = func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		kind := from.Kind()
		if kind == reflect.String {
			switch to.Kind() {
			case reflect.Slice:
				var res interface{}
				body, err := json.Marshal(data)
				if err != nil {
					return nil, err
				}

				if len(body) < 2 {
					return data, nil
				}

				txt := []byte(f.Connection.ParseText(string(body[1 : len(body)-1])))

				err = json.Unmarshal(txt, &res)
				if err != nil {
					return []interface{}{}, nil
				}
				return res, nil
			case reflect.String:
				return f.Connection.ParseText(data.(string)), nil
			case reflect.Interface:
				return f.Connection.ParseText(data.(string)), nil
			//fixme added more types
			case reflect.Int, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Int64, reflect.Int32:
				v := f.Connection.ParseText(data.(string))
				if v == "" {
					return 0, nil
				}

				if strings.Index(v, ".") > -1 {
					res, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return 0, err
					}
					return res, nil
				}
				return v, nil
			case reflect.Bool:
				return f.Connection.ParseText(data.(string)), nil
			}
		}
		return data, nil
	}

	return f.decode(in, out, hook)
}

func (f *Flow) DecodeSrc(in interface{}, out interface{}) *model.AppError {
	var hook mapstructure.DecodeHookFuncType = func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if from.Kind() == reflect.String {
			switch to.Kind() {
			case reflect.String:
				return data.(string), nil
			case reflect.Interface:
				return data.(string), nil
			case reflect.Int:
				v := data.(string)
				if v == "" {
					return 0, nil
				}
				return v, nil
			case reflect.Bool:
				return data.(string), nil
			}
		}
		return data, nil
	}

	return f.decode(in, out, hook)
}

func (f *Flow) decode(in interface{}, out interface{}, hook mapstructure.DecodeHookFuncType) *model.AppError {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		DecodeHook:       hook,
		Result:           &out,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return model.NewAppError("Parser", "parser.decode.create.err", nil, err.Error(), http.StatusBadRequest)
	}
	err = decoder.Decode(in)
	if err != nil {
		return model.NewAppError("Parser", "parser.decode.parse.err", nil, err.Error(), http.StatusBadRequest)
	}

	return nil
}

func GetTopStringArg(args []interface{}) string {
	if args != nil && len(args) > 0 {
		return fmt.Sprintf("%v", args[0])
	}

	return ""
}

func GetTopIntArg(args []interface{}) int {
	var v = 0
	if str := GetTopStringArg(args); str != "" {
		v, _ = strconv.Atoi(str)
	}

	return v
}
