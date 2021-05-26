grammar vinLang;

program: sequenceStatement;

sequenceStatement: statement (statement)*;

statement: declaration;

//assign: ID ASSIGN expression SEMICOLIN;

declaration: VAR ID COLON ID ('=' expression)* SEMICOLIN;

expression: NUMBER
    | ID
    | expression (MUL|DIV) expression
    | expression (ADD|SUB) expression
    | '(' expression ')';

VAR: 'biến';
COLON: ':';
SEMICOLIN: ';';
ASSIGN: '=';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';

// TODO: them dau
FOR: 'lap';

NUMBER: DIGIT+ ('.' DIGIT+)?;
DIGIT: [0-9];

ID: CHAR+;
CHAR: ([a-zA-Z0-9]|'\u00c0'..'\u1ed1');
//STRING: '"' CHAR* '"';


SPACE
 : [ \t\r\n] -> skip
 ;

 //TODO: update features. More new test files