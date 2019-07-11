package main

import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
	//"strings"
	"io/ioutil"
	"net/url"
	"strings"
)

var db *sql.DB
var (
	tos      string
	ip       string
	endpoint string
	v        int
	metric   []uint8
)

func main() {
	n := time.Now()
	yesterDay := n.AddDate(0, 0, -1)
	fmt.Println(yesterDay.Format("2006-01-02"))
	currentDay := yesterDay.Format("2006-01-02")
	initialTime := currentDay + " 00:00:00"
	fmt.Println(initialTime)

	endData := currentDay + " 23:59:59"
	fmt.Println(endData)
	resu, err := db.Query("select tos,ip,endpoint ,metric,count(endpoint) as v from alarm_info where date between  ?  and  ? GROUP BY endpoint order by v DESC LIMIT 5", initialTime, endData)
	if err != nil {
		fmt.Println("something error", err)
	}
	defer resu.Close()
	for resu.Next() {
		err := resu.Scan(&tos, &ip, &endpoint, &metric, &v)
		if err != nil {
			fmt.Println("Ther error:", err)
		}
		fmt.Println("ip: , endpoint: , count:  metric :", ip, endpoint, v, string(metric))

		mess := fmt.Sprintf("您下辖主机昨日报警过多请酌情优化:\n %-6s: %-6s\n %-6s: %-6s\n %-6s: %-6v\n %-6s: %-6v\n", "ip_address", string(ip), "endpoint", string(endpoint), "metric", string(metric), "昨日总计发送消息", v)
		usersList := strings.Split(tos, ",")
		for _, name := range usersList {
			fmt.Printf("name is: %v\n", name)
			resp, err := http.PostForm("http://192.168.5.11:8888/api/v1/sender/bee?action=bee", url.Values{"tos": {"yalei.zhang@quanshi.com"}, "content": {string(mess)}, "tag": {"cmdb"}})
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}

			fmt.Println(string(body))
		}

	}

}
func init() {
	var err error
	db, err = sql.Open("mysql", "root:quanshi@tcp(192.168.5.19:3306)/gaea?charset=utf8")
	if err != nil {
		fmt.Println("something error:", err)
	}
}
