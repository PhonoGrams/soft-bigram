package soft_bigram

// NormalizeSoftBigram computes the Soft-Bigram distance and converts the distance into a similarity score in the range [0, 1].
func NormalizeSoftBigram(s1, s2 string, weights Weights) float64 {
	rawDistance := SoftBigramDistance(s1, s2, weights)

	bigrams1 := ComputeBigrams(s1)
	bigrams2 := ComputeBigrams(s2)

	// Calculate the maximum possible distance for normalization
	maxDistance := float64(len(bigrams1))*weights.Delete + float64(len(bigrams2))*weights.Insert

	// Handle the case when maxDistance is zero to avoid division by zero
	if maxDistance == 0 {
		if rawDistance == 0 {
			return 1.0 // Identical strings
		}
		return 0.0 // Completely different strings
	}

	normalizedDistance := rawDistance / maxDistance
	similarity := 1 - normalizedDistance

	return similarity
}

// // HybridSimilarity combines Soft-Bigram and Jaro-Winkler scores with equal weight
// func HybridSimilarity(s1, s2 string, weights map[string]float64) float64 {
// 	// Normalize Soft-Bigram similarity
// 	softBigramSimilarity := NormalizeSoftBigram(s1, s2, weights)

// 	// Calculate the Jaro-Winkler similarity
// 	jaroWinklerSimilarity := JaroWinkler(s1, s2)

// 	// Balanced approach: equal weights for both methods
// 	alpha, beta := 0.5, 0.5
// 	hybridSimilarity := (alpha * softBigramSimilarity) + (beta * jaroWinklerSimilarity)

// 	return hybridSimilarity
// }

// func main() {
// 	name1 := "jennifer"
// 	name2 := "jeniffer"

// 	// Example weights
// 	weights := map[string]float64{
// 		"match":    0,   // Exact match of bigrams
// 		"replace":  1.0, // Different characters in corresponding positions
// 		"insert":   1.0, // Insert a character
// 		"delete":   1.0, // Delete a character
// 		"transpos": 0.5, // Transpose characters
// 		"lcs":      0.3, // Weight for LCS adjustment factor
// 	}

// 	hybridSim := HybridSimilarity(name1, name2, weights)
// 	fmt.Printf("Hybrid Similarity between '%s' and '%s': %.2f\n", name1, name2, hybridSim)
// }
