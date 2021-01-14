grammar symonlang;

// PARSER HERE
expression
    : constant
    | function
    | expression operator expression
    ;

function
    : FUNCTION_NAME BOPEN argumentList BCLOSE
    ;

constant
    : stringLiteral
    | integerLiteral
    | floatLiteral
    | boolLiteral
    ;

operator
    : PLUS
    |  MINUS
    |  DIV
    |  MUL
    |  MOD
    |  LOGICAND
    |  LOGICOR
    |  EQ
    |  NEQ
    |  LTE
    |  LT
    |  GT
    |  GTE
    ;


argumentList
    : ( expression (',' expression )* )?
    ;

stringLiteral
    : STRING_LIT
    ;

integerLiteral
    : MINUS? INT_LIT
    ;

floatLiteral
    : MINUS? REAL_LIT
    ;

boolLiteral
    : TRUE | FALSE
    ;

// LEXER HERE
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;
fragment DEC_DIGIT          : [0-9];
fragment DEC_DIGITS         : DEC_DIGIT+;


PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;
MOD                         : '%' ;
LOGICAND                    : '&&' ;
LOGICOR                     : '||' ;
EQ                          : '==';
NEQ                         : '!=';
LTE                         : '<=';
LT                          : '<';
GT                          : '>';
GTE                         : '>=';

DOT                         : '.' ;

BOPEN                       : '(' ;
BCLOSE                      : ')' ;


STRING_LIT                  : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
INT_LIT                     : '0'
                            | [1-9] DEC_DIGITS?
                            ;
DECIMAL_EXPONENT            : E (PLUS|MINUS)? DEC_DIGITS;
REAL_LIT                    : INT_LIT DOT DEC_DIGITS DECIMAL_EXPONENT?
                            | INT_LIT DECIMAL_EXPONENT
                            | DOT DEC_DIGITS DECIMAL_EXPONENT?
                            ;

TRUE                        : T R U E ;
FALSE                       : F A L S E ;


FUNCTION_NAME               : [a-zA-Z] [a-zA-Z0-9]*;

// IGNORED TOKENS
SPACE                       : [ \t\r\n]+    -> skip;