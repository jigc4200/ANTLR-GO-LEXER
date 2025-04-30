package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"antlr-go-example/utils"
)

//go:embed files/file.txt
var fileTxt embed.FS

func main() {
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

	switch option {
	case "1":
		data, err := fileTxt.ReadFile("files/file.txt")
		if err != nil {
			log.Fatalf("Error leyendo archivo embebido: %v", err)
		}
		utils.RunBatchAnalysis(data, string(data))
	case "2":
		fmt.Println("Ingrese su código (finalice con una línea vacía):")
		reader := bufio.NewReader(os.Stdin)
		var lines []string
		for {
			line, _ := reader.ReadString('\n')
			trimmedLine := strings.TrimRight(line, "\r\n")
			if trimmedLine == "" {
				break
			}
			lines = append(lines, trimmedLine)
		}
		inputText := strings.Join(lines, "\n")
		utils.RunBatchAnalysis([]byte(inputText), inputText)
	case "3":
		utils.RunRealtimeAnalysis()
	default:
		log.Fatal("Opción inválida.")
	}

	fmt.Println("Fin del programa.")
}
