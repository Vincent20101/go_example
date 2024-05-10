package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	path := "/tmp/tracing"
	if _, err := os.Stat(path); err != nil {
		fmt.Println(err)
		if err = os.MkdirAll(path, fs.ModePerm); err != nil {
			fmt.Println(err)
		}
	}
}
