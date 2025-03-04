package rex

import (
	"fmt"
	"os"
	"unicode"
)




type reader struct {
	data []rune
	index int
}

func newReader(input string) *reader {
	reader := &reader{
		data: []rune{},
		index: 0,
	}
	for _, r := range input {
		reader.data = append(reader.data, r)
	}
	return reader
}

func (r *reader) eat() (rune, bool) {
	if r.index >= len(r.data) {
		return 0, false
	}
	result := r.data[r.index]
	r.index++
	return result, true
}

type Regex struct {
	chain []func(*reader) (string, error)
}

func Parse(expr string) (*Regex, error) {
	for _, r := range expr {
		switch r {
		case '[':

		}
	}
}

func (r *Regex) Match(input string) ([]string, error) {
	reader := newReader(input)
	for _, fn := range r.chain {
		match, err := fn(reader)
		if err!=nil {
			return nil, err
		}
		println(match)
	}
	return []string{}, nil
}

func main() {
	regex, err := Parse("[a-z]")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	groups := regex.Match("a")
	if len(groups) < 1 {
		fmt.Println("doesnt match")
		os.Exit(1)
	}
}
