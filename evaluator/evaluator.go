package evaluator

import (
	"github.com/ilya-tokar/interpreter/ast"
	"github.com/ilya-tokar/interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
