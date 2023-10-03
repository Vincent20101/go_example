package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-version"
)

func main() {
	image := "registry.gitlab.casa-systems.com/platform/operators/axyom-sidecar:v0.10.1"
	parts := strings.Split(image, ":")
	fmt.Println(parts)

	repository := strings.Join(parts[:len(parts)-1], ":")
	fmt.Println(repository)
	fmt.Println(strings.Join([]string{"author", "v8.M10.0"}, ":"))

	var (
		smfsm8M80_tag string = "1.20.25"
		tag           string = "1.19.2"
	)

	v1, err1 := version.NewVersion(tag)
	v2, err2 := version.NewVersion(smfsm8M80_tag)
	if err1 == nil && err2 == nil && v1.LessThanOrEqual(v2) {
		fmt.Println("v1 less than or equal v2")
	}

	fmt.Println(os.Remove("./ttt"))

	type UUC struct {
		total    int64
		upload   int64
		download int64
	}

}
