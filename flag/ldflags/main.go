package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

var buildstamp = ""
var githash = ""

const GX_ULI_TYPE_ECGI = 129

func main() {
	fmt.Println("lhb")
	debug.SetTraceback("single")
	fmt.Println(string(debug.Stack()))
	fmt.Println(os.Getenv("GOTRACEBACK"))
	fmt.Println("lhb")
	var el []byte
	el = append(el, GX_ULI_TYPE_ECGI)
	fmt.Println(el)
	workPath, getWdErr := os.Getwd()
	if getWdErr != nil {
		fmt.Printf("OS.Getwd Failed, %v\n", getWdErr)
		os.Exit(1)
	}
	fmt.Println(workPath)
	args := os.Args
	fmt.Println(strings.Join(os.Args, " "))
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s\n", githash)
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		return
	}
}

// flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git describe --long --dirty --abbrev=14`"
// go build -ldflags "$flags" -x -o build-version main.go
// ./build-version -v
