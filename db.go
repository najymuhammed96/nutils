package nutils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

func InitDB(ip, port, dbname, user, password, driver string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if driver == "mysql" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+ip+":"+port+")/"+dbname+"?charset=utf8")
	} else if driver == "mssql" {
		connString := fmt.Sprintf("server=%s;user id=%s; password=%s;portNumber=%s;database=%s;encrypt=disable", ip, user, password, port, dbname)
		db, err = sql.Open("mssql", connString)
	} else {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+ip+":"+port+")/"+dbname+"?charset=utf8")
	}
	return db, err
}

func MapScan(rows *sql.Rows, columns []string) map[string]interface{} {
	resMap := make(map[string]interface{})
	records := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range records {
		columnPointers[i] = &records[i]
	}

	rows.Scan(columnPointers...)
	for i := range columns {
		val := columnPointers[i].(*interface{})
		actVal, ok := (*val).([]uint8)
		if ok {
			resMap[columns[i]] = string(actVal)
		} else if reflect.TypeOf(*val) == nil {
			resMap[columns[i]] = nil
		} else {
			switch reflect.TypeOf(*val).String() {
			case "int64":
				resMap[columns[i]], _ = (*val).(int64)
			case "int32":
				resMap[columns[i]], _ = (*val).(int32)
			case "float64":
				myfloat, _ := (*val).(float64)
				resMap[columns[i]] = json.Number(fmt.Sprintf("%.1f", myfloat))
			case "float32":
				myfloat, _ := (*val).(float32)
				resMap[columns[i]] = json.Number(fmt.Sprintf("%.1f", myfloat))
			default:
				resMap[columns[i]] = fmt.Sprintf("%v", *val)
			}
		}
	}
	return resMap
}

func PrepareAndInsertDataBulk(l [][]interface{}, f func(data [][]interface{}), max int) {
	list := l
	sublist := make([][]interface{}, 0)
	for {
		if len(list) > max {
			sublist = list[:max]
			list = list[max:]
			f(sublist)
		} else if len(list) > 0 {
			f(list)
			break
		} else {
			break
		}
	}
}
