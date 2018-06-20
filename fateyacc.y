%{
package main
%}

%union{
    exprNode ExprNode
	exprNode_list []ExprNode
	unary_operator byte
	assignment_operator string
	stmtNode StmtNode
}

%token <exprNode>     INT_LITERAL
%token <exprNode>     FLOAT_LITERAL
%token <exprNode>     STRING_LITERAL
%token <exprNode>     CHAR_LITERAL
%token <exprNode>     BOOL_LITERAL
%token <exprNode>     IDENTIFIER

%token INC_OP DEC_OP LEFT_OP RIGHT_OP LE_OP GE_OP EQ_OP NE_OP
%token AND_OP OR_OP MUL_ASSIGN DIV_ASSIGN MOD_ASSIGN ADD_ASSIGN
%token SUB_ASSIGN LEFT_ASSIGN RIGHT_ASSIGN AND_ASSIGN
%token XOR_ASSIGN OR_ASSIGN

%token INT8, INT16, INT32, INT64, UINT8, UINT16, UINT32, UINT64, FLOAT32, FLOAT64
%token BOOL
%token CHAR STRING
%token TYPE_CONST TYPE_STRUCT
%token STRUCT CONST EXTERN

%token LET SET FUNC TYPE CASE DEFAULT IF ELSE SWITCH FOR GOTO CONTINUE BREAK RETURN

%type <exprNode> primary_expression expression assignment_expression constant_expression conditional_expression 
	  logical_or_expression logical_and_expression inclusive_or_expression exclusive_or_expression and_expression equality_expression
	  relational_expression shift_expression additive_expression cast_expression multiplicative_expression unary_expression postfix_expression
%type <unary_operator> unary_operator
%type <exprNode_list> argument_expression_list
%type <assignment_operator> assignment_operator
%type <stmtNode> statement labeled_statement compound_statement expression_statement selection_statement iteration_statement jump_statement
%start translation_unit

%%

type_specifier
    : INT8
	| INT16
	| INT32
	| INT64
	| UINT8
	| UINT16
	| UINT32
	| UINT64
    | FLOAT32
    | FLOAT64
	| BOOL
    | CHAR
    | STRING
    | TYPE_CONST
    | TYPE_STRUCT
    ; /* 声明过的IDENTIFIER返回TYPE_STRUCT或者TYPE_CONST，没有声明过则返回IDENTIFIER */

translation_unit
	: external_declaration
	| translation_unit external_declaration
    ;

external_declaration
	: declaration
    | function_definition
	| variable_definition
    ;

declaration
    : type_declaration
    ;

type_declaration
	: const_declaration
	| struct_declaration
    ;

function_declaration
    : EXTERN FUNC IDENTIFIER '(' parameter_list ')' ';'
    | EXTERN FUNC IDENTIFIER '(' parameter_list ')' type_specifier ';'
    ;

function_definition
    : FUNC IDENTIFIER '(' parameter_list ')' compound_statement
    | FUNC IDENTIFIER '(' parameter_list ')' type_specifier compound_statement
    ;

variable_definition
    : LET declarator type_specifier ';'
    | SET declarator type_specifier ';'
    ;

parameter_list
    :
    | declarator type_specifier
    | declarator type_specifier ',' parameter_list
    ;

statement
	: labeled_statement
	| compound_statement
	| expression_statement
	| selection_statement
	| iteration_statement
	| jump_statement
	;

compound_statement
	: '{' '}'
	{
		$$ = aa()
	}
	| '{' block_item_list '}'
	{
		$$ = aa()
	}
	;

block_item_list
	: block_item
	| block_item_list block_item
	;

block_item
	: type_declaration
    | variable_definition
	| statement
	;

labeled_statement
	: IDENTIFIER ':' statement
	{
		$$ = aa()
	}
	| CASE constant_expression ':' statement
	{
		$$ = aa()
	}
	| DEFAULT ':' statement
	{
		$$ = aa()
	}
	;

expression_statement
	: ';'
	{
		$$ = StmtNode{}
	}
	| expression ';'
	{
		$$ = createExprStmtNode($1)
	}
	;

selection_statement
	: IF '(' expression ')' statement
	{
		$$ = createIfNode($3,$5,(StmtNode{}))
	}
	| IF '(' expression ')' statement ELSE statement
	{
		$$ = createIfNode($3,$5,$7)
	}
	| SWITCH '(' expression ')' statement
	{
		$$ = aa()
	}
	;

iteration_statement
	: FOR '(' expression_statement expression_statement ')' statement
	{
		$$ = createForNode($3,$4,(ExprNode{}),$6)
	}
	| FOR '(' expression_statement expression_statement expression ')' statement
	{
		$$ = createForNode($3,$4,$5,$7)
	}
	;

jump_statement
	: GOTO IDENTIFIER ';'
	{
		$$ = aa()
	}
	| CONTINUE ';'
	{
		$$ = aa()
	}
	| BREAK ';'
	{
		$$ = aa()
	}
	| RETURN ';'
	{
		$$ = aa()
	}
	| RETURN expression ';'
	{
		$$ = aa()
	}
	;

constant_expression
	: conditional_expression
	;

conditional_expression
	: logical_or_expression
	| logical_or_expression '?' expression ':' conditional_expression
	{
		$$ = createCondExprNode($1,$3,$5)
	}
	;

logical_or_expression
	: logical_and_expression
	| logical_or_expression OR_OP logical_and_expression
	{
		$$ = ceateBinaryOpNode($1,"||",$3)
	}
	;

logical_and_expression
	: inclusive_or_expression
	| logical_and_expression AND_OP inclusive_or_expression
	{
		$$ = ceateBinaryOpNode($1,"&&",$3)
	}
	;

inclusive_or_expression
	: exclusive_or_expression
	| inclusive_or_expression '|' exclusive_or_expression
	{
		$$ = ceateBinaryOpNode($1,"|",$3)
	}
	;

exclusive_or_expression
	: and_expression
	| exclusive_or_expression '^' and_expression
	{
		$$ = ceateBinaryOpNode($1,"^",$3)
	}
	;

and_expression
	: equality_expression
	| and_expression '&' equality_expression
	{
		$$ = ceateBinaryOpNode($1,"&",$3)
	}
	;

equality_expression
	: relational_expression
	| equality_expression EQ_OP relational_expression
	{
		$$ = ceateBinaryOpNode($1,"==",$3)
	}
	| equality_expression NE_OP relational_expression
	{
		$$ = ceateBinaryOpNode($1,"!=",$3)
	}
	;

relational_expression
	: shift_expression
	| relational_expression '<' shift_expression
	{
		$$ = ceateBinaryOpNode($1,"<",$3)
	}
	| relational_expression '>' shift_expression
	{
		$$ = ceateBinaryOpNode($1,">",$3)
	}
	| relational_expression LE_OP shift_expression
	{
		$$ = ceateBinaryOpNode($1,"<=",$3)
	}
	| relational_expression GE_OP shift_expression
	{
		$$ = ceateBinaryOpNode($1,">=",$3)
	}
	;

shift_expression
	: additive_expression
	| shift_expression LEFT_OP additive_expression
	{
		$$ = ceateBinaryOpNode($1,"<<",$3)
	}
	| shift_expression RIGHT_OP additive_expression
	{
		$$ = ceateBinaryOpNode($1,">>",$3)
	}
	;

additive_expression
	: multiplicative_expression
	| additive_expression '+' multiplicative_expression
	{
		$$ = ceateBinaryOpNode($1,"+",$3)
	}
	| additive_expression '-' multiplicative_expression
	{
		$$ = ceateBinaryOpNode($1,"-",$3)
	}
	;

multiplicative_expression
	: cast_expression
	| multiplicative_expression '*' cast_expression
	{
		$$ = ceateBinaryOpNode($1,"*",$3)
	}
	| multiplicative_expression '/' cast_expression
	{
		$$ = ceateBinaryOpNode($1,"/",$3)
	}
	| multiplicative_expression '%' cast_expression
	{
		$$ = ceateBinaryOpNode($1,"%",$3)
	}
	;

cast_expression
	: unary_expression
	;

unary_operator
	: '+'
	{
		$$='+'
	}
	| '-'
	{
		$$='-'
	}
	| '!'
	{
		$$='!'
	}
	;

unary_expression
	: postfix_expression
	| INC_OP unary_expression
	{
		$$=createPrefixOpNode(INC_OP,$2)
	}
	| DEC_OP unary_expression
	{
		$$=createPrefixOpNode(DEC_OP,$2)
	}
	| unary_operator cast_expression
	{
		$$=createUnaryOpNode($1,$2)
	}
	;

postfix_expression
	: primary_expression
	| postfix_expression '[' expression ']'
	{
		$$=createArefNode($1,$3)
	}
    | postfix_expression '(' argument_expression_list ')'
	{
		$$=createFuncallNode($1,$3)
	}
	| postfix_expression '.' IDENTIFIER
	{
		$$=createMemberNode($1,$3)
	}
	| postfix_expression INC_OP
	{
		$$=createSuffixOpNode(INC_OP,$1)
	}
	| postfix_expression DEC_OP
	{
		$$=createSuffixOpNode(INC_OP,$1)
	}
	;

primary_expression
	: IDENTIFIER
	| INT_LITERAL      /* 整形常量 */
    | BOOL_LITERAL     /* bool常量 */
    | CHAR_LITERAL     /* char常量 */
    | FLOAT_LITERAL    /* 浮点常量 */
	| STRING_LITERAL    /* 字符串常量 */
	| '(' expression ')'
	{
		$$=$2
	}
	;

expression
	: assignment_expression
	| expression ',' assignment_expression
	;

assignment_expression
	: conditional_expression
	| unary_expression assignment_operator assignment_expression
	{
		$$ = createAssignNode($1,$2,$3)
	}
	;

assignment_operator
	: '='
	{
		$$ = "="
	}
	| MUL_ASSIGN
	{
		$$ = "*="
	}
	| DIV_ASSIGN
	{
		$$ = "/="
	}
	| MOD_ASSIGN
	{
		$$ = "%="
	}
	| ADD_ASSIGN
	{
		$$ = "+="
	}
	| SUB_ASSIGN
	{
		$$ = "-="
	}
	| LEFT_ASSIGN
	{
		$$ = "<<="
	}
	| RIGHT_ASSIGN
	{
		$$ = ">>="
	}
	| AND_ASSIGN
	{
		$$ = "&="
	}
	| XOR_ASSIGN
	{
		$$ = "^="
	}
	| OR_ASSIGN
	{
		$$ = "|="
	}
	;

struct_declaration
    : TYPE IDENTIFIER STRUCT '{' struct_declaration_list '}' ';' /*{ft.newWord(TYPE_STRUCT, $2)}*/
    ;

struct_declaration_list
	: declarator type_specifier
	| struct_declaration_list ',' declarator type_specifier
	;

const_declaration /* 枚举类型 */
    : CONST '(' consterator_list ')'
    | CONST IDENTIFIER '(' consterator_list ')' /*{ft.newWord(TYPE_CONST, $2)}*/
    ;

consterator_list
	: consterator
	| consterator_list ',' consterator
	;

consterator
	: IDENTIFIER
	| IDENTIFIER '=' constant_expression
	;

declarator /* 支持多重数组和map */
	: IDENTIFIER
	| declarator '[' ']' /* 数组 */
    | declarator '<' type_specifier '>' /* map */
	;

argument_expression_list
    : assignment_expression
	{
		$$=bb()
	}
    | argument_expression_list ',' assignment_expression
	{
		$$=bb()
	}
    ;

%%