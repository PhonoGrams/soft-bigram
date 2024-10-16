package soft_bigram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComputeBigrams(t *testing.T) {
	got := ComputeBigrams("nicolás maduro")
	expected := []string{"ni", "ic", "co", "ol", "lá", "ás", "s ", " m", "ma", "ad", "du", "ur", "ro"}
	require.Equal(t, expected, got)

	got = ComputeBigrams("Raúl Rodríguez")
	expected = []string{"Ra", "aú", "úl", "l ", " R", "Ro", "od", "dr", "rí", "íg", "gu", "ue", "ez"}
	require.Equal(t, expected, got)
}

func TestSoftBigram_DefaultWeights(t *testing.T) {
	require.InDelta(t, 0.489, SoftBigramDistance("adam", "adams", DefaultWeights), 0.01)
	require.InDelta(t, 0.93, NormalizeSoftBigram("adam", "adams", DefaultWeights), 0.01)
}

var Table4Weights = Weights{
	Delete:        1.0, // Corresponds to wt2
	Insert:        0.2, // Corresponds to wt4
	Match:         0.0, // Corresponds to wt1
	Replace:       0.2, // Corresponds to wt5
	Transpose:     0.0, // Corresponds to wt3
	PartialTrans1: 1.0, // Corresponds to wt6
	PartialTrans2: 1.0, // Corresponds to wt7
	LCS1:          0.5, // Corresponds to wt8 (first LCS adjustment)
	LCS2:          0.5, // Corresponds to wt9 (second LCS adjustment)
}

func TestSoftBigram_Table4(t *testing.T) {
	require.InDelta(t, 0.981, NormalizeSoftBigram("precede", "preceed", Table4Weights), 0.01)
	require.InDelta(t, 0.991, NormalizeSoftBigram("promise", "promiss", Table4Weights), 0.01)
	require.InDelta(t, 0.975, NormalizeSoftBigram("absence", "absense", Table4Weights), 0.01)
	require.InDelta(t, 0.953, NormalizeSoftBigram("achieve", "acheive", Table4Weights), 0.01)
	require.InDelta(t, 0.981, NormalizeSoftBigram("accidentally", "accidentaly", Table4Weights), 0.01)
	require.InDelta(t, 0.984, NormalizeSoftBigram("algorithm", "algorythm", Table4Weights), 0.01)
	require.InDelta(t, 0.953, NormalizeSoftBigram("similar", "Similer", Table4Weights), 0.01)
	require.InDelta(t, 0.991, NormalizeSoftBigram("dilemma", "Dilemma", Table4Weights), 0.01)
	require.InDelta(t, 0.991, NormalizeSoftBigram("almost", "allmost", Table4Weights), 0.01)
	require.InDelta(t, 0.990, NormalizeSoftBigram("amend", "ammend", Table4Weights), 0.01)
	require.InDelta(t, 0.969, NormalizeSoftBigram("occurred", "occured", Table4Weights), 0.01)
	require.InDelta(t, 0.973, NormalizeSoftBigram("embarrass", "embarass", Table4Weights), 0.01)
	require.InDelta(t, 0.991, NormalizeSoftBigram("harass", "harrass", Table4Weights), 0.01)
	require.InDelta(t, 0.92, NormalizeSoftBigram("really", "Realy", Table4Weights), 0.01)
}

func TestSoftBigramDistance_Pharmaceutical(t *testing.T) {
	t.Run("default weights", func(t *testing.T) {
		require.InDelta(t, 1.125, SoftBigramDistance("Acetaminophen", "Acetaminaphen", DefaultWeights), 0.01)
		require.InDelta(t, 0.537, SoftBigramDistance("Alprazolam", "Alprazolan", DefaultWeights), 0.01)
		require.InDelta(t, 3.098, SoftBigramDistance("Celexa", "Zyprexa", DefaultWeights), 0.01)
		require.InDelta(t, 5.000, SoftBigramDistance("Prozac", "Paxil", DefaultWeights), 0.01)
		require.InDelta(t, 2.689, SoftBigramDistance("Hydroxyzine", "Hydralazine", DefaultWeights), 0.01)
		require.InDelta(t, 2.889, SoftBigramDistance("Lamisil", "Lamictal", DefaultWeights), 0.01)
		require.InDelta(t, 0.600, SoftBigramDistance("Xanax", "Zanax", DefaultWeights), 0.01)
		require.InDelta(t, 4.860, SoftBigramDistance("Amlodipine", "Amiodarone", DefaultWeights), 0.01)
		require.InDelta(t, 2.889, SoftBigramDistance("Warfarin", "Heparin", DefaultWeights), 0.01)
		require.InDelta(t, 3.098, SoftBigramDistance("Toprol", "Topomax", DefaultWeights), 0.01)
	})
}

func TestSoftBigramDistance_Names(t *testing.T) {
	require.InDelta(t, 0.48, SoftBigramDistance("john", "johnny", OptimizedWeights), 0.01)
	require.InDelta(t, 0.96, SoftBigramDistance("john", "johnthan", OptimizedWeights), 0.01)

	require.InDelta(t, 0, SoftBigramDistance("Zhao Wei", "Zhao Wei", OptimizedWeights), 0.01)
	require.InDelta(t, 0, SoftBigramDistance("zhao wei", "zhao wei", OptimizedWeights), 0.01)

	// both directions
	require.InDelta(t, 5.010, SoftBigramDistance("jane doe", "jan lahore", OptimizedWeights), 0.01)
	require.InDelta(t, 5.010, SoftBigramDistance("jan lahore", "jane doe", OptimizedWeights), 0.01)

	// real world case
	require.InDelta(t, 4.687, SoftBigramDistance("john doe", "paul john", OptimizedWeights), 0.01)
	require.InDelta(t, 4.298, SoftBigramDistance("john doe", "john othername", OptimizedWeights), 0.01)

	// close match
	require.InDelta(t, 0.24, SoftBigramDistance("jane doe", "jane doe2", OptimizedWeights), 0.01)

	// real-ish world examples
	require.InDelta(t, 5.690, SoftBigramDistance("kalamity linden", "kala limited", OptimizedWeights), 0.01)

	// examples used in demos / commonly
	require.InDelta(t, 0.00, SoftBigramDistance("nicolas", "nicolas", OptimizedWeights), 0.01)
	require.InDelta(t, 1.44, SoftBigramDistance("nicolas moros maduro", "nicolas maduro", OptimizedWeights), 0.01)

}
