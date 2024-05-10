package pkg

import "os"

var PACKAGE_PATH string

func init() {
	//PACKAGE_PATH = os.Getenv("GOPATH")

	getPath()
}

func getPath() {
	PACKAGE_PATH = os.Getenv("GOROOT")
}
