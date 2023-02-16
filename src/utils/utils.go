package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

// LogFormat LogFormat
type LogFormat struct{}

func (logFormat *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer

	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	buffer.WriteString("[" + entry.Level.String()[0:4] + "] ")
	buffer.WriteString(entry.Time.Format("2006/01/02 15:04:05.000 "))

	for key, value := range entry.Data {
		buffer.WriteByte('[')
		buffer.WriteString(key)
		buffer.WriteByte(':')
		fmt.Fprint(buffer, value)
		buffer.WriteString("] ")
	}

	buffer.WriteString(entry.Message)
	buffer.WriteByte('\n')

	return buffer.Bytes(), nil
}

func HasData[T any](slice []T) bool {
	LenOfSlice := len(slice)
	if LenOfSlice > 0 {
		return true
	} else {
		return false
	}
}

func HasDataWithMin[T any](slice []T, min int) bool {
	LenOfSlice := len(slice)
	if LenOfSlice > min {
		return true
	} else {
		return false
	}
}

func RecoverError() {
	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Printf("")
		}
	}()
	panic("Error")
}

func RemoveLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func PrintProcessDay(Logger *logrus.Entry, start, from, to time.Time) {
	Logger.Debugf("[EventService] End of process day:[" + from.Format("01-02-2006") + "]")
	Logger.Debugf("[EventService] Ejecution time    :[" + strconv.FormatFloat(time.Since(start).Minutes(), 'f', 1, 64) + "]")
	Logger.Debugf("")
}

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func RemoveFrom(s interface{}, idx int) interface{} {
	if v := reflect.ValueOf(s); v.Len() > idx {
		return reflect.AppendSlice(v.Slice(0, idx), v.Slice(idx+1, v.Len())).Interface()
	}
	return s
}
