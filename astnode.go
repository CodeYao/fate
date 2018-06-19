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
	variable    Variable
	initializer ExprNode //表达式节点
	ir          Expr
	sequence    int64
	symbol      Symbol //符号表
}

type Entity struct {
	name      string
	isPrivate bool
	typeNode  TypeNode
	nRefered  int64
	//menref    MemoryReference//汇编时用，暂不使用
	//address   Operand//汇编时用，暂不使用
}

type Constant struct {
	entity       Entity
	typeConstant TypeNode
	name         string
	value        ExprNode
}
type Function struct {
	entity Entity
	//callingSymbol Symbol//暂不使用
	//label Label//暂不使用
}
type Variable struct {
	entity Entity
}

//表达式语句节点
type ExprNode struct {
	abstractAssignNode *AbstractAssignNode //赋值
	addressNode        *AddressNode        //地址表达式，可能有问题
	binaryOpNode       *BinaryOpNode       //二元运算表达式 x+y...
	castNode           *CastNode           //类型转换
	condExprNode       *CondExprNode       //条件表达式(a?b:c)
	funcCallNode       *FuncCallNode       //函数调用表达式
	lhNode             *LHSNode            //能够成为赋值的左值节点
	literalNode        *LiteralNode        //字面量
	sizeofExprNode     *SizeofExprNode     //计算表达式的sizeof的表达式
	sizeofTYpeNode     *SizeofTypeNode     //计算类型的sizeof的表达式
	unaryOpNode        *UnaryOpNode        //一元运算表达式
}

type ExprNodeList struct {
	exprNode ExprNode
	next     *ExprNodeList
}

type UndefinedVariable struct {
	//未定义变量，暂不处理
}

type DefinedFunction struct {
	function Function
	params   Params
	body     BlockNode
	//scope    LocalScope
	//ir       []Stmt
}
type UndefinedFunction struct {
	function Function
	params   Params
}

type Params struct {
	location         Location
	paramDescriptors []DefinedVariable
	vararg           bool
}

//***************************

type StmtNode struct {
	breakNode    *BreakNode    //break语句
	caseNode     *CaseNode     //case标签
	continueNode *ContinueNode //contiune语句
	exprStmtNode *ExprStmtNode //单独构成语句的表达式
	blocknode    *BlockNode    //程序块{...}
	forNode      *ForNode      //for语句
	ifNode       *IfNode       //if语句
	labelNode    *LabelNode    //goto标签
	gotoNode     *GotoNode     //goto语句
	returnNode   *ReturnNode   //return语句
	switchNode   *SwitchNode   //switch语句
	whileNode    *WhileNode    //while语句
	doWhileNode  *DoWhileNode  //do...while语句
}
type BlockNode struct {
	location  Location
	variables []DefinedVariable
	stmts     []StmtNode
	//scope     LocalScope //暂不使用
}

type BreakNode struct {
	location Location
}
type CaseNode struct {
	location Location
	//label Label//暂不使用
	values []ExprNode
	body   BlockNode
}
type ContinueNode struct {
	location Location
}
type ExprStmtNode struct {
	location Location
	expr     ExprNode
}
type ForNode struct {
	location Location
	init     StmtNode
	cond     ExprNode
	incr     StmtNode
	body     StmtNode
}
type GotoNode struct {
	location Location
	target   string
}

type LabelNode struct {
	location Location
	name     string
	stmt     StmtNode
}

type IfNode struct {
	location Location
	cond     ExprNode
	thenBody StmtNode
	elseBody StmtNode
}

type ReturnNode struct {
	location Location
	expr     ExprNode
}
type SwitchNode struct {
	location Location
	cond     ExprNode
	cases    []CaseNode
}
type WhileNode struct {
	location Location
	body     StmtNode
	cond     ExprNode
}
type DoWhileNode struct {
	location Location
	body     StmtNode
	cond     ExprNode
}

//************************

type AbstractAssignNode struct {
	assignNode   *AssignNode   //赋值
	opAssignNode *OpAssignNode //带运算的赋值+=...
}
type AssignNode struct {
	lhs ExprNode
	rhs ExprNode
}

type OpAssignNode struct {
	operator string
	lhs      ExprNode
	rhs      ExprNode
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

type StructNode struct {
	compositeTypeDefinition CompositeTypeDefinition
}
type CompositeTypeDefinition struct {
	member         []Slot
	typeDefinition TypeDefinition
}

//结构体成员列表节点
type Slot struct {
	typeNode TypeNode
	name     string
	offset   int64
}

//类型节点，可能会有问题
type TypeNode struct {
	typeRef BasicType
	types   BasicType
}
type TypedefNode struct {
	real           TypeNode
	typeDefinition TypeDefinition
}
type TypeDefinition struct {
	name     string
	location Location
	typeNode TypeNode
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
	types        BasicType
	variableNode VariableNode
	//暂时不处理
}
type VariableNode struct {
	location Location
	name     string
	entity   Entity
}
type LiteralNode struct {
	integerLiteralNode    *IntegerLiteralNode
	floatLiteralNode      *FloatLiteralNode
	charLiteralNode       *CharLiteralNode
	boolLiteralNode       *BoolLiteralNode
	stringLiteralNode     *StringLiteralNode
	identifierLiteralNode *IdentifierLiteralNode
}
type IntegerLiteralNode struct {
	location Location
	typeNode TypeNode
	value    uint64
}
type FloatLiteralNode struct {
	location Location
	typeNode TypeNode
	value    float64
}
type CharLiteralNode struct {
	location Location
	typeNode TypeNode
	value    byte
}
type BoolLiteralNode struct {
	location Location
	typeNode TypeNode
	value    bool
}
type IdentifierLiteralNode struct {
	location Location
	typeNode TypeNode
	value    string
	//entry    ConsrantEntry
}
type StringLiteralNode struct {
	location Location
	typeNode TypeNode
	value    string
	entry    ConsrantEntry
}
type ConsrantEntry struct {
	value  string
	symbol Symbol
	//memref MemoryReference//汇编时用，暂不使用
	//address ImmediateValue//立即数
}

type Location struct {
	location string
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
	amount   int64
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
