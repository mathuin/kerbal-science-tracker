package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// https://gist.github.com/mashbridge/4365101
	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
		break
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		check(err)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}

	savefile := &SaveFile{}
	savefile.Load(data)

	scienceTerms := savefile.GetTerms("Science")

	var numterms int
	var totsci float64
	var capsci float64

	for _, term := range scienceTerms {
		//		repr.Println(term)
		var st *ScienceTerm
		st, err = Fill(term)
		if err != nil {
			continue
		}
		numterms = numterms + 1
		totsci = totsci + st.Sci
		capsci = capsci + st.Cap
	}

	fmt.Printf("%d missions, %0.2f total science, %0.2f capacity science\n", numterms, totsci, capsci)

}
