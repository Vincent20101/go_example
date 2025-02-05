package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	path := "/nudm-sdm/v2/notification/imsi-450051230109838/sm-data"
	fmt.Println(strings.Split(path, "/"))
	fmt.Println(strings.Split(path, "/")[1])
	fmt.Println(GetSessIdFromPath("/nsmf-pdusession/{apiVersion}/sm-contexts/{pduSessionRef}/modify"))
	fmt.Println(GetSessIdFromPath("/{apiPrefix}/nsmf-pdusession/{apiVersion}/sm-contexts/{pduSessionRef}/modify"))
}

func GetSessIdFromPath(path string) (sessId string, err error) {
	///nsmf-pdusession/v1/sm-contexts/{smContextRef}/[release/update/retrieve]
	sl := strings.Split(path, "/")
	if len(sl) < 5 {
		err = errors.New(fmt.Sprintf("wrong uri [path: %v] received", path))
		return
	}
	sessId = sl[len(sl)-2]
	return sessId, nil
}
