package main

import (
	"fmt"
)

func findIdentifierLiteralNode(lval LiteralNode) {

	fmt.Println("chenyao**************IdentifierLiteralNode is ", lval.identifierLiteralNode.value)
}

func aa() ExprNode {
	return ExprNode{}
}
