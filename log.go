package nutils

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func WriteLog(logFileName, event string) {

	t := time.Now()
	today := t.Day()
	old := false
	var dir string
	var logname string

	if strings.Contains(logFileName, string(os.PathSeparator)) {
		dir = path.Dir(logFileName)
	} else {
		dir, _ = os.Getwd()
		dir += string(os.PathSeparator) + "log"
		logFileName = dir + string(os.PathSeparator) + logFileName
	}
	_, err := os.Stat(dir)
	if (err != nil) && (os.IsNotExist(err)) {
		os.Mkdir(dir, 0777)
	}

	// Check current log date, if it is old, overwrite it
	logname = logFileName + "-" + strconv.Itoa(today) + ".log"
	logstat, err := os.Stat(logname)
	if err == nil {
		if t.Month() != logstat.ModTime().Month() {
			old = true
		}
	}
	var f *os.File
	if old {
		os.Remove(logname)
		f, _ = os.OpenFile(logname, os.O_CREATE|os.O_RDWR, 0666)

	} else {
		f, _ = os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	}
	_, er := f.WriteString(t.String()[1:22] + ": " + event + "\n")
	if er != nil {
		println("Error in writing log: ", er.Error())
	}
	f.Close()

}

func WriteServerLog(logFileName, ip, method, queryParams string, statusCode int) {
	t := time.Now()
	day := t.Day()
	month := int(t.Month())
	year := t.Year()

	var dir string

	if strings.Contains(logFileName, string(os.PathSeparator)) {
		dir = path.Dir(logFileName)
	} else {
		dir, _ = os.Getwd()
		dir += string(os.PathSeparator) + "serverlog"
		logFileName = dir + string(os.PathSeparator) + logFileName
	}

	_, err := os.Stat(dir)
	if (err != nil) && (os.IsNotExist(err)) {
		os.Mkdir(dir, 0777)
	}

	logname := logFileName + "-" + strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day) + ".txt"

	f, _ := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()

	_, er := f.WriteString(ip + ` -- [` + t.Format("2/Jan/2006:15:04:05") + `] "` + method + ` ` + queryParams + `" ` + strconv.Itoa(statusCode) + " -\n")
	if er != nil {
		println("Error in writing log: " + er.Error())
	}

}
