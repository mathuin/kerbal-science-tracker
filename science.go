package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ScienceSubject is the object holding the data pertaining to this subject.
// https://kerbalspaceprogram.com/api/class_science_subject.html
type ScienceSubject struct {
	// FIXME: break ID down into:
	// - what experiment (evaReport)
	// - what body (Kerbin)
	// - what situation (SrfSplashed) https://kerbalspaceprogram.com/api/_science_8cs.html
	// - what biome (Water)
	// API has this as two calls:
	// static string ScienceUtil.GetExperimentBodyName(string subjectID)
	//
	// static void ScienceUtil.GetExperimentFieldsFromScienceID	(
	//	string subjectID,
	//	out string BodyName,
	//	out string Situation,
	//	out string Biome
	// )
	ID              string
	Title           string
	DataScale       float64
	ScientificValue float64
	SubjectValue    float64
	Science         float64
	ScienceCap      float64
}

// Fill fills the ScienceSubject with data from the Term object.
func Fill(t *Term) (*ScienceSubject, error) {
	s := &ScienceSubject{}
	if t.Terms == nil {
		return nil, errors.New("no terms in science term")
	}
	for _, term := range t.Terms {
		var err error
		if term.Property == "" {
			return nil, errors.New("no properties in science term")
		}
		terms := strings.SplitN(term.Property, " = ", 2)
		if len(terms) != 2 {
			return nil, fmt.Errorf("improper property '%s' in science term", term.Property)
		}
		switch terms[0] {
		case "id":
			s.ID = terms[1]
		case "title":
			s.Title = terms[1]
		case "dsc":
			s.DataScale, err = strconv.ParseFloat(terms[1], 64)
			if err != nil {
				return nil, err
			}
		case "scv":
			s.ScientificValue, err = strconv.ParseFloat(terms[1], 64)
			if err != nil {
				return nil, err
			}
		case "sbv":
			s.SubjectValue, err = strconv.ParseFloat(terms[1], 64)
			if err != nil {
				return nil, err
			}
		case "sci":
			s.Science, err = strconv.ParseFloat(terms[1], 64)
			if err != nil {
				return nil, err
			}
		case "cap":
			s.ScienceCap, err = strconv.ParseFloat(terms[1], 64)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid name '%s' in science term", terms[0])
		}
	}
	return s, nil
}
