package main

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	test()
}

func test() {
	fmt.Println("===testing")
	logFile()
	fmt.Println(path.Join("/", "configs", "template", "/etc/main_launcher.json"))
}

func logFile() {
	var funcName string
	var buf bytes.Buffer
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	} else {
		f := runtime.FuncForPC(pc)
		if f != nil {
			funcName = f.Name()
			funcNameSp := strings.Split(funcName, ".")
			if len(funcNameSp) > 0 {
				funcName = funcNameSp[len(funcNameSp)-1]
			}
		}
	}

	buf.WriteString(`"casa":{"src":{"file":"`)
	buf.WriteString(file)
	buf.WriteString(`",`)
	buf.WriteString(`"line":`)
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(`,`)
	buf.WriteString(`"function":"`)
	buf.WriteString(funcName)
	buf.WriteString(`"}}`)

	// message
	buf.WriteString(`,"message": test"`)

	fmt.Println(buf.String())
}
