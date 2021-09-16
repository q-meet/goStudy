package julian

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Eval : 判断 map 是否符合 bool 表达式
//	expr = `a > 1 && b < 0`
func Eval(m map[string]string, expr string) (bool, error) {
	fset := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, err
	}

	ast.Print(fset, exprAst)
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

// dfs
func judge(bop ast.Node, m map[string]string) bool {
	if isLeaf(bop) {
		// do the leaf logic
		expr := bop.(*ast.BinaryExpr)
		x := expr.X.(*ast.Ident)
		y := expr.Y.(*ast.BasicLit)

		a, _ := strconv.Atoi(m[x.Name])
		b, _ := strconv.Atoi(y.Value)
		switch expr.Op {
		case token.GTR:
			return a > b
		case token.GEQ:
			return a >= b
		case token.LSS:
			return a < b
		case token.LEQ:
			return a <= b
		case token.EQL:
			return a == b
		}
	}
	// for parenese
	parenExpr, ok:= bop.(*ast.ParenExpr)
	if ok{
		return judge(ast.Node(parenExpr.X), m)
	}

	// not leaf
	// 那么一定是 binary expression
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		println("this cannot be true")
		return false
	}

	a := judge(expr.X, m)
	b := judge(expr.Y, m)
	switch expr.Op {
	case token.LAND:
		return a && b
	case token.LOR:
		return a || b
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
		testCase{
			m: map[string]string{
				"invest": "15000", "posts": "101",
			},
			expr:   `invest > 15000 || (posts > 100  && invest > 10000)`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "15000", "posts": "101",
			},
			expr:   `invest > 15000 || posts > 100  && invest > 10000`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "15000", "posts": "99",
			},
			expr:   `(invest > 15000 || posts > 100 ) && invest > 10000`,
			result: false,
		},
		testCase{
			m: map[string]string{
				"invest": "20000", "posts": "144",
			},
			expr:   `(invest > 10000) && (posts > 100)`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "20000", "posts": "144",
			},
			expr:   `invest > 10000 && posts > 100`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "20000", "posts": "144",
			},
			expr:   `invest > 10000 && posts < 140`,
			result: false,
		},
		testCase{
			m: map[string]string{
				"invest": "20000", "posts": "139",
			},
			expr:   `invest >= 10000 && posts < 140`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "20000", "posts": "139",
			},
			expr:   `invest >= 10000 && posts <= 140`,
			result: true,
		},
		testCase{
			m: map[string]string{
				"invest": "2000", "posts": "144",
			},
			expr:   `invest > 10000 && posts > 100`,
			result: false,
		},
	}

	for _, cas := range cases {
		res, err := Eval(cas.m, cas.expr)
		assert.Nil(t, err)
		assert.Equal(t, res, cas.result)
	}
}
