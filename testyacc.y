%{
package main
%}
%union{
    exprNode ExprNode
    exprNodeList []ExprNode
}

%token <exprNode>     INT_LITERAL
%token <exprNode>     FLOAT_LITERAL
%token <exprNode>     STRING_LITERAL
%token <exprNode>     CHAR_LITERAL
%token <exprNode>     BOOL_LITERAL
%token <exprNode>     IDENTIFIER
%token INC_OP DEC_OP LEFT_OP RIGHT_OP LE_OP GE_OP EQ_OP NE_OP
       AND_OP OR_OP MUL_ASSIGN DIV_ASSIGN MOD_ASSIGN ADD_ASSIGN
       SUB_ASSIGN LEFT_ASSIGN RIGHT_ASSIGN AND_ASSIGN
       XOR_ASSIGN OR_ASSIGN
// %type   <parameter_list> parameter_list
// %type   <argument_list> argument_list
%type   <exprNode> expression expression_opt
        assignment_expression logical_and_expression logical_or_expression
        equality_expression relational_expression
        additive_expression multiplicative_expression
        unary_expression primary_expression primary_no_new_array
        array_literal array_creation identifier_opt label_opt
%type   <exprNodeList> expression_list 
// %type   <statement> statement
//         if_statement while_statement for_statement foreach_statement
//         return_statement break_statement continue_statement try_statement
//         throw_statement declaration_statement
// %type   <statement_list> statement_list
// %type   <block> block
// %type   <elsif> elsif elsif_list
// %type   <assignment_operator> assignment_operator
// %type   <identifier> identifier_opt label_opt
// %type   <type_specifier> type_specifier
// %type   <basic_type_specifier> basic_type_specifier
// %type   <array_dimension> dimension_expression dimension_expression_list
//         dimension_list

%start translation_unit
%%
translation_unit
        : definition_or_statement
        | translation_unit definition_or_statement
        ;
definition_or_statement
        : function_definition
        | statement
        ;
basic_type_specifier
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
        ;
type_specifier
        : basic_type_specifier
        | '[' ']' type_specifier
        ;
function_definition
        : FUNC IDENTIFIER '(' parameter_list ')' type_specifier block
        | FUNC IDENTIFIER '(' ')' type_specifier block 
        ;
parameter_list
        : IDENTIFIER type_specifier
        | parameter_list ',' IDENTIFIER type_specifier 
        ;
argument_list
        : assignment_expression
        | argument_list ',' assignment_expression
        ;
statement_list
        : statement
        | statement_list statement
        ;
expression
        : assignment_expression
        | expression ',' assignment_expression
        ;
assignment_expression
        : logical_or_expression
        | primary_expression assignment_operator assignment_expression
        ;
assignment_operator
        : '='
        | ADD_ASSIGN
        | SUB_ASSIGN
        | MUL_ASSIGN
        | DIV_ASSIGN
        | MOD_ASSIGN_T
        | LEFT_ASSIGN
        | RIGHT_ASSIGN
        | AND_ASSIGN
        | XOR_ASSIGN
        | OR_ASSIGN
        ;
logical_or_expression
        : logical_and_expression
        | logical_or_expression OR_OP logical_and_expression
        ;
logical_and_expression
        : equality_expression
        | logical_and_expression AND_OP equality_expression
        ;
equality_expression
        : relational_expression
        | equality_expression EQ_OP relational_expression
        | equality_expression NE_OP relational_expression
        ;
relational_expression
        : additive_expression
        | relational_expression '>' additive_expression
        | relational_expression GE_OP additive_expression
        | relational_expression '<' additive_expression
        | relational_expression LE_OP additive_expression
        ;
additive_expression
        : multiplicative_expression
        | additive_expression '+' multiplicative_expression
        | additive_expression '-' multiplicative_expression
        ;
multiplicative_expression
        : unary_expression
        | multiplicative_expression '*' unary_expression
        | multiplicative_expression '/' unary_expression
        | multiplicative_expression '%' unary_expression
        ;
unary_expression
        : primary_expression
        // | '-' unary_expression
        // | '!' unary_expression
        ;
primary_expression
        : primary_no_new_array
        | array_creation
        ;
primary_no_new_array
        : primary_no_new_array '[' argument_list ']'
        | primary_expression '.' IDENTIFIER
        | primary_expression '(' argument_list ')'
        | primary_expression '(' ')'
        | primary_expression INCREMENT
        | primary_expression DECREMENT
        | '(' expression ')'
        | IDENTIFIER
        | INT_LITERAL
        | FLOAT_LITERAL
        | STRING_LITERAL
        | CHAR_LITERAL
        | BOOL_LITERAL
        | array_literal
        ;
expression_list
        : /* empty */
        // {
        //     $$ = nil;
        // }
        | assignment_expression
        | expression_list ',' assignment_expression
        ;
array_literal
        : '{' expression_list '}'
        | '{' expression_list ',' '}'
        ;
array_creation
        : dimension_expression_list basic_type_specifier
        | dimension_expression_list dimension_list basic_type_specifier
        ;
dimension_expression_list
        : dimension_expression
        | dimension_expression_list dimension_expression
        ;
dimension_expression
        : '[' expression ']'
        ;
dimension_list
        : '[' ']'
        | dimension_list '[' ']'
        ;
statement
        : expression ';'
        | if_statement
        | while_statement
        | for_statement
        | foreach_statement
        | return_statement
        | break_statement
        | continue_statement
        | try_statement
        | throw_statement
        | declaration_statement
        ;
if_statement
        : IF '(' expression ')' block
        | IF '(' expression ')' block ELSE block
        | IF '(' expression ')' block elsif_list
        | IF '(' expression ')' block elsif_list ELSE block
        ;
elsif_list
        : elsif
        | elsif_list elsif
        ;
elsif
        : ELSE IF '(' expression ')' block
        ;
label_opt
        : /* empty */
        // {
        //     $$ = ""
        // }
        | IDENTIFIER ':'
        // {
        //     $$ = $1
        // }
        ;
while_statement
        : label_opt WHILE '(' expression ')' block
        ;
for_statement
        : label_opt FOR '(' expression_opt ';' expression_opt ';' expression_opt ')' block
        ;
expression_opt
        : /* empty */
        // {
        //     $$ = nil
        // }
        | expression
        ;
return_statement
        : RETURN expression_opt ';'
        ;
identifier_opt
        : /* empty */
        // {
        //     $$ = ""
        // }
        | IDENTIFIER
        ;
break_statement 
        : BREAK identifier_opt ';'
        ;
continue_statement
        : CONTINUE identifier_opt ';'
        ;
declaration_statement
        : SET IDENTIFIER type_specifier ';'
        | SET IDENTIFIER type_specifier '=' expression ';'
        | LET IDENTIFIER type_specifier ';'
        | LET IDENTIFIER type_specifier '=' expression ';'
        ;
block
        : '{'
        // {
        //     $<block>$ = fate_open_block()
        // }
          statement_list '}'
        // {
        //     $<block>$ = fate_close_block($<block>2, $3)
        // }
        | '{' '}'
        // {
        //     empty_block := fate_open_block()
        //     $<block>$ = fate_close_block(empty_block, nil)
        // }
        ;
%%