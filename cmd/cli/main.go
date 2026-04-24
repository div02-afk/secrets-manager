package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseByteList(input string) ([]byte, error) {
	cleaned := strings.TrimSpace(input)
	cleaned = strings.TrimPrefix(cleaned, "[")
	cleaned = strings.TrimSuffix(cleaned, "]")

	if strings.TrimSpace(cleaned) == "" {
		return nil, fmt.Errorf("empty input")
	}

	parts := strings.Fields(cleaned)
	result := make([]byte, 0, len(parts))

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %q: %w", p, err)
		}
		if n < 0 || n > 255 {
			return nil, fmt.Errorf("value out of byte range: %d", n)
		}
		result = append(result, byte(n))
	}

	return result, nil
}

func main() {
	input := strings.Join(os.Args[1:], " ")
	if strings.TrimSpace(input) == "" {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil && strings.TrimSpace(line) == "" {
			fmt.Fprintln(os.Stderr, "provide input as args or stdin, e.g. [91 182 216]")
			os.Exit(1)
		}
		input = line
	}

	b, err := parseByteList(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err)
		os.Exit(1)
	}

	fmt.Println(hex.EncodeToString(b))
}
