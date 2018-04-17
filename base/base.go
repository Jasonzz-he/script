package base

import (
	"go/ast"
	"go/token"
)

/*
   创建日期：2018/4/11
   修改日期：2018/4/11
   作者：Jason
*/

type AstFile struct {
	PackageName string
	Imports     []Import
}

func GetAstFileInstance() *AstFile {
	return &AstFile{}
}

type Import struct {
	Name    string
	SelName string
	Body    string
}

var fset = token.NewFileSet()

// 解析语法树
func (af *AstFile) ParseAstFile(astFile *ast.File) {
	ast.Print(fset, astFile)
	//ast.Inspect(astFile, func(node ast.Node) bool {
	//	switch node := node.(type) {
	//	case *ast.File:
	//		ast.Print(fset, node)
	//	case *ast.Package:
	//		af.PackageName = node.Name
	//	case *ast.ImportSpec:
	//		ast.Print(fset, node)
	//	case *ast.TypeSpec:
	//		ast.Print(fset, node)
	//	default:
	//		fmt.Printf("node Type %T\n", node)
	//	}
	//	return true
	//})
}
