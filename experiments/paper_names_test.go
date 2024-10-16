package experiments

import (
	"testing"

	"github.com/PhonoGrams/soft_bigram"
)

func TestPaperNames(t *testing.T) {
	weights := make(map[string]soft_bigram.Weights)
	weights["default"] = soft_bigram.DefaultWeights
	weights["table4"] = soft_bigram.Table4Weights
	weights["optimized"] = soft_bigram.OptimizedWeights
	weights["high precision"] = soft_bigram.HighPrecisionWeights
	weights["phonetic"] = soft_bigram.PhoneticWeights

	pairs := [][]string{
		{"precede", "preceed"},
		{"promise", "promiss"},
		{"absence", "absense"},
		{"achieve", "acheive"},
		{"accidentally", "accidentaly"},
		{"algorithm", "algorythm"},
		{"similar", "Similer"},
		{"dilemma", "Dilemma"},
		{"almost", "allmost"},
		{"amend", "ammend"},
		{"occurred", "occured"},
		{"embarrass", "embarass"},
		{"harass", "harrass"},
		{"really", "Realy"},
	}

	runExperiments(t, weights, pairs)
}
