package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}
	//ast.Print(nil, f.Decls)
	for _, decl := range f.Decls {
		fmt.Printf("%T\n", decl)

		//if s, ok := decl.(*ast.GenDecl); ok && s.Tok == token.IMPORT {
		//	for _, v := range s.Specs {
		//		fmt.Println(v.(*ast.ImportSpec).Path.Value)
		//	}
		//}
	}

}

const src = `package foo

import (
	"fmt"
	"time"
)

// wire a is a int
var a int

func bar() {
	fmt.Println(a, time.Now())
}
`

// 0  *ast.File {
// 1  .  Package: 1
// 2  .  Name: *ast.Ident {
// 3  .  .  NamePos: 9
// 4  .  .  Name: "foo"
// 5  .  }
// 6  .  Scope: *ast.Scope {
// 7  .  .  Objects: map[string]*ast.Object (len = 0) {}
// 8  .  }
// 9  }
