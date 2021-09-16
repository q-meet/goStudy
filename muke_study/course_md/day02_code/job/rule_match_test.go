package main

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// Eval : 判断 map 是否符合 bool 表达式
//	expr = `a > 1 && b < 0`
func Eval(m map[string]string, expr string) (bool, error) {
	fest := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, err
	}
	ast.Print(fest, exprAst)

	return judge(exprAst, m), nil
}

func isLeaf(bop ast.Node) bool {
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	_, okL := expr.X.(*ast.Ident)
	_, okR := expr.Y.(*ast.BasicLit)
	if okL && okR {
		return true
	}
	return false
}

func comparison(x string, y string, op token.Token) bool {
	switch op {
	case token.EQL:
		return x == y
	case token.LSS:
		return x < y
	case token.GTR:
		return x > y
	case token.NEQ:
		return x != y
	}
	return false
}

// dfs
func judge(bop ast.Node, m map[string]string) bool {
	if isLeaf(bop) {
		// do the leaf logic
		expr := bop.(*ast.BinaryExpr)
		x := expr.X.(*ast.Ident)
		y := expr.Y.(*ast.BasicLit)

		// 获取符号 常量值
		op := expr.Op

		return comparison(m[x.Name], y.Value, op)
		// FIXME，修正这里的逻辑，使 test 能够正确通过
		//return m[x.Name] == y.Value
	}
	// not leaf
	// 那么一定是 binary expression
	expr, ok := bop.(*ast.BinaryExpr)
	//expr, ok := bop.(*ast.BinaryExpr)

	if !ok {
		expr1, ok1 := bop.(*ast.ParenExpr)
		if !ok1 {
			println("this cannot be true")
			return false
		}
		expr = expr1.X.(*ast.BinaryExpr)
	}
	/*
	x1 := expr.X.(*ast.BinaryExpr)
	x2 := x1.X.(*ast.Ident)
	println("name", x2.Name)
	*/

	switch expr.Op {
	case token.LAND:
		return judge(expr.X, m) && judge(expr.Y, m)
	case token.LOR:
		return judge(expr.X, m) || judge(expr.Y, m)
	}

	println("unsupported operator")
	return false
}


type testCase struct {
	m      map[string]string
	expr   string
	result bool
}

func TestMapExpr(t *testing.T) {
	cases := []testCase{
		{
			m: map[string]string{
				"invest": "0", "posts": "11",
			},
			expr:   `posts > 10`,
			result: true,
		},
		{
			m: map[string]string{
				"invest": "20000", "posts": "144",
			},
			expr:   `invest > 1000`,
			result: true,
		},
		{
			m: map[string]string{
				"invest": "150", "posts": "10000",
			},
			expr:   `(invest > 10000 && posts > 100) || (posts > 10000 || (invest != 0 && posts > 10))`,
			result: true,
		},
	}

	for _, cas := range cases {
		res, err := Eval(cas.m, cas.expr)
		assert.Nil(t, err)
		assert.Equal(t, res, cas.result)
	}
}
