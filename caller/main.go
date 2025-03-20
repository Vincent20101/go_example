package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getFunctionName() string {
	pc, _, _, ok := runtime.Caller(2) // 1 表示获取调用者的栈帧
	if !ok {
		return "unknown"
	}
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName)
}

func middleware() string {
	return getFunctionName()
}
func myFunction() {
	fmt.Println("Current function name:", middleware())
}

func main() {

	supi := "imsi-3423423424234234"
	s := strings.Split(supi, " ")
	fmt.Println(s[0])
	myFunction()
	fmt.Println(path.Base("gitlab.casa-systems.com/mobility/smf/infra/protocol/gen/nnrf/disc.(*NFInstancesStoreApiService).SearchNFInstances"))
}
