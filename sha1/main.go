package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))

}

func main() {
	fmt.Println(SHA1("123456sdfsdfssdfsdfdsfdfsdfdsfsdfsdfdsfsdf"))
	fmt.Println(len(SHA1("123456sdfsdfssdfsdfdsfdfsdfdsfsdfsdfdsfsdf")))

	rand.Seed(time.Now().UnixNano())
	//s := strconv.FormatInt(time.Now().UnixNano(), 16) + sdEp.GetSessionKey().ToString() + string(msgType)
	s := strconv.FormatInt(time.Now().UnixNano()+rand.Int63(), 16)
	b := sha1.Sum([]byte(s))
	hexStr := hex.EncodeToString(b[:])
	fmt.Println(hexStr)
	fmt.Println(len(hexStr))

	s1 := strconv.FormatUint(100, 16)
	fmt.Println(s1)
	fmt.Println(strconv.ParseUint(s1, 10, 64))
}
