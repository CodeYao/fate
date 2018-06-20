package main

import (
	"fmt"
)

func findIdentifierLiteralNode(lval LiteralNode) {

	fmt.Println("chenyao**************IdentifierLiteralNode is ", lval.identifierLiteralNode.value)
}

func aa() StmtNode {
	return StmtNode{}
}
func bb() []ExprNode {
	return []ExprNode{}
}
func createPrefixOpNode(opcode int, unary_expression ExprNode) ExprNode {
	fmt.Println("chenyao**************unary_expression is ", unary_expression.literalNode.identifierLiteralNode.value)
	var exprNode ExprNode
	if exprNode.unaryOpNode == nil {
		exprNode.unaryOpNode = &UnaryOpNode{}
	}
	exprNode.unaryOpNode.amount = 1
	exprNode.unaryOpNode.expr = unary_expression
	if opcode == INC_OP {
		exprNode.unaryOpNode.operator = "++"
	} else if opcode == DEC_OP {
		exprNode.unaryOpNode.operator = "--"
	}
	exprNode.unaryOpNode.isPer = true
	exprNode.unaryOpNode.opType = INC_OP
	return exprNode
}

func createUnaryOpNode(opcode byte, cast_expression ExprNode) ExprNode {
	fmt.Println("chenyao**************cast_expression is ", cast_expression.literalNode.identifierLiteralNode.value)
	var exprNode ExprNode
	if exprNode.unaryOpNode == nil {
		exprNode.unaryOpNode = &UnaryOpNode{}
	}
	exprNode.unaryOpNode.expr = cast_expression
	exprNode.unaryOpNode.operator = string(opcode)
	exprNode.unaryOpNode.opType = int(opcode)
	return exprNode
}

func createSuffixOpNode(opcode int, unary_expression ExprNode) ExprNode {
	fmt.Println("chenyao**************unary_expression is ", unary_expression.literalNode.identifierLiteralNode.value)
	var exprNode ExprNode
	if exprNode.unaryOpNode == nil {
		exprNode.unaryOpNode = &UnaryOpNode{}
	}
	exprNode.unaryOpNode.amount = 1
	exprNode.unaryOpNode.expr = unary_expression
	if opcode == INC_OP {
		exprNode.unaryOpNode.operator = "++"
	} else if opcode == DEC_OP {
		exprNode.unaryOpNode.operator = "--"
	}
	exprNode.unaryOpNode.isPer = false
	exprNode.unaryOpNode.opType = INC_OP
	return exprNode
}

func createArefNode(expr ExprNode, index ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.lhNode == nil {
		exprNode.lhNode = &LHSNode{}
	}
	if exprNode.lhNode.arefNode == nil {
		exprNode.lhNode.arefNode = &ArefNode{}
	}
	exprNode.lhNode.arefNode.expr = expr
	exprNode.lhNode.arefNode.index = index
	return exprNode
}

func createFuncallNode(expr ExprNode, args []ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.funcCallNode == nil {
		exprNode.funcCallNode = &FuncCallNode{}
	}
	exprNode.funcCallNode.expr = expr
	exprNode.funcCallNode.args = args
	return exprNode
}

func createMemberNode(expr ExprNode, menber ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.lhNode == nil {
		exprNode.lhNode = &LHSNode{}
	}
	if exprNode.lhNode.memberNode == nil {
		exprNode.lhNode.memberNode = &MemberNode{}
	}
	exprNode.lhNode.memberNode.expr = expr
	exprNode.lhNode.memberNode.member = menber.literalNode.identifierLiteralNode.value
	fmt.Println("chenyao********memberNode", exprNode.lhNode.memberNode.member)
	return exprNode
}

func ceateBinaryOpNode(left ExprNode, op string, right ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.binaryOpNode == nil {
		exprNode.binaryOpNode = &BinaryOpNode{}
	}
	exprNode.binaryOpNode.left = left
	exprNode.binaryOpNode.right = right
	exprNode.binaryOpNode.operator = op
	fmt.Println("chenyao*************ceateBinaryOpNode.op ", op)
	return exprNode
}

func createCondExprNode(cond ExprNode, thenExpr ExprNode, elseExpr ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.condExprNode == nil {
		exprNode.condExprNode = &CondExprNode{}
	}
	exprNode.condExprNode.cond = cond
	exprNode.condExprNode.thenExpr = thenExpr
	exprNode.condExprNode.elseExpr = elseExpr

	return exprNode
}

func createAssignNode(lhs ExprNode, op string, rhs ExprNode) ExprNode {
	var exprNode ExprNode
	if exprNode.abstractAssignNode == nil {
		exprNode.abstractAssignNode = &AbstractAssignNode{}
	}
	if op == "=" {
		exprNode.abstractAssignNode.assignNode = &AssignNode{}
		exprNode.abstractAssignNode.assignNode.lhs = lhs
		exprNode.abstractAssignNode.assignNode.rhs = rhs
	} else {
		exprNode.abstractAssignNode.opAssignNode = &OpAssignNode{}
		exprNode.abstractAssignNode.opAssignNode.lhs = lhs
		exprNode.abstractAssignNode.opAssignNode.rhs = rhs
		exprNode.abstractAssignNode.opAssignNode.operator = op
	}
	fmt.Println("chenyao*************createAssignNode.op ", op)
	return exprNode
}

func createIfNode(cond ExprNode, thenBody StmtNode, elseBody StmtNode) StmtNode {
	var stmtNode StmtNode
	if stmtNode.ifNode == nil {
		stmtNode.ifNode = &IfNode{}
	}
	stmtNode.ifNode.cond = cond
	stmtNode.ifNode.thenBody = thenBody
	if elseBody == (StmtNode{}) {
		stmtNode.ifNode.elseBody = StmtNode{}
		fmt.Println("chenyao**********there is no else", elseBody)
	} else {
		fmt.Println("chenyao**********there is have else", elseBody)
		stmtNode.ifNode.elseBody = elseBody
	}
	return stmtNode
}

func createExprStmtNode(expr ExprNode) StmtNode {
	var stmtNode StmtNode
	if stmtNode.exprStmtNode == nil {
		stmtNode.exprStmtNode = &ExprStmtNode{}
	}
	stmtNode.exprStmtNode.expr = expr
	return stmtNode
}

func createForNode(init StmtNode, cond StmtNode, incr ExprNode, body StmtNode) StmtNode {
	fmt.Println("chenyaoXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", cond.exprStmtNode.expr.binaryOpNode.operator)
	var stmtNode StmtNode
	if stmtNode.forNode == nil {
		stmtNode.forNode = &ForNode{}
	}
	stmtNode.forNode.init = init
	stmtNode.forNode.cond = cond.exprStmtNode.expr
	if incr != (ExprNode{}) {
		stmtNode.forNode.incr = createExprStmtNode(incr)
	}
	stmtNode.forNode.body = body
	fmt.Println("chenyao*****************for cond is", cond.exprStmtNode.expr.binaryOpNode.operator)
	return stmtNode
}

func createBlockItemFormVar(variable Variable) BlockItem {
	var blockItem BlockItem
	if blockItem.variables == nil {
		blockItem.variables = &DefinedVariable{}
	}
	blockItem.variables.variable = variable
	return blockItem
}

func createBlockItemFormStmt(stmt StmtNode) BlockItem {
	var blockItem BlockItem
	if blockItem.stmts == nil {
		blockItem.stmts = &StmtNode{}
	}
	blockItem.stmts = &stmt
	return blockItem
}

func createDefinedVariable(dectype int, expr ExprNode, type_specifier int) Variable {
	var variable Variable

	variable.entity.isPrivate = false
	variable.entity.name = expr.literalNode.identifierLiteralNode.value
	variable.entity.typeNode.types = type_specifier
	variable.entity.typeNode.dectype = dectype
	return variable
}
