// 创建的M正常情况下是无法被回收
package main

import "C"
import (
	"net/http"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void output(char *str){
	usleep(1000000);
	printf("%s\n",str);
}
 */
import "C"
import (
	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe(":9999",nil)
}

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			str := "hello cgo"
			//change to char*
			cstr := C.CString(str)
			C.output(cstr)
			C.free(unsafe.Pointer(cstr))
		}()
	}
	select {}
}
