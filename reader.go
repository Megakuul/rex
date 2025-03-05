package main

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

func (r *reader) peak(count int) string {
	output := ""

	for i:=1; i <= count; i++ {
		if r.index+i >= len(r.data) {
			return output
		}
		output += string(r.data[r.index+i])
	}
	return output
}
