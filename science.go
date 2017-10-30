package main

import (
	"errors"
	"strconv"
	"strings"
)

// ScienceTerm is the object holding the science term data.
type ScienceTerm struct {
	// FIXME: break ID down into:
	// - what scientific act (evaReport)
	// - what body (Kerbin)
	// - what region (Srf for surface)
	// - what mode (Splashed)
	// - what biome (Water)
	ID    string
	Title string
	DSC   int
	SCV   float64
	SBV   float64
	Sci   float64
	Cap   float64
}

// Fill fills the ScienceTerm with data from the Term object.
func Fill(t *Term) (*ScienceTerm, error) {
	s := &ScienceTerm{}
	if t.Terms == nil {
		return nil, errors.New("no terms in science term")
	}
	for _, term := range t.Terms {
		var err error
		if term.Values == nil {
			return nil, errors.New("no values in science term")
		}
		val := strings.Join(term.Values, ", ")
		switch term.Name {
		case "id":
			s.ID = val
		case "title":
			s.Title = val
		case "dsc":
			s.DSC, err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
		case "scv":
			s.SCV, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "sbv":
			s.SBV, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "sci":
			s.Sci, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "cap":
			s.Cap, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid name in science term")
		}
	}
	return s, nil
}
