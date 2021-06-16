grammar vinLang;

program: sequenceStatement;

sequenceStatement: statement (statement)*;

statement: declaration SEMICOLIN
    | forStatement
    | ifStatement
    | functionCall SEMICOLIN
    | functionDeclaration
    | (assign SEMICOLIN)
    | ID SEMICOLIN
    | RETURN expression SEMICOLIN
    | CONTINUE SEMICOLIN
    | BREAK SEMICOLIN
    | LEFT_BRACE sequenceStatement RIGHT_BRACE;

assign: ID ASSIGN expression;

declaration: VAR ID COLON ID (ASSIGN expression)*;

forStatement: FOR LEFT_PARENTHESE declaration SEMICOLIN boolExpression SEMICOLIN assign RIGHT_PARENTHESE statement;

functionDeclaration: FUNCTION ID '(' functionArgs ')'  ':' ID '{' sequenceStatement '}';

functionArgs: (ID ':' ID (',' ID ':' ID)*)?;

functionCall: ID '(' (expression (',' expression)*)? ')';

ifStatement: IF '(' boolExpression ')' '{' sequenceStatement '}' (ELIF '(' boolExpression ')' '{' sequenceStatement '}')* (ELSE '{' sequenceStatement '}')?;

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
    | functionCall
    | ID
    | expression (MUL|DIV) expression
    | expression (ADD|SUB) expression
    | LEFT_PARENTHESE expression RIGHT_PARENTHESE
    | LEFT_BRACE expression RIGHT_BRACE;

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
NE: '!=';
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
FUNCTION: 'hàm';

NUMBER: '-'* DIGIT+ ('.' DIGIT+)?;
DIGIT: [0-9];

ID: CHAR+;
CHAR: ([0-9]
|'\u0041'..'\u005A'
|'\u0061'..'\u007A'
|'\u00C0'..'\u00C3'
|'\u00C8'..'\u00CA'
|'\u00CC'..'\u00CD'
|'\u00D2'..'\u00D5'
|'\u00D9'..'\u00DA'
|'\u00DD'
|'\u00E0'..'\u00E3'
|'\u00E8'..'\u00EA'
|'\u00EC'..'\u00ED'
|'\u00F2'..'\u00F5'
|'\u00F9'..'\u00FA'
|'\u00FD'
|'\u0102'..'\u0103'
|'\u0110'..'\u0111'
|'\u0128'..'\u0129'
|'\u0168'..'\u0169'
|'\u01A0'..'\u01A1'
|'\u01AF'..'\u01B0'
|'\u1EA0'..'\u1EF9'
);
//STRING: '"' CHAR* '"';


SPACE
 : [ \t\r\n] -> skip
 ;
