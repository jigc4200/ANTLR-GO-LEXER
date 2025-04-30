package utils

import (
	"bufio"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

func countChars(data []byte, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()
	res.Chars = len(data)
}

func countLines(data []byte, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()
	lines := 0
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		lines++
	}
	if len(data) > 0 && lines == 0 {
		lines = 1
	}
	res.Lines = lines
}

func countTokens(stream *antlr.CommonTokenStream, wg *sync.WaitGroup, res *Result) {
	defer wg.Done()
	stream.Fill()
	res.Tokens = len(stream.GetAllTokens())
}
