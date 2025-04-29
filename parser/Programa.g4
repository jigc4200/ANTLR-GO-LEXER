grammar Programa;

// ---------------------------
// PARSER
// ---------------------------

programa 
    : 'PROGRAMA' IDENTIFICADOR_VARIABLE ';' declaraciones bloque 'FIN' ';' ;

declaraciones 
    : 'VARIABLES' lista_variables ';' ;
     

lista_variables 
    : IDENTIFICADOR_VARIABLE (',' IDENTIFICADOR_VARIABLE)* ;
    
bloque 
    : 'INICIO' instrucciones ;

instrucciones 
    : instruccion* ;

instruccion 
    : asignacion ';'
    | 'LEER' IDENTIFICADOR_VARIABLE ';'
    | 'ESCRIBIR' expresion ';'
    | condicional ;

asignacion 
    : IDENTIFICADOR_VARIABLE asignacionvar expresion ;

condicional 
    : 'SI' expresion comparador expresion 'ENTONCES' instrucciones ;

expresion 
    : termino (operador expresion)? ;

termino 
    : NUMERO_ENTERO 
    | IDENTIFICADOR_VARIABLE
    | '(' expresion ')' ;

asignacionvar 
    : ASIGNACION
    | IGUALDAD ;

comparador 
    : '=' 
    | '<' 
    | '>'
    | '<='
    | '>='
    | '!=' ;

operador 
    : '+' 
    | '-' 
    | '*' 
    | '/' ;

// ---------------------------
// LÉXICO
// ---------------------------
Reservada_PROGRAMA : 'PROGRAMA' ;
Reservada_FIN      : 'FIN' ;
Reservada_INICIO   : 'INICIO' ;
Reservada_SI       : 'SI' ;
Reservada_ENTONCES : 'ENTONCES' ;
Reservada_LEER     : 'LEER' ;
Reservada_ESCRIBIR : 'ESCRIBIR' ;
Reservada_VARIABLES : 'VARIABLES' ;

ASIGNACION     : '<--' ;
IGUALDAD       : '=' ;
PUNTOYCOMA     : ';' ;
COMA           : ',' ;
PUNTO          : '.' ;
DOS_PUNTOS     : ':' ;
CORCHETE_IZQ   : '[' ;
CORCHETE_DER   : ']' ;

CADENA         : '"' ( ~["] | '""' )* '"' ;
CADENA_CARACTER : '\'' ( ~['] | '\'\'')* '\'' ;

BOOLEANO       : 'VERDADERO' | 'FALSO' ;
NUMERO_ENTERO  : [0-9]+ ;
IDENTIFICADOR_VARIABLE : [a-zA-Z_] [a-zA-Z0-9_]* ;

// Operadores y comparadores
OPERADOR_SUMA              : '+' ;
OPERADOR_RESTA             : '-' ;
OPERADOR_MULTIPLICACION    : '*' ;
OPERADOR_DIVISION          : '/' ;
OPERADOR_PARENTESIS_IZQ      : '(' ;
OPERADOR_PARENTESIS_DER      : ')' ;
COMPARADOR_IGUAL           : '==' ;
COMPARADOR_MAYOR           : '>' ;
COMPARADOR_MENOR           : '<' ;
COMPARADOR_MAYOR_IGUAL     : '>=' ;
COMPARADOR_MENOR_IGUAL     : '<=' ;
COMPARADOR_DIFERENTE       : '!=' ;

OPERADOR_MODULO            : '%' ;
OPERADOR_POTENCIA          : '^' ;
OPERADOR_INCREMENTO        : '++' ;
OPERADOR_DECREMENTO        : '--' ;

// Ignorar espacios, tabs, saltos de línea
WS             : [ \t\r\n]+ -> skip ;

// Ignorar comentarios con #
COMMENT        : '#' ~[\r\n]* ;  // mantener los comentarios de una sola línea
COMMENT_MULTILINE : '/*' .*? '*/' ; // mantener los comentarios multilínea
