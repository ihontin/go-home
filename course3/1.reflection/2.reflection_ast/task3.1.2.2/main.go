package main

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	"github.com/dave/dst/decorator/resolver/guess"
	east "gitlab.com/ptflp/goast"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename) + "/models.go"
	data, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fset := token.NewFileSet()
	// создаем декоратор
	d := decorator.NewDecoratorWithImports(fset, "models", goast.New())

	f, err := d.Parse(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// создаем методы TableName, которые возвращают строку, используя east.Method
	//func (u *User) TableName() string { return "users" }
	tableNameMethod := east.Method{
		Name:         "TableName",
		Receiver:     "u",
		ReceiverType: "User",
		Arguments:    []east.Param{},
		Return: []east.Param{
			{
				Name: "",
				Type: "string",
			},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: "\"users\"",
						},
					},
				},
			},
		},
	}

	//func (a *Address) TableName() string { return "address" }
	tableNameMethodA := east.Method{
		Name:         "TableName",
		Receiver:     "a",
		ReceiverType: "Address",
		Arguments:    []east.Param{},
		Return: []east.Param{
			{
				Name: "",
				Type: "string",
			},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: "\"address\"",
						},
					},
				},
			},
		},
	}
	_, err = east.AddMethod(f, "User", tableNameMethod)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*methods*/ _, err = east.AddMethod(f, "Address", tableNameMethodA)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(methods)

	// получаем структуры из файла с помощью east.GetStructs
	structs := east.GetStructs(f)

	// добавляем теги в структуры с помощью east.ModifyStructs, east.AddDBTags, east.AddDBTypeTags
	err = east.ModifyStructs(structs, east.AddDBTags, east.AddDBTypeTags)
	if err != nil {
		fmt.Println(err)
		return
	}

	// синхронизируем код с измененными структурами с помощью east.SyncStructs
	err = east.SyncStructs(f, structs)
	if err != nil {
		fmt.Println(err)
		return
	}

	//reult,_:=east.PrintAST(f)

	//создать путь
	//newFileDir := path.Dir(dir) + "/models"
	//err = os.MkdirAll(newFileDir, os.FileMode(0776))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//открыть файл
	//newModels, err := os.OpenFile(newFileDir+"/models.go", os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer newModels.Close()

	// сохранить результат в файл
	res := decorator.NewRestorerWithImports("models", guess.New())

	//if err = res.Fprint(newModels, f); err != nil {
	//	log.Fatal(err)
	//}
	_ = res.Print(f)
}
