package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func main01() {
	_, err := ReadCofig()
	if err != nil {
		fmt.Printf("originnal error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v", err)
		os.Exit(1)
	}
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	return nil, nil
}

func ReadCofig() ([]byte, error) {
	home := os.Getenv("HOME")
	fmt.Println("home:", home)
	config, err := ReadFile(filepath.Join(home, ".setting.xml"))
	return config, errors.WithMessage(err, "could not read config")
}

func main() {
	_, err := ReadCofig()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("stack error : %+v\n", err)
		os.Exit(1)
	}
}
