package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	SECRET  = "W0MwRTI4MjQyRUE1NUM1Qjc0RDIxMEQ5QUZBQzhFQTA2REIzQ0EzOTJd"
	AWS_ACCESS_KEY_ID = "A2A598647C54B13B2078"
	SIGNATURE_METHOD = "HmacSHA256"
	SIGNATURE_VERSION = "2"
	VERSION = "latest"
)

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	message :=""

	methon := "GET"
	IP := "10.1.240.6"
	Path:="/"
	itemSlice := sort.StringSlice{
		"SignatureMethod=" + SIGNATURE_METHOD,
		"Action=DescribeInstances",
		"Timestamp="+url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05")+".000Z"),
		"SignatureVersion=" + SIGNATURE_VERSION,
		"AWSAccessKeyId=" + AWS_ACCESS_KEY_ID,
		"Version=" + VERSION}

	sort.Sort(itemSlice)

	items := strings.Join(itemSlice, "&")
	message = methon+"\n"+ IP+"\n"+Path+"\n"+items

	fmt.Println(HmacSha256(message, SECRET))
	fmt.Println(url.QueryEscape(HmacSha256(message, SECRET)))
}