package EasyLogger

import (
	"fmt"
	"time"
	"os"
	"net/http"
	"encoding/json"
	"bytes"
	"runtime"
	"strconv"
	"net/http/httputil"
)

type Log struct {
	LogID string
	AppName string
	Date time.Time
	LogLevel LogLevel
	Message string
	Object string
	HttpBodyDump string
	FuncCall string
}

type LogLevel string

const (
	Info LogLevel = "INFO"
	Warn LogLevel = "WARN"
	Error LogLevel = "ERROR"
	Fatal LogLevel = "FATAL"
	Debug LogLevel = "DEBUG"
	Trace LogLevel = "TRACE"
)

func acceptLog(log Log) error {

	body, err := json.Marshal(log)
	object, err := json.Marshal(log.Object)
	log.Object = string(object)

	_, err = http.Post(fmt.Sprintf("%s/%s/%s", os.Getenv("ES_HOST"), os.Getenv("APP_NAME"), log.LogID),
		"application/json", bytes.NewBuffer(body))

	return err
}

func LogIt(magnitude LogLevel, message string, object interface{}) error {

	pc, file, line, _ := runtime.Caller(1)
	funcCall := runtime.FuncForPC(pc).Name()

	return acceptLog(Log{LogID:fmt.Sprintf("%v", time.Now().UnixNano()), AppName:os.Getenv("APP_NAME"),
	Date:time.Now().UTC(),
	LogLevel:magnitude,
	Message:message, Object:fmt.Sprintf("%v", object), FuncCall:fmt.Sprintf("FILE:%s, FUNC:%s, LINE:%s", file, funcCall, strconv.Itoa(line))})
}

func LogItWithHttpResponse(magnitude LogLevel, message string, object interface{}, resp *http.Response) error {

	pc, file, line, _ := runtime.Caller(1)
	funcCall := runtime.FuncForPC(pc).Name()

	body, _ := httputil.DumpResponse(resp, true)

	return acceptLog(Log{HttpBodyDump:string(body), LogID:fmt.Sprintf("%v", time.Now().UnixNano()), AppName:os.Getenv("APP_NAME"),
	Date:time.Now().UTC(), LogLevel:magnitude, Message:message, Object:fmt.Sprintf("%v", object),
	FuncCall:fmt.Sprintf("FILE:%s, FUNC:%s, LINE:%s", file, funcCall, strconv.Itoa(line))})
}

func LogItWithHttpRequest(magnitude LogLevel, message string, object interface{}, resp *http.Request) error {

	pc, file, line, _ := runtime.Caller(1)
	funcCall := runtime.FuncForPC(pc).Name()

	body, _ := httputil.DumpRequest(resp, true)

	return acceptLog(Log{HttpBodyDump:string(body), LogID:fmt.Sprintf("%v", time.Now().UnixNano()), AppName:os.Getenv("APP_NAME"),
		Date:time.Now().UTC(), LogLevel:magnitude, Message:message, Object:fmt.Sprintf("%v", object),
		FuncCall:fmt.Sprintf("FILE:%s, FUNC:%s, LINE:%s", file, funcCall, strconv.Itoa(line))})
}
