package main

import (
	"bytes"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var savefileLexer = lexer.Must(lexer.Regexp(
	`(?m)` +
		`^([\t\f\r ]+)` +
		`|(?P<Property>[^\n]+)(?P<Equal> = ?)(?P<Value>[^\n]+)?` +
		`|(?P<Name>[^\n}{]+)` +
		`|(?P<Open>{)` +
		`|(?P<Close>})` +
		`|(?P<Newline>\n)`,
))

// Load fills a SaveFile with Terms from a byte array.
func (s *SaveFile) Load(b []byte) error {
	gstring := bytes.NewReader(b)

	parser, err := participle.Build(&SaveFile{}, participle.Lexer(savefileLexer))
	if err != nil {
		return err
	}
	err = parser.Parse(gstring, s)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile represents the entire save file.
type SaveFile struct {
	Term *Term `parser:"{ @@ }"`
}

// Term is either a Group or a Property.
// Group information is stored in Terms while Property information is stored in Values.
type Term struct {
	Name     string  `parser:"( @Name Newline"`
	Terms    []*Term `parser:"Open Newline { @@ } Close { Newline } )"`
	Property string  `parser:"| @Property Newline"`
}

// GetTerms returns a list of pointers to Term objects with names matching the string argument.
func (s *SaveFile) GetTerms(m string) []*Term {
	var terms []*Term

	ch := Walker(s, m)
	for {
		term, ok := <-ch
		if !ok {
			break
		}
		terms = append(terms, term)
	}

	return terms
}

// Walk traverses a tree of Term objects in search of one with a particular name.
// See https://golang.org/doc/play/tree.go
func Walk(t *Term, m string, ch chan *Term) {
	if t.Name == m {
		ch <- t
	}
	if t.Terms == nil {
		return
	}
	for _, term := range t.Terms {
		Walk(term, m, ch)
	}
}

// Walker drives the Walk function.
// See https://golang.org/doc/play/tree.go
func Walker(s *SaveFile, m string) <-chan *Term {
	ch := make(chan *Term)
	go func() {
		Walk(s.Term, m, ch)
		close(ch)
	}()
	return ch
}
