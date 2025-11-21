package unixsort

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/andreyxaxa/unixsort/pkg/unixsort/helpers"
)

// Params is core struct for sort util
type Params struct {
	BlanksIgnore bool
	Month        bool
	Human        bool
	Numeric      bool
	Reverse      bool
	Unique       bool
	Column       int
	//Sorted       bool // TODO

	reader io.ReadCloser
}

// NewParams returns new 'params' for sort util
func NewParams() *Params {
	return &Params{}
}

func (p *Params) parse() error {
	expArgs := []string{}
	for _, a := range os.Args[1:] {
		// допустим, -nru
		if strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") && len(a) > 2 {
			// [1:] - 'n', 'r', 'u'
			for _, c := range a[1:] {
				// ["-n", "-r", "-u"]
				expArgs = append(expArgs, "-"+string(c))
			}
		} else {
			expArgs = append(expArgs, a)
		}
	}

	// пустой набор флагов
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// ["app", "-nru", "text.txt"] -> ["app", "-n", "-r", "-u", text.txt]
	os.Args = append([]string{os.Args[0]}, expArgs...)

	flag.BoolVar(&p.Numeric, "n", false, "by string numeric value")
	flag.BoolVar(&p.Reverse, "r", false, "reversed order")
	flag.BoolVar(&p.Unique, "u", false, "without duplicates")
	flag.IntVar(&p.Column, "k", 0, "by column number")
	flag.BoolVar(&p.BlanksIgnore, "b", false, "ignore trailing blanks")
	flag.BoolVar(&p.Month, "M", false, "by month name")
	flag.BoolVar(&p.Human, "h", false, "by nums with suffix(K, M, G, T)")
	//flag.BoolVar(&p.Sorted, "c", false, "sorted or no") // TODO

	flag.Parse()
	args := flag.Args()

	// проверяем, как поступают данные
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		p.reader = os.Stdin
		return nil
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s - %v\n", args[0], err)
		return err
	}
	p.reader = file

	return nil
}

// Start run sort util
func (p *Params) Start() error {
	if err := p.parse(); err != nil {
		return err
	}

	defer p.reader.Close()

	lines := make([]string, 0)
	helpers.ReadLines(p.reader, &lines)

	sorted := p.Sort(lines)

	helpers.Out(sorted)

	return nil
}

// Sort takes unsorted 'lines' and returns a sorted 'lines'
func (p *Params) Sort(lines []string) []string {
	kf := func(line string) string {
		if p.BlanksIgnore {
			line = strings.TrimRight(line, " \t")
		}
		if p.Column > 0 {
			return helpers.GetColumn(line, p.Column)
		}
		return line
	}

	// сортировка
	sort.Slice(lines, func(i, j int) bool {
		// каждую строку отдаем в 'kf', если '-k', тогда будет извлекаться столбец, иначе просто целая строка
		ki := kf(lines[i])
		kj := kf(lines[j])

		// если -M
		m := months()
		if p.Month {
			mi := m[strings.ToLower(ki)]
			mj := m[strings.ToLower(kj)]
			if mi != 0 && mj != 0 {
				return mi < mj
			}
		}

		// Если -h
		if p.Human {
			hi, ok1 := helpers.HumanSuffNums(ki)
			hj, ok2 := helpers.HumanSuffNums(kj)
			if ok1 && ok2 {
				return hi < hj
			}
		}

		// если -n
		// конвертируем в числа и сравниваем
		if p.Numeric {
			fi, err1 := strconv.ParseFloat(strings.TrimSpace(ki), 64)
			fj, err2 := strconv.ParseFloat(strings.TrimSpace(kj), 64)

			if err1 == nil && err2 == nil {
				return fi < fj
			}
		}
		// обычная сортировка строк
		return ki < kj
	})

	// если -r
	if p.Reverse {
		slices.Reverse(lines)
	}

	// если -u
	if p.Unique {
		uniquesort(&lines)
	}

	return lines
}

// for -u
// отдаем уже отсортированные данные
func uniquesort(s *[]string) {
	if len(*s) == 0 {
		return
	}

	out := make([]string, 0, len(*s))
	prev := (*s)[0]
	out = append(out, prev)

	for i := 1; i < len(*s); i++ {
		if (*s)[i] != prev {
			out = append(out, (*s)[i])
			prev = (*s)[i]
		}
	}
	*s = out
}

// for -m
func months() map[string]int {
	return map[string]int{
		"jan": 1, "feb": 2, "mar": 3, "apr": 4,
		"may": 5, "jun": 6, "jul": 7, "aug": 8,
		"sep": 9, "oct": 10, "nov": 11, "dec": 12,
	}
}
