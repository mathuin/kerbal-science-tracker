package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
)

var lexertests = []struct {
	in  []string
	out []string
}{
	{[]string{"GAME", "{", "    Hello = World", "}"}, []string{"Name", "Newline", "Open", "Newline", "Property", "Newline", "Close", "EOF"}},
	{[]string{"Science", "{", "	id = temperatureScan@KerbinSrfLandedLaunchPad", "	title = Temperature Scan from LaunchPad", "	dsc = 1", "	scv = 0", "	sbv = 0.300000012", "	sci = 1.20000005", "	cap = 1.20000005", "}"}, []string{"Name", "Newline", "Open", "Newline", "Property", "Newline", "Property", "Newline", "Property", "Newline", "Property", "Newline", "Property", "Newline", "Property", "Newline", "Property", "Newline", "Close", "EOF"}},
	{[]string{"file", "{", "		filename = launch.ks", "		binary = H4sIADjSWlwA,12NQQrCMBRE180pPlm1UIq4syAoKlqEVpq6lq+mTSAmIUkFFe9uu1NnOe8Ns1D4fHTKnFGBaduMWHR444E7WFoD0sN0MuYbiGs3gtlQEtnC,qjlnTvP86outkW5WRdNVcMcKDvsKBgHTEibl8M6XxkdUGofU8aaiibwIlHkem0xiJgq7PVFnLwPhqZoTTo8JRl5A1eej+q,if7yI5IPAoH67M8AAAA=", "}"}, []string{"Name", "Newline", "Open", "Newline", "Property", "Newline", "Property", "Newline", "Close", "EOF"}},
	{[]string{"bad", "{", "    linkURL =", "}"}, []string{"Name", "Newline", "Open", "Newline", "Property", "Newline", "Close", "EOF"}},
}

func TestLexer(t *testing.T) {
	for _, tt := range lexertests {

		gstring := strings.NewReader(strings.Join(tt.in, "\n"))
		glex, err := savefileLexer.Lex(gstring)
		if err != nil {
			t.Error(err)
		}
		tokens, err := lexer.ConsumeAll(glex)
		if err != nil {
			t.Error(err)
		}

		revSymbols := make(map[rune]string)
		for k, v := range savefileLexer.Symbols() {
			revSymbols[v] = k
		}

		if len(tokens) != len(tt.out) {
			t.Errorf("token count %d not equal expected %d", len(tokens), len(tt.out))
		}

		for i, token := range tokens {
			if token.Type != savefileLexer.Symbols()[tt.out[i]] {
				t.Errorf("token %d (%s) was %v not %v", i, token, revSymbols[token.Type], tt.out[i])
				// } else {
				// 	fmt.Printf("token %d (%s) was %v as expected\n", i, token, tt.out[i])
			}
		}

	}
}

var parsertests = []struct {
	in  []string
	out *SaveFile
}{
	{[]string{""}, &SaveFile{}},
	{[]string{"GAME", "{", "}"}, &SaveFile{&Term{Name: "GAME"}}},
	{[]string{"GAME", "{", "    Hello = World", "}"}, &SaveFile{&Term{Name: "GAME", Terms: []*Term{&Term{Property: "Hello = World"}}}}},
	{[]string{"GAME", "{", "    INNER", "    {", "    }", "}"}, &SaveFile{&Term{Name: "GAME", Terms: []*Term{&Term{Name: "INNER"}}}}},
	{[]string{"GAME", "{", "    INNER", "    {", "        Hello = World", "    }", "}"}, &SaveFile{&Term{Name: "GAME", Terms: []*Term{&Term{Name: "INNER", Terms: []*Term{&Term{Property: "Hello = World"}}}}}}},
}

func TestParser(t *testing.T) {
	for _, tt := range parsertests {

		savefile := &SaveFile{}
		err := savefile.Load([]byte(strings.Join(tt.in, "\n")))
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(savefile, tt.out) {
			errout := new(bytes.Buffer)
			rerr := repr.New(errout)
			errout.WriteString("expected ")
			rerr.Print(tt.out)
			errout.WriteString(", got ")
			rerr.Print(savefile)
			t.Error(errout.String())
		}
	}
}

var walkertests = []struct {
	in  []string
	out int
}{
	{[]string{"GAME", "{", "}"}, 0},
	{[]string{"GAME", "{", "    INNER", "    {", "    }", "}"}, 1},
	{[]string{"GAME", "{", "    INNER", "    {", "        Hello = World", "    }", "}"}, 1},
}

func TestWalker(t *testing.T) {
	for _, tt := range walkertests {

		savefile := &SaveFile{}
		savefile.Load([]byte(strings.Join(tt.in, "\n")))

		terms := savefile.GetTerms("INNER")

		if len(terms) != tt.out {
			t.Errorf("%s: expected %d terms got %d", tt.in, tt.out, len(terms))
		}

	}
}
