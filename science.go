package main

import (
	"errors"
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
	DataScale       int
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
			s.DataScale, err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
		case "scv":
			s.ScientificValue, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "sbv":
			s.SubjectValue, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "sci":
			s.Science, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		case "cap":
			s.ScienceCap, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid name in science term")
		}
	}
	return s, nil
}
