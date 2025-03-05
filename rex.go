package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

// character classes:
// - macros (\p{} \d \xy) -> resolved in macro stage
// - single chars (a, b, c)
// - custom char classes (all that is encapsulated in [])

type Regex struct {
	chain []func(*reader) (string, error)
}

func Parse(expr string) (*Regex, error) {
	regex := &Regex{
		chain: []func(*reader) (string, error){},
	}

	var lastFunc func(*reader) (string, error) = nil
	
	for _, r := range expr {
		switch r {
		case '*':
			
		default:
			// TODO parse custom char classes aswell (this only supports single chars)
			table := &unicode.RangeTable{
				R16: []unicode.Range16{
					{Lo: uint16(r), Hi: uint16(r), Stride: 1},
				},
			}
			lastFunc = singleMatch(table)
		}
		regex.chain = append(regex.chain, lastFunc)
	}

	return regex, nil
}

func singleMatch(table *unicode.RangeTable) func(*reader) (string, error) {
	return func(reader *reader) (string, error) {
		c, ok := reader.eat()
		if !ok {
			return "", io.EOF
		}
		if unicode.Is(table, c) {
			return string(c), nil
		} else {
			return "", fmt.Errorf("rune '%c' does not match expected class", c)
		}
	}
}

// just poc. I know the code is crap. stfu
func multiMatch(loopConditions []func(*reader) (string, error), stopCondition func(reader *reader) (string, error)) func(*reader) (string, error) {
	return func(reader *reader) (string, error) {
		for {
			peakReader := newReader(reader.peak(len(loopConditions)))
			for _, condition := range loopConditions {
				match, err := condition(peakReader)
				if err!=nil {
					return "", err
				}
			}
			if !ok {
				return "", io.EOF
			}
			c, ok := reader.eat()
			if !ok {
				return "", io.EOF
			}
			
			if unicode.Is(table, c) {
				return string(c), nil
			} else {
				return "", fmt.Errorf("rune '%c' does not match expected class", c)
			}
		}
	}
}

func (r *Regex) Match(input string) ([]string, error) {
	rootMatch := ""
	reader := newReader(input)
	for _, fn := range r.chain {
		match, err := fn(reader)
		if err!=nil {
			return nil, err
		}
		rootMatch += match
	}
	return []string{rootMatch}, nil
}

func main() {
	regex, err := Parse("abdc")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	groups, err := regex.Match("abdcasdf")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Printf("Matches %v", groups)
}
