package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	//return errors.Wrap(sql.ErrNoRows, "foo failed")
	return errors.WithStack(sql.ErrNoRows)
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main0101() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Println("====================")
		fmt.Printf("%+v\n", err)
		fmt.Println("====================")
		return
	}
	if err != nil {
		// unknown error
	}
}

func main() {
	err1 := errors.New("第一个错误")
	err2 := errors.Wrap(err1, errors.New("第二个错误").Error())
	fmt.Println(errors.Unwrap(err2))
}
