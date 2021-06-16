grammar vinLang;

program: sequenceStatement;

sequenceStatement: statement (statement)*;

statement: (declaration SEMICOLIN)
    | forStatement
    | (assign SEMICOLIN)
    | LEFT_BRACE sequenceStatement RIGHT_BRACE;

assign: ID ASSIGN expression;

declaration: VAR ID COLON ID (ASSIGN expression)*;

forStatement: FOR LEFT_PARENTHESE declaration SEMICOLIN boolExpression SEMICOLIN assign RIGHT_PARENTHESE statement;

boolExpression: TRUE|FALSE
    | expression E expression
    | expression G expression
    | expression L expression
    | expression LE expression
    | expression GE expression
    | expression NE expression
    | NEG boolExpression
    | LEFT_PARENTHESE boolExpression RIGHT_PARENTHESE
    ;

expression: NUMBER
    | ID
    | expression (MUL|DIV) expression
    | expression (ADD|SUB) expression
    | LEFT_PARENTHESE expression RIGHT_PARENTHESE
    | LEFT_BRACE SequenceStatement RIGHT_BRACE;


VAR: 'biến';

COLON: ':';
SEMICOLIN: ';';
ASSIGN: '=';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/' | '%';
NEG: '!';
LE: '<=';
L: '<';
G: '>';
GE: '>=';
E: '==';
NE: '==';
INCREMENT: '++';
DECREMENT: '--';
LEFT_BRACE: '{';
RIGHT_BRACE: '}';
LEFT_PARENTHESE: '(';
RIGHT_PARENTHESE: ')';

FOR: 'lặp';
CONTINUE: 'tiếp';
BREAK: 'dừng';
RETURN: 'trả';
IF: 'nếu';
ELSE:'thì';
ELIF:'không_thì';
TRUE:'đúng';
FALSE:'sai';

NUMBER: SUB* DIGIT+ ('.' DIGIT+)?;
DIGIT: [0-9];

ID: CHAR+;
CHAR: ([a-zA-Z0-9]|'\u00c0'..'\u1ed1');
//STRING: '"' CHAR* '"';


SPACE
 : [ \t\r\n] -> skip
 ;
