package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/Jasonzz-he/script/base"
)

/*
    创建日期：2018/4/11
    修改日期：2018/4/11
    作者：Jason
	目的：接口实现类
*/

var (
	fileName string
	fileSet  = token.NewFileSet()
)

func Flags() {
	flag.StringVar(&fileName, "filename", "", "文件名称，如果没有则扫描文件夹")
}

func main() {
	Flags()
	flag.Parse()
	var astFile *ast.File
	var err error
	if len(fileName) > 0 && strings.HasSuffix(fileName, ".go") {
		astFile, err = parseFile(fileName)
	} else {
		astFile, err = parseDir("./")
	}
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	// 解析语法树
	base.GetAstFileInstance().ParseAstFile(astFile)

}

// 解析文件获取语法树
func parseFile(fileName string) (*ast.File, error) {
	astFile, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if nil != err {
		er := fmt.Errorf("parseFile [ fileName %v ] err: %v", fileName, err)
		return nil, er
	}
	return astFile, nil
}

// 解析目录获取文件语法树
func parseDir(dir string) (*ast.File, error) {
	return nil, errors.New("not impl parseDir")
}
