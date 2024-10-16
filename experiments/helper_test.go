package experiments

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"testing"
	"text/tabwriter"

	"github.com/PhonoGrams/soft_bigram"

	"golang.org/x/exp/maps"
)

func runExperiments(tb testing.TB, weights map[string]soft_bigram.Weights, pairs [][]string) {
	weightNames := maps.Keys(weights)
	slices.Sort(weightNames)

	var lines []string
	for _, pair := range pairs {
		for i := 1; i < len(pair); i++ {
			var total float64

			var buf bytes.Buffer
			buf.WriteString(fmt.Sprintf("%s\t%s", pair[0], pair[i]))

			for _, w := range weightNames {
				norm := soft_bigram.NormalizeSoftBigram(pair[0], pair[i], weights[w])
				total += norm
				buf.WriteString(fmt.Sprintf("\t%.2f", norm))
			}

			buf.WriteString("\t|")
			buf.WriteString(fmt.Sprintf("\t%.2f", total/float64(len(weightNames))))

			lines = append(lines, buf.String())

		}
	}

	// Write results as table to stdout
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	tb.Cleanup(func() { w.Flush() })

	var buf bytes.Buffer
	buf.WriteString("s1\ts2")
	for w := range weightNames {
		buf.WriteString("\t" + weightNames[w])
	}
	buf.WriteString("\t|")
	buf.WriteString("\t" + "average")
	buf.WriteString("\n")
	w.Write(buf.Bytes())

	slices.Sort(lines)
	for _, line := range lines {
		fmt.Fprint(w, line+"\n")
	}
}
