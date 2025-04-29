package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"antlr-go-example/parser" // Importa el paquete generado por ANTLR para el análisis

	"github.com/antlr4-go/antlr/v4"
)

//go:embed files/file.txt
var fileTxt embed.FS // Embebe un archivo de texto dentro del binario

// Estructura para almacenar los resultados del conteo
type Result struct {
	Chars  int // Cantidad de caracteres
	Lines  int // Cantidad de líneas
	Tokens int // Cantidad de tokens
}

// Función para contar caracteres en los datos
func countChars(data []byte, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()       // Marca la goroutine como terminada
	res.Chars = len(data) // Cuenta los caracteres
}

// Función para contar líneas en los datos
func countLines(data []byte, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()
	lines := 0
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		lines++ // Incrementa el contador por cada línea
	}
	// Si los datos no están vacíos y no se detectaron líneas, cuenta una línea
	if len(data) > 0 && lines == 0 {
		lines = 1
	}
	res.Lines = lines
}

// Función para contar tokens usando ANTLR
func countTokens(stream *antlr.CommonTokenStream, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()
	stream.Fill()                           // Llena el stream con todos los tokens
	res.Tokens = len(stream.GetAllTokens()) // Cuenta los tokens
}

// Listener personalizado para manejar errores de sintaxis
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	inputText     string // Almacena el texto de entrada para resaltar errores
	lastErrorLine int    // Evita múltiples errores en la misma línea
}

// Método que se ejecuta cuando ocurre un error de sintaxis
func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Evita reportar múltiples errores en la misma línea
	if l.lastErrorLine == line {
		return
	}
	l.lastErrorLine = line // Marca la línea como reportada

	// Imprime el mensaje de error
	fmt.Printf("⚠️ Error de sintaxis en línea %d, columna %d: %s\n", line, column, msg)
	lines := strings.Split(l.inputText, "\n")
	if line-1 < len(lines) {
		// Muestra la línea con el error
		fmt.Println("Código:")
		fmt.Printf("%d | %s\n", line, lines[line-1])
		// Subraya el error
		fmt.Printf("    %s^\n", strings.Repeat(" ", column))
	}
}

// Función que realiza el análisis con ANTLR
func performAnalysis(inputText string) {
	if inputText == "" {
		fmt.Println("(Entrada vacía, saltando análisis ANTLR)")
		return
	}

	// Crea el lexer y el stream de tokens
	input := antlr.NewInputStream(inputText)
	lexer := parser.NewProgramaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// --- Listado de tokens ---
	fmt.Printf("\n--- Listado Detallado de Tokens ---\n")
	stream.Fill()
	tokens := stream.GetAllTokens()
	fmt.Printf("Cantidad de tokens:---> %d\n", len(tokens))

	// Imprime los tokens con sus tipos
	fmt.Println("Tokens:")
	for i, token := range tokens {
		tokenTypeName := "UNKNOWN"
		if token.GetTokenType() >= 0 && token.GetTokenType() < len(lexer.SymbolicNames) {
			tokenTypeName = lexer.SymbolicNames[token.GetTokenType()]
		} else if token.GetTokenType() == antlr.TokenEOF {
			tokenTypeName = "EOF"
		}
		fmt.Printf(" %d.Token: %q -------> Type: %s\n", i, token.GetText(), tokenTypeName)
	}

	// --- Análisis sintáctico con listener de errores ---
	p := parser.NewProgramaParser(stream)

	// Configura el listener de errores personalizado
	p.RemoveErrorListeners()
	errorListener := &CustomErrorListener{inputText: inputText}
	p.AddErrorListener(errorListener)

	// Intenta parsear el texto
	fmt.Println("\n--- Iniciando Parseo ---")
	errorListener.lastErrorLine = -1
	tree := p.Programa()

	// Imprime el árbol de parseo
	fmt.Println("\n--- Árbol de Parseo ---")
	if tree != nil {
		fmt.Println(tree.ToStringTree(nil, p))
	} else {
		fmt.Println("No se pudo generar el árbol de parseo debido a errores sintácticos.")
	}

	fmt.Println("\n------------------------------------")
}

// Modo de análisis en tiempo real
func runRealtimeAnalysis() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("👨‍💻 Analizador en Tiempo Real (escribe tu código línea por línea, escribe ':exit' para salir)")
	fmt.Println("-----------------------------------------------------------------------------------")

	var inputBuffer strings.Builder
	lineNumber := 1 // Número de línea para el prompt

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

		// Verifica si el usuario quiere salir
		trimmedLine := strings.TrimRight(line, "\r\n")
		if trimmedLine == ":exit" {
			fmt.Println("Saliendo del modo tiempo real.")
			break
		}

		// Agrega la línea al buffer
		inputBuffer.WriteString(line)

		// Realiza el análisis con el texto acumulado
		currentInputText := inputBuffer.String()
		performAnalysis(currentInputText)

		lineNumber++
	}
}

// Función principal
func main() {
	// --- Menú de opciones ---
	fmt.Println("Seleccione la opción:")
	fmt.Println("1. Analizar archivo embebido (files/file.txt)")
	fmt.Println("2. Analizar datos desde teclado (entrada única)")
	fmt.Println("3. Analizar en Tiempo Real (línea por línea)")
	fmt.Print("Opción (1/2/3): ")

	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatalf("Error leyendo opción: %v", err)
	}

	// Manejo de la opción seleccionada
	var data []byte
	var inputText string

	if option == "1" {
		// Opción 1: Analizar archivo embebido
		data, err = fileTxt.ReadFile("files/file.txt")
		if err != nil {
			log.Fatalf("Error leyendo archivo embebido: %v", err)
		}
		inputText = string(data)
		runBatchAnalysis(data, inputText)

	} else if option == "2" {
		// Opción 2: Analizar entrada única desde teclado
		fmt.Println("Ingrese su código (finalice con una línea vacía):")
		reader := bufio.NewReader(os.Stdin)
		var lines []string
		for {
			line, _ := reader.ReadString('\n')
			trimmedLine := strings.TrimRight(line, "\r\n")
			if trimmedLine == "" && len(line) <= 2 {
				break
			}
			lines = append(lines, trimmedLine)
		}
		inputText = strings.Join(lines, "\n")
		data = []byte(inputText)
		runBatchAnalysis(data, inputText)

	} else if option == "3" {
		// Opción 3: Modo de análisis en tiempo real
		runRealtimeAnalysis()

	} else {
		log.Fatal("Opción inválida.")
	}

	fmt.Println("Fin del programa.")
}

// Función auxiliar para realizar análisis en modo batch
func runBatchAnalysis(data []byte, inputText string) {
	var wg sync.WaitGroup
	result := &Result{}

	// Crea el stream de tokens
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

	performAnalysis(inputText)
}
