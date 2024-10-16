package soft_bigram

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestExperiments(t *testing.T) {
	weightNames := []string{"default", "table4", "optimized", "high precision", "phonetic"}
	weightValues := []Weights{DefaultWeights, Table4Weights, OptimizedWeights, HighPrecisionWeights, PhoneticWeights}

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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	t.Cleanup(func() { w.Flush() })

	var buf bytes.Buffer
	buf.WriteString("s1\ts2")
	for w := range weightNames {
		buf.WriteString("\t" + weightNames[w])
	}
	buf.WriteString("\n")
	w.Write(buf.Bytes())

	for _, pair := range pairs {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%s\t%s", pair[0], pair[1]))

		for w := range weightNames {
			norm := NormalizeSoftBigram(pair[0], pair[1], weightValues[w])
			buf.WriteString(fmt.Sprintf("\t%.2f", norm))
		}

		buf.WriteString("\n")
		w.Write(buf.Bytes())
	}
}
