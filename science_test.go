package main

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/alecthomas/repr"
)

var filltests = []struct {
	in  *Term
	out *ScienceSubject
	err error
}{
	{
		&Term{Name: "empty"},
		nil,
		errors.New("no terms in science term"),
	},
	{
		&Term{Name: "noproperties", Terms: []*Term{
			&Term{Name: "novalue"},
		}},
		nil,
		errors.New("no properties in science term"),
	},
	{
		&Term{Name: "badproperties", Terms: []*Term{
			&Term{Property: "wrong"},
		}},
		nil,
		errors.New("improper property 'wrong' in science term"),
	},
	{
		&Term{Name: "wrongproperties", Terms: []*Term{
			&Term{Property: "wrong = 1"},
		}},
		nil,
		errors.New("invalid name 'wrong' in science term"),
	},
	{
		&Term{Name: "good", Terms: []*Term{
			&Term{Property: "id = recovery@MinmusFlewBy"},
			&Term{Property: "title = Recovery of a vessel returned from a flight over Minmus"},
			&Term{Property: "dsc = 1"},
			&Term{Property: "scv = 0.027777778"},
			&Term{Property: "sbv = 15"},
			&Term{Property: "sci = 17.5"},
			&Term{Property: "cap = 18"},
		}},
		&ScienceSubject{
			ID:              "recovery@MinmusFlewBy",
			Title:           "Recovery of a vessel returned from a flight over Minmus",
			DataScale:       1.0,
			ScientificValue: 0.027777778,
			SubjectValue:    15.0,
			Science:         17.5,
			ScienceCap:      18.0},
		nil,
	},
}

func TestFill(t *testing.T) {
	for _, tt := range filltests {
		var out *ScienceSubject
		var err error
		out, err = Fill(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			errout := new(bytes.Buffer)
			rerr := repr.New(errout)
			errout.WriteString("expected ")
			rerr.Print(tt.out)
			errout.WriteString(", got ")
			rerr.Print(out)
			t.Error(errout.String())
		}
		if !reflect.DeepEqual(err, tt.err) {
			errout := new(bytes.Buffer)
			rerr := repr.New(errout)
			errout.WriteString("expected ")
			if tt.err == nil {
				errout.WriteString("nil")
			} else {
				rerr.Print(tt.err)
			}
			errout.WriteString(", got ")
			if err == nil {
				errout.WriteString("nil")
			} else {
				rerr.Print(err)
			}
			t.Error(errout.String())
		}
	}
}
