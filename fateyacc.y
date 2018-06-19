%{
package main
%}

%union{
    exprNode ExprNode
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

%type <exprNode> primary_expression expression assignment_expression expression_statement constant_expression conditional_expression 
	  logical_or_expression logical_and_expression inclusive_or_expression exclusive_or_expression and_expression equality_expression
	  relational_expression shift_expression additive_expression cast_expression multiplicative_expression unary_expression postfix_expression

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
	| '{' block_item_list '}'
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
	| CASE constant_expression ':' statement
	| DEFAULT ':' statement
	;

expression_statement
	: ';'
	{
		$$= aa()
	}
	| expression ';'
	{
		$$=$1
	}
	;

selection_statement
	: IF '(' expression ')' statement
	| IF '(' expression ')' statement ELSE statement
	| SWITCH '(' expression ')' statement
	;

iteration_statement
	: FOR '(' expression_statement expression_statement ')' statement
	| FOR '(' expression_statement expression_statement expression ')' statement
	;

jump_statement
	: GOTO IDENTIFIER ';'
	| CONTINUE ';'
	| BREAK ';'
	| RETURN ';'
	| RETURN expression ';'
	;

constant_expression
	: conditional_expression
	;

conditional_expression
	: logical_or_expression
	| logical_or_expression '?' expression ':' conditional_expression
	;

logical_or_expression
	: logical_and_expression
	| logical_or_expression OR_OP logical_and_expression
	;

logical_and_expression
	: inclusive_or_expression
	| logical_and_expression AND_OP inclusive_or_expression
	;

inclusive_or_expression
	: exclusive_or_expression
	| inclusive_or_expression '|' exclusive_or_expression
	;

exclusive_or_expression
	: and_expression
	| exclusive_or_expression '^' and_expression
	;

and_expression
	: equality_expression
	| and_expression '&' equality_expression
	;

equality_expression
	: relational_expression
	| equality_expression EQ_OP relational_expression
	| equality_expression NE_OP relational_expression
	;

relational_expression
	: shift_expression
	| relational_expression '<' shift_expression
	| relational_expression '>' shift_expression
	| relational_expression LE_OP shift_expression
	| relational_expression GE_OP shift_expression
	;

shift_expression
	: additive_expression
	| shift_expression LEFT_OP additive_expression
	| shift_expression RIGHT_OP additive_expression
	;

additive_expression
	: multiplicative_expression
	| additive_expression '+' multiplicative_expression
	| additive_expression '-' multiplicative_expression
	;

multiplicative_expression
	: cast_expression
	| multiplicative_expression '*' cast_expression
	| multiplicative_expression '/' cast_expression
	| multiplicative_expression '%' cast_expression
	;

cast_expression
	: unary_expression
	;

unary_operator
	: '+'
	| '-'
	| '!'
	;

unary_expression
	: postfix_expression
	| INC_OP unary_expression
	{
		$$=aa()
	}
	| DEC_OP unary_expression
	{
		$$=aa()
	}
	| unary_operator cast_expression
	{
		$$=aa()
	}
	;

postfix_expression
	: primary_expression
	| postfix_expression '[' expression ']'
    | postfix_expression '(' argument_expression_list ')'
	| postfix_expression '.' IDENTIFIER
	| postfix_expression INC_OP
	| postfix_expression DEC_OP
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
	;

assignment_operator
	: '='
	| MUL_ASSIGN
	| DIV_ASSIGN
	| MOD_ASSIGN
	| ADD_ASSIGN
	| SUB_ASSIGN
	| LEFT_ASSIGN
	| RIGHT_ASSIGN
	| AND_ASSIGN
	| XOR_ASSIGN
	| OR_ASSIGN
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
    | argument_expression_list ',' assignment_expression
    ;

%%