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
		&Term{Name: "novalues", Terms: []*Term{
			&Term{Name: "novalue"},
		}},
		nil,
		errors.New("no values in science term"),
	},
	{
		&Term{Name: "wrongvalues", Terms: []*Term{
			&Term{Name: "wrongvalue", Values: []string{"wrong"}},
		}},
		nil,
		errors.New("invalid name in science term"),
	},
	{
		&Term{Name: "good", Terms: []*Term{
			&Term{Name: "id", Values: []string{"id"}},
			&Term{Name: "title", Values: []string{"title"}},
			&Term{Name: "dsc", Values: []string{"1"}},
			&Term{Name: "scv", Values: []string{"2.0"}},
			&Term{Name: "sbv", Values: []string{"3.4"}},
			&Term{Name: "sci", Values: []string{"5.6"}},
			&Term{Name: "cap", Values: []string{"7.8"}},
		}},
		&ScienceSubject{
			ID:              "id",
			Title:           "title",
			DataScale:       1,
			ScientificValue: 2,
			SubjectValue:    3.4,
			Science:         5.6,
			ScienceCap:      7.8},
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
