package main

import (
	"encoding/csv"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

type DistanceRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type DistanceResponse struct {
	Success   bool
	Message   string     `json:",omitempty"`
	Distances []Distance `json:"results,omitempty"`
}

type Distance struct {
	Source   string
	Target   string
	Distance int
}

func getDistanceResponse(dr DistanceRequest) DistanceResponse {

	sources, sourceWordsErr := csvToSlice(dr.Source)

	if sourceWordsErr != nil {
		message := fmt.Sprintf("Error parsing source words comma delimted string: %v", sourceWordsErr)
		return DistanceResponse{Success: false, Message: message}
	}

	targets, targetWordsErr := csvToSlice(dr.Target)

	if targetWordsErr != nil {
		message := fmt.Sprintf("Error parsing target words comma delimted string: %v", targetWordsErr)
		return DistanceResponse{Success: false, Message: message}
	}

	response := getFuzzyMatches(sources, targets)

	return response
}

func csvToSlice(csvString string) ([]string, error) {

	csvString = strings.Replace(csvString, "\n", "", -1)
	reg, _ := regexp.Compile(`",\s+`)
	csvString = reg.ReplaceAllString(csvString, `",`)

	r := csv.NewReader(strings.NewReader(csvString))

	records, err := r.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("Error parsing comma delimited list.  Make sure there are no spaces before or after commas: %v", err)
	}

	return records[0], nil

}

func getFuzzyMatches(sources, targets []string) DistanceResponse {

	var distanceResponse DistanceResponse
	var distances []Distance
	var firstPassUnmatched []string
	matched := false

	for _, sourceWord := range sources {
		matches := fuzzy.RankFindFold(sourceWord, targets)
		sort.Sort(matches)

		for _, match := range matches {
			d := Distance{Source: match.Source, Target: match.Target, Distance: match.Distance}
			distances = append(distances, d)
			matched = true
		}

		if !matched {
			firstPassUnmatched = append(firstPassUnmatched, sourceWord)
		}

		matched = false

	}

	for _, sourceWord := range targets {
		matches := fuzzy.RankFindFold(sourceWord, firstPassUnmatched)
		sort.Sort(matches)

		for _, match := range matches {
			d := Distance{Source: match.Target, Target: match.Source, Distance: match.Distance}
			distances = append(distances, d)
			matched = true
		}

	}

	distanceResponse.Distances = distances
	distanceResponse.Success = true
	return distanceResponse

}
