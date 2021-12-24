package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

const (
	secret  = "W0MwRTI4MjQyRUE1NUM1Qjc0RDIxMEQ5QUZBQzhFQTA2REIzQ0EzOTJd"
)

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	//fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	message :=""
	methon := "GET"
	IP := "10.1.240.6"
	Path:="/"
	items := "AWSAccessKeyId=A2A598647C54B13B2078&Action=DescribeAvailabilityZones&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2021-12-24T08%3A05%3A35.000Z&Version=latest"

	message = methon+"\n"+ IP+"\n"+Path+"\n"+items
	fmt.Println(HmacSha256(message, secret))
	fmt.Println(url.QueryEscape(HmacSha256(message, secret)))
}