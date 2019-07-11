package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hpcloud/tail"
	//"strings"
	"encoding/json"
	"strings"
	"time"
)

type Fileds struct {
	Generate       string     `json:"generate"`
	Tos            string     `json:"tos"`
	Level          int        `json:"level"`
	Endpoint       string     `json:"endpoint"`
	Formatted_time string     `json:"formatted_time"`
	Template_id    string     `json:"template_id"`
	Date           *time.Time `json:"date"`
	Ip             string     `json:"ip"`
	Metric         string     `json:"metric"`
}

var Db *sql.DB

func main() {

	//errs := getJsonData()
	//if errs != nil {
	//	fmt.Println("get file error:", errs)
	//}
	filename := "/var/log/alarm_minions/new_alarm.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file error:", err)
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail filename close :%s\n", tails)
			time.Sleep(1000 * time.Millisecond)
			continue
		}
		fmt.Printf("Get line is: %v\n", msg.Text)
		var f Fileds
		json.Unmarshal([]byte(msg.Text), &f)
		if len(msg.Text) <= 200 {
			fmt.Println("The len is <= 200 ")
		} else {
			ip := getIp(f.Endpoint)
			res, err := Db.Exec("INSERT into alarm_info (generate,tos,level,endpoint,formatdate,template_id,date,ip,metric) VALUES (?,?,?,?,?,?,?,?,?)", f.Generate, f.Tos, f.Level, f.Endpoint, f.Formatted_time, f.Template_id, time.Now().Format("2006-01-02 15:04:05"), ip, f.Metric)
			if err != nil {
				fmt.Println("error:", err)

			}
			id, err := res.LastInsertId()
			fmt.Printf("The id is %d\n", id)

		}

	}
}

func getIp(endpoint string) string {
	var l []string
	l = strings.Split(endpoint, "-")
	return l[len(l)-1]
}

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:quanshi@tcp(192.168.5.19:3306)/gaea?charset=utf8")
	if err != nil {
		fmt.Println("something error:", err)
	}
	//defer Db.Close()
}

//func getJsonData() error {
//	filename := "/var/log/alarm_minions/alarm.log"
//	tails, err := tail.TailFile(filename, tail.Config{
//		ReOpen:    true,
//		Follow:    true,
//		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
//		MustExist: false,
//		Poll:      true,
//	})
//
//	if err != nil {
//		fmt.Println("tail file error:", err)
//		return err
//	}
//	var msg *tail.Line
//	var ok bool
//	for true {
//		msg, ok = <-tails.Lines
//		if !ok {
//			fmt.Printf("tail filename close :%s\n", tails)
//			time.Sleep(1000 * time.Millisecond)
//			continue
//		}
//		errs := writeNewFile(msg.Text, "/var/log/alarm_minions/new_alarm.log")
//		if errs != nil {
//			return errs
//		}
//	}
//	return err
//}

//func writeNewFile(line string, filePath string) error {
//	jsonString := strings.Split(line, " ")
//	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
//	if err != nil {
//		fmt.Println("some error msg:", err)
//		return err
//	}
//	lineStr := fmt.Sprintf("%v\n", jsonString[1])
//	n, err := f.WriteString(lineStr)
//	fmt.Println("Tatol is: ", n)
//	if err != nil {
//		fmt.Printf("What Wrong. %s\n", err)
//	}
//	defer f.Close()
//	return err
//}
