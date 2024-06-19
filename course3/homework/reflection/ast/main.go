package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
		package main

		import "fmt"


		func trizna(){fmt.Println("life is shert")} 
		func main() {
			fmt.Println("Hello, world!")
			trizna()
		}
		
	`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if ident, ok2 := callExpr.Fun.(*ast.Ident); ok2 {
				fmt.Println("Вызов функции:", ident.Name)
			}
		}
		return true
	})
}

//В этом примере мы используем пакет ast для анализа исходного кода программы.
//Мы передаем исходный код в функцию parser. ParseFile, которая возвращает абстрактное синтаксическое дерево программы.
//Затем мы используем функцию ast. Inspect, чтобы пройтись по всему дереву и найти вызовы функций.
//В этом случае мы выводим имена вызываемых функций.

//func main() {
//	fset := token.NewFileSet()
//
//	// Создаем новый пакет
//	pkg := &ast.Package{
//		Name:  "main",
//		Files: make(map[string]*ast.File),
//	}
//
//	// Создаем новый файл
//	file := &ast.File{
//		Name:  ast.NewIdent("main"),
//		Decls: []ast.Decl{},
//	}
//
//	// Создаем новую функцию
//	funcDecl := &ast.FuncDecl{
//		Name: ast.NewIdent("main"),
//		Type: &ast.FuncType{},
//		Body: &ast.BlockStmt{
//			List: []ast.Stmt{
//				&ast.ExprStmt{
//					X: &ast.CallExpr{
//						Fun: &ast.SelectorExpr{
//							X:   ast.NewIdent("fmt"),
//							Sel: ast.NewIdent("Println"),
//						},
//						Args: []ast.Expr{
//							&ast.BasicLit{
//								Kind:  token.STRING,
//								Value: "\"Hello, world!\"",
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	// Добавляем функцию в файл
//	file.Decls = append(file.Decls, funcDecl)
//
//	// Добавляем файл в пакет
//	pkg.Files["main.go"] = file
//
//	// Генерируем код
//	err := printer.Fprint(os.Stdout, fset, file)
//	if err != nil {
//		fmt.Println("Ошибка:", err)
//		return
//	}
//}

//В этом примере мы используем пакет ast для генерации кода. Мы создаем новый пакет, файл и функцию,
//а затем добавляем функцию в файл и файл в пакет. Затем мы используем функцию printer.
//Fprint для генерации кода и выводим его в стандартный вывод. В этом случае мы генерируем код для функции main,
//которая выводит “Hello, world!”.
