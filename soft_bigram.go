package soft_bigram

import (
	"math"
)

// ComputeBigrams generates bigrams from a given string using runes.
func ComputeBigrams(s string) []string {
	runes := []rune(s)
	var bigrams []string
	for i := 0; i < len(runes)-1; i++ {
		// if runes[i] == ' ' || runes[i+1] == ' ' {
		// 	continue
		// }
		bigram := string(runes[i]) + string(runes[i+1])
		bigrams = append(bigrams, bigram)
	}
	return bigrams
}

// LCS computes the Longest Common Subsequence length between two bigram sequences.
func LCS(bigrams1, bigrams2 []string) int {
	m, n := len(bigrams1), len(bigrams2)
	lcsTable := make([][]int, m+1)
	for i := range lcsTable {
		lcsTable[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if bigrams1[i-1] == bigrams2[j-1] {
				lcsTable[i][j] = lcsTable[i-1][j-1] + 1
			} else {
				lcsTable[i][j] = int(math.Max(float64(lcsTable[i-1][j]), float64(lcsTable[i][j-1])))
			}
		}
	}
	return lcsTable[m][n]
}

// SoftBigramDistance computes the Soft-Bidist distance between two strings using runes.
func SoftBigramDistance(s1, s2 string, weights Weights) float64 {
	bigrams1 := ComputeBigrams(s1)
	bigrams2 := ComputeBigrams(s2)

	// Distance matrix
	dist := make([][]float64, len(bigrams1)+1)
	for i := range dist {
		dist[i] = make([]float64, len(bigrams2)+1)
	}

	// Initialize base cases
	for i := 0; i <= len(bigrams1); i++ {
		dist[i][0] = float64(i) * weights.Delete
	}
	for j := 0; j <= len(bigrams2); j++ {
		dist[0][j] = float64(j) * weights.Insert
	}

	// Fill the distance matrix using Soft-Bidist rules
	for i := 1; i <= len(bigrams1); i++ {
		for j := 1; j <= len(bigrams2); j++ {
			if bigrams1[i-1] == bigrams2[j-1] {
				// Case 1: Match
				dist[i][j] = dist[i-1][j-1] + weights.Match
			} else {
				// Case 2: Replacement (Different bigrams)
				dist[i][j] = math.Min(
					math.Min(dist[i-1][j]+weights.Delete, dist[i][j-1]+weights.Insert),
					dist[i-1][j-1]+weights.Replace,
				)

				// Case 3: Transposition
				if i > 1 && j > 1 && bigrams1[i-1] == bigrams2[j-2] && bigrams1[i-2] == bigrams2[j-1] {
					dist[i][j] = math.Min(dist[i][j], dist[i-2][j-2]+weights.Transpose)
				}

				// Case 4 & 5: Partial Matches
				if i > 1 && bigrams1[i-1] == bigrams2[j-1] && bigrams1[i-2] != bigrams2[j-2] {
					dist[i][j] = math.Min(dist[i][j], dist[i-2][j-2]+weights.PartialTrans1)
				}
				if j > 1 && bigrams1[i-1] == bigrams2[j-1] && bigrams1[i] != bigrams2[j-2] {
					dist[i][j] = math.Min(dist[i][j], dist[i-1][j-2]+weights.PartialTrans2)
				}
			}
		}
	}

	// Incorporate LCS as an additional adjustment factor using two LCS weights
	lcsLength := LCS(bigrams1, bigrams2)
	lcsFactor := float64(lcsLength) / math.Min(float64(len(bigrams1)), float64(len(bigrams2)))
	finalDistance := dist[len(bigrams1)][len(bigrams2)] * (1 - weights.LCS1*lcsFactor) * (1 - weights.LCS2*lcsFactor)
	finalDistance = math.Max(finalDistance, 0)

	return finalDistance
}
