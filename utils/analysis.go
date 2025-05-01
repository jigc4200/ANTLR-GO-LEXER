package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"antlr-go-example/parser"

	"github.com/antlr4-go/antlr/v4"
)

// ANALISIS EN TIEMPO REAL

type Result struct {
	Chars  int
	Lines  int
	Tokens int
}

func RunBatchAnalysis(data []byte, inputText string) {
	var wg sync.WaitGroup
	result := &Result{}

	input := antlr.NewInputStream(inputText)
	lexer := parser.NewProgramaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	wg.Add(3)
	go countChars(data, &wg, result)
	go countLines(data, &wg, result)
	go countTokens(stream, &wg, result)

	wg.Wait()

	fmt.Printf("\n--- Resultados del Conteo ---\n")
	fmt.Printf("Cantidad de caracteres: %d\n", result.Chars)
	fmt.Printf("Cantidad de líneas: %d\n", result.Lines)
	fmt.Printf("Cantidad de tokens: %d\n", result.Tokens)

	PerformAnalysis(inputText)
}

func RunRealtimeAnalysis() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("👨‍💻 Analizador en Tiempo Real (escribe tu código línea por línea, escribe ':exit' para salir)")
	fmt.Println("-----------------------------------------------------------------------------------")

	var inputBuffer strings.Builder
	lineNumber := 1

	for {
		fmt.Printf("%d | ", lineNumber)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\nSaliendo del modo tiempo real (EOF detectado).")
				break
			}
			log.Printf("Error leyendo línea: %v", err)
			continue
		}

		trimmedLine := strings.TrimRight(line, "\r\n")
		if trimmedLine == ":exit" {
			fmt.Println("Saliendo del modo tiempo real.")
			break
		}

		inputBuffer.WriteString(line)
		PerformAnalysis(inputBuffer.String())
		lineNumber++
	}
}
