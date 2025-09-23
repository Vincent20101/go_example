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
	arr := strings.Split("lo:smfsm,eth0:smfsm", ":")
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println("============")
	fmt.Println(ParseInterfaceString("lo:smfsm,eth0:smfsm"))
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

func ParseInterfaceString(str string) (intf string, serviceName string, podName string) {
	arr := strings.Split(str, ":")
	if len(arr) > 3 || len(arr) == 0 {
		return "", "", ""
	}

	intf = arr[0]

	if len(arr) > 1 {
		serviceName = arr[1]
	}

	if len(arr) > 2 {
		podName = arr[2]
	}

	return
}
