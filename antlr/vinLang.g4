grammar vinLang;

program: sequenceStatement;

sequenceStatement: statement (statement)*;

statement: (declaration SEMICOLIN) | forStatement;

//assign: ID ASSIGN expression SEMICOLIN;

declaration: VAR ID COLON ID ('=' expression)*;

forStatement: FOR '(' declaration ';' boolExpression ';' statement ')';

boolExpression: TRUE|FALSE
    | expression '<' expression
    | expression '>' expression
    | expression '==' expression
    | expression '<=' expression
    | expression '>=' expression
    | expression '!=' expression
    | 'không' Bool
    | Bool 'var' Bool
    | '(' Bool ')'
    ;

expression: NUMBER
    | ID
    | expression (MUL|DIV) expression
    | expression (ADD|SUB) expression
    | '(' expression ')'
    | '{' SequenceStatement '}';

VAR: 'biến';

COLON: ':';
SEMICOLIN: ';';
ASSIGN: '=';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/' | '%';

FOR: 'lặp';
CONTINUE: 'tiếp';
BREAK: 'dừng';
RETURN: 'trả';
IF: 'nếu';
ELSE:'thì';
ELIF:'không_thì';
TRUE:'đúng';
FALSE:'sai';

NUMBER: DIGIT+ ('.' DIGIT+)?;
DIGIT: [0-9];

ID: CHAR+;
CHAR: ([a-zA-Z0-9]|'\u00c0'..'\u1ed1');
//STRING: '"' CHAR* '"';


SPACE
 : [ \t\r\n] -> skip
 ;
