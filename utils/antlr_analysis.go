package utils

import (
	"fmt"
	"strings"

	"antlr-go-example/parser"

	"github.com/antlr4-go/antlr/v4"
)

// ANALISIS USANDO ANTLR
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	inputText     string
	lastErrorLine int
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if l.lastErrorLine == line {
		return
	}
	l.lastErrorLine = line
	fmt.Printf("⚠️ Error de sintaxis en línea %d, columna %d: %s\n", line, column, msg)
	lines := strings.Split(l.inputText, "\n")
	if line-1 < len(lines) {
		fmt.Println("Código:")
		fmt.Printf("%d | %s\n", line, lines[line-1])
		fmt.Printf("    %s^\n", strings.Repeat(" ", column))
	}
}

func PerformAnalysis(inputText string) {
	if inputText == "" {
		fmt.Println("(Entrada vacía, saltando análisis ANTLR)")
		return
	}

	input := antlr.NewInputStream(inputText)
	lexer := parser.NewProgramaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	fmt.Printf("\n--- Listado Detallado de Tokens ---\n")
	stream.Fill()
	tokens := stream.GetAllTokens()
	fmt.Printf("Cantidad de tokens: %d\n", len(tokens))

	for i, token := range tokens {
		tokenTypeName := "UNKNOWN"
		if token.GetTokenType() >= 0 && token.GetTokenType() < len(lexer.SymbolicNames) {
			tokenTypeName = lexer.SymbolicNames[token.GetTokenType()]
		} else if token.GetTokenType() == antlr.TokenEOF {
			tokenTypeName = "EOF"
		}
		fmt.Printf(" %d.Token: %q -------> Type: %s\n", i, token.GetText(), tokenTypeName)
	}

	p := parser.NewProgramaParser(stream)
	p.RemoveErrorListeners()
	errorListener := &CustomErrorListener{inputText: inputText}
	p.AddErrorListener(errorListener)

	fmt.Println("\n--- Iniciando Parseo ---")
	errorListener.lastErrorLine = -1
	tree := p.Programa()

	fmt.Println("\n--- Árbol de Parseo ---")
	if tree != nil {
		fmt.Println(tree.ToStringTree(nil, p))
	} else {
		fmt.Println("No se pudo generar el árbol de parseo debido a errores sintácticos.")
	}
}
