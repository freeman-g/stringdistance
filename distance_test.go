package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFuzzyMatches(t *testing.T) {
	sources := []string{
		"A river runs",
		"The lord of the rings",
		"Hunger Games",
		"The golden sun",
		"Red rising",
	}

	targets := []string{
		"A river runs through it",
		"lord of the rings",
		"The Hunger games",
		"The Golden Son",
	}

	resp := getFuzzyMatches(sources, targets)

	expectedDistances := []Distance{
		Distance{"A river runs", "A river runs through it", 11},
		Distance{"Hunger Games", "The Hunger games", 5},
		Distance{"The lord of the rings", "lord of the rings", 4},
	}
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, expectedDistances, resp.Distances)

}

func TestGetDistanceParsingErrors(t *testing.T) {

	//bad data in sources
	sources := `"A river runs through it"    , "Lotr"`
	targets := `"A river runs","Lord of tr"`

	dr := DistanceRequest{Source: sources, Target: targets}
	resp := getDistanceResponse(dr)

	assert.Equal(t, false, resp.Success)
	assert.Equal(t, `Error parsing source words comma delimted string: Error parsing comma delimited list.  Make sure there are no spaces before or after commas: line 1, column 24: extraneous " in field`, resp.Message)

	//bad data in targets
	sources = `"A river runs through it","Lotr"`
	targets = `"A river runs"  ,"Lord of tr"`

	dr = DistanceRequest{Source: sources, Target: targets}
	resp = getDistanceResponse(dr)

	assert.Equal(t, false, resp.Success)
	assert.Equal(t, `Error parsing target words comma delimted string: Error parsing comma delimited list.  Make sure there are no spaces before or after commas: line 1, column 13: extraneous " in field`, resp.Message)

}

func TestCSVToUpperSlice(t *testing.T) {

	//tests the clean up to remove carriage returns and extra whitespace
	dirtySources := `"A river runs through it",
	      "The lord of the rings",
		"Hunger games",
		    "The golden sun",
		     "Red rising,"`

	result, _ := csvToSlice(dirtySources)

	expected := []string{
		"A river runs through it",
		"The lord of the rings",
		"Hunger games",
		"The golden sun",
		"Red rising,",
	}

	assert.Equal(t, expected, result)

}
