package helpers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads from 'r' to 'lines' (bufio.Scanner)
func ReadLines(r io.Reader, lines *[]string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		*lines = append(*lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatalf("file scan error: %v\n", err)
	}
}

// GetColumn returns 'column'th word from 'line'
func GetColumn(line string, column int) string {
	parts := strings.Split(line, " ")
	if column <= 0 || column > len(parts) {
		return ""
	}

	return parts[column-1]
}

// Out writes 'lines' to stdout
func Out(lines []string) {
	for _, l := range lines {
		fmt.Fprintln(os.Stdout, l)
	}
}

func HumanSuffNums(line string) (float64, bool) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return 0, false
	}

	// последний символ - мб суффикс
	last := line[len(line)-1]
	// запоминаем
	l := line

	// множитель
	var m float64 = 1

	switch last {
	case 'K', 'k':
		m = 1024
		// всё до последнего символа
		l = line[:len(line)-1]
	case 'M', 'm':
		m = 1024 * 1024
		l = line[:len(line)-1]
	case 'G', 'g':
		m = 1024 * 1024 * 1024
		l = line[:len(line)-1]
	case 'T', 't':
		m = 1024 * 1024 * 1024 * 1024
		l = line[:len(line)-1]
	}

	// переводим в число
	n, err := strconv.ParseFloat(l, 64)
	if err != nil {
		return 0, false
	}

	return n * m, true
}
