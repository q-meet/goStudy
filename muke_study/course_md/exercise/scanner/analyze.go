package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	//要分析的代码
	var src = []byte(`println("你好，世界")`)

	//通过token.NewFileSet()创建一个文件集 Token的位置信息必须通过文件集定位
	//并且需要通过文件集创建扫描器的Init方法需要的File参数。
	var fset = token.NewFileSet()
	//fset文件集添加一个新的文件，文件名为“hello.go”，文件的长度就是src要分析代码的长度。
	var file = fset.AddFile("hello.go", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
