package test

import "go/ast"

/*
   创建日期：2018/4/11
   修改日期：2018/4/11
   作者：Jason
*/

type AAA interface {
	bbb() error
	ccc(file ast.File)
}
