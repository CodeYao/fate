package main

//根节点
type AST struct {
	//source Location
	decarations Declarations //声明
	// scope         ToplevelScope //顶层范围
	// constantTable ConstantTable //常量表
}

// type Location struct {
// 	sourceName string
// 	token      Token
// }
type Declarations struct {
	defvars    []DefinedVariable   //已定义的变量
	vardecls   []UndefinedVariable //未定义变量
	defuns     []DefinedFunction   //定义的函数
	funcdecls  []UndefinedFunction //未定义的函数
	constants  []Constant          //常量
	defstructs []StructNode        //结构体
	typedefs   []TypedefNode       //类型
}

type DefinedVariable struct {
	initializer ExprNode //表达式节点
	ir          Expr
	sequence    int64
	symbol      Symbol //符号表
}

//表达书语句节点
type ExprNode struct {
	abstractAssignNode AbstractAssignNode //赋值
	addressNode        AddressNode        //地址表达式，可能有问题
	binaryOpNode       BinaryOpNode       //二元运算表达式 x+y...
	castNode           CastNode           //类型转换
	condExprNode       CondExprNode       //条件表达式(a?b:c)
	funcCallNode       FuncCallNode       //函数调用表达式
	lhNode             LHSNode            //能够成为赋值的左值节点
	literalNode        LiteralNode        //字面量
	sizeofExprNode     SizeofExprNode     //计算表达式的sizeof的表达式
	sizeofTYpeNode     SizeofTypeNode     //计算类型的sizeof的表达式
	unaryOpNode        UnaryOpNode        //一元运算表达式
}

type UndefinedVariable struct {
	//未定义变量，暂不处理
}

type DefinedFunction struct {
	params Params
	body   BlockNode
	scope  LocalScope
	ir     []Stmt
}

type Params struct {
}
type AbstractAssignNode struct {
	lhs ExprNode
	rhs ExprNode
}
type AddressNode struct {
	expr ExprNode
	//可能有问题
	types BasicType
}
type BinaryOpNode struct {
	operator string
	left     ExprNode
	right    ExprNode
	types    BasicType
}
type CastNode struct {
	typeNode TypeNode
	expr     ExprNode
}

//类型节点，可能会有问题
type TypeNode struct {
	typeRef BasicType
	types   BasicType
}
type CondExprNode struct {
	cond     ExprNode
	thenExpr ExprNode
	elseExpr ExprNode
}
type FuncCallNode struct {
	expr ExprNode
	args []ExprNode
}
type LHSNode struct {
	//暂时不处理
}
type LiteralNode struct {
	location Location
	typeNode TypeNode
}
type Location struct {
	Location string
	token    Token
}
type SizeofExprNode struct {
	expr     ExprNode
	typeNode TypeNode
}
type SizeofTypeNode struct {
	operand  TypeNode
	typeNode TypeNode
}
type UnaryOpNode struct {
	operator string
	expr     ExprNode
	opType   BasicType
}
type Expr struct {
	types BasicType
	//中间代码阶段，暂未设计，可能有问题
}
type Symbol struct {
	//字符表暂不处理
}
