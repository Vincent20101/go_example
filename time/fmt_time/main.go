package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//var timestamp time.Time
	//itimeStr := fmt.Sprintf("%02d%02d%02d", timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	//fmt.Println(itimeStr)

	now := time.Now()
	formatStr := "1504-0700"
	fmt.Println(now.Format(formatStr))
	//fmt.Println(now.UTC())
	//fmt.Println(now.UTC().Hour())
	//timeStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d + %02d:%02d:%02d",
	//	now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.UTC().Hour(), now.UTC().Minute(), now.UTC().Second())
	//fmt.Printf("time:%s\n", timeStr)

	timeStr := fmt.Sprintf("%02d%02d", now.Hour(), now.Minute())
	// determine timezone
	dateArray := strings.Fields(now.String())
	fmt.Println(dateArray)
	// the 2nd field in the date array is the timezone offset
	timeStr += dateArray[2]
	fmt.Println(timeStr)

}
