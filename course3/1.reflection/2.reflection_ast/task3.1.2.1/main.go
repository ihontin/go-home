package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"runtime"
)

type MyStruct struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Address   string
	Status    int
	DeletedAt string
}

func main() {
	fset := token.NewFileSet()
	_, filename, _, _ := runtime.Caller(0)
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if typeSpec.Name.Name == "MyStruct" {
				typeSpec.Name.Name = "User"
			}
		}
		return true
	})
	printer.Fprint(os.Stdout, fset, file)
}
