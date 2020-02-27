package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ApplicationHandler func(c Connection, args interface{}) (Response, *AppError)
type ApplicationArgsParser func(c Connection, args ...interface{}) interface{}

type ApplicationObject map[string]interface{}
type Applications []ApplicationObject
type ApplicationHandlers map[string]*Application

type Response interface {
	String() string
}

type Application struct {
	AllowNoConnect bool
	Handler        ApplicationHandler
	ArgsParser     ApplicationArgsParser
}

func (j Applications) Value() (driver.Value, error) {
	str, err := json.Marshal(j)
	return string(str), err
}

func (j *Applications) Scan(src interface{}) error {
	if bytes, ok := src.([]byte); ok {
		return json.Unmarshal(bytes, &j)
	}
	return errors.New(fmt.Sprintf("unmarshal %v", src))
}

func (a ApplicationHandlers) Has(id string) bool {
	if _, ok := a[id]; ok {
		return true
	} else {
		return false
	}
}

func (a ApplicationHandlers) Register(id string, handler ApplicationHandler, allowNoConnect bool) *AppError {
	if a.Has(id) {
		return NewAppError("Applications", "applications.register.exists", nil,
			fmt.Sprintf("application %s handler exists", id), 500)
	}

	a[id] = &Application{
		AllowNoConnect: allowNoConnect,
		Handler:        handler,
	}

	return nil
}

func UnionApplicationMap(src ...ApplicationHandlers) ApplicationHandlers {
	res := make(ApplicationHandlers)
	for _, m := range src {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
