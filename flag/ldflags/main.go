package main

import (
	"fmt"
	"os"
)

var buildstamp = ""
var githash = ""

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s\n", githash)
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		return
	}
}

// flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git describe --long --dirty --abbrev=14`"
// go build -ldflags "$flags" -x -o build-version main.go
// ./build-version -v
