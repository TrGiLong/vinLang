grammar vinLang;

program: sequenceStatement;

sequenceStatement: statement (statement)*;

statement: declaration | forStatement;

//assign: ID ASSIGN expression SEMICOLIN;

declaration: VAR ID COLON ID ('=' expression)* SEMICOLIN;

forStatement: FOR '(' declaration ';' expression ';' statement ')';

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
