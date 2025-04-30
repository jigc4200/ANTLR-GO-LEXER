# Arquitectura del Proyecto `antlr-go-example`

Este proyecto es un ejemplo básico que demuestra cómo usar ANTLR (ANother Tool for Language Recognition) con Go para crear un parser para un lenguaje simple (una calculadora aritmética).

## Componentes Principales
 
1.  **Definición de la Gramática (`Programa.g4`)**:
    *   Este archivo es el corazón del parser. Define las reglas léxicas (tokens como `INT`) y sintácticas (reglas como `expr` para sumas, restas, multiplicaciones, divisiones y paréntesis) del lenguaje de la calculadora.

2.  **Generación del Parser (Implícito + `parser/` dir)**:
    *   Se asume que la herramienta ANTLR se ejecutó sobre `Calc.g4`.
    *   Esto generó el código Go necesario para el lexer (`parser/calc_lexer.go`) y el parser (`parser/calc_parser.go`), junto con archivos auxiliares (`.tokens`, `.interp`) y una interfaz de listener (`parser/calc_listener.go`) y una implementación base (`parser/calc_base_listener.go`). Estos archivos contienen la lógica para reconocer la gramática definida en `Calc.g4`.

3.  **Aplicación Principal (`main.go`)**:
    *   Este es el punto de entrada del programa.
    *   Importa el paquete `parser` generado y la librería runtime de ANTLR para Go (`github.com/antlr4-go/antlr/v4`).
    *   Define una cadena de entrada (`"3 + 4 * 5"`).
    *   Crea una instancia del lexer y el parser generados.
    *   Configura un `CustomErrorListener` para un manejo de errores más informativo que el predeterminado.
    *   Invoca al parser para analizar la cadena de entrada comenzando por la regla `expr`.
    *   Imprime el árbol de sintaxis abstracto (AST) resultante de la expresión analizada. No *evalúa* la expresión, solo muestra su estructura.

4.  **Dependencias (`go.mod`, `go.sum`)**:
    *   Estos archivos gestionan las dependencias del proyecto, principalmente las librerías runtime de ANTLR necesarias para que el código generado funcione.

5.  go run main.go // EJECUCION
    antlr4 -Dlanguage=Go Programa.g4 // DENTRO DE LA CARPETA PARSER
   
## Resumen

Es una arquitectura estándar para un proyecto ANTLR: defines una gramática, generas el parser, y luego usas ese parser en tu aplicación principal para analizar texto que se ajuste a esa gramática. El foco de este ejemplo está en el *análisis* (parsing) y no en la *evaluación* de la expresión.

import "github.com/antlr4-go/antlr/v4" // usa este funciona mejor 
