package main

import (
	"fmt"
	"strings"
	"time"
)

var formatArr = []string{"2006", "2006年", "2006-01-02", "2006/01/02", "", "2006-01-02 15:04:05", "2006-01-02T15:04:05",
	"2006-01-02T15:04:05Z", "Jan 02 15:04:05 2006 GMT", "1月 02 15:04:05 2006 GMT"}

// StrToDate 时间转换
func StrToDate(dateStr string) time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if len(dateStr) == 0 {
		return time.Time{}
	}
	for _, format := range formatArr {
		dateTime, err := time.ParseInLocation(format, strings.TrimSpace(dateStr), time.Local)
		if err != nil {
			continue
		}
		return dateTime
	}
	return time.Time{}
}

func main() {
	year := time.Now().Year() // 获取当前年份
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year+1, time.January, 1, 23, 59, 59, 0, time.UTC).AddDate(0, 0, -1)

	fmt.Println("一年的开始日期：", startDate.Format("2006-01-02 15:04:05"))
	fmt.Println("一年的结束日期：", endDate.Format("2006-01-02 15:04:05"))

	time := StrToDate("2023年")
	fmt.Println(time.Year())
}
