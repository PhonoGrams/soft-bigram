package soft_bigram

// In the Soft-Bigram paper, weights \( w_1 \) through \( w_9 \) are used to represent different scenarios when comparing two bigrams (pairs of characters) in strings. These weights align with specific operations and how characters match or differ. Here’s a simplified explanation of how these weights connect to a structure that measures similarity between strings:

// w1: Match - This weight is used when characters from both strings match exactly. It directly corresponds to a "Match" condition.
// w2: Replacement - This is for situations where characters in both strings are different. It aligns with the idea of replacing one character with another.
// w3: Transposition - This applies when two characters are swapped between the strings. It accounts for character swaps and is crucial for measuring similarity when characters are out of order.
// w4: Insertion - This weight is used when a character is inserted in one string but doesn’t appear in the other.
// w5: Deletion - This applies when a character is missing in one string but present in the other.
// w6: Partial Transposition - This handles cases where only part of a character swap occurs (e.g., only one of the two characters is in the correct spot).
// w7: Partial Matching - This applies when characters align but there are small differences in the surrounding or adjacent characters.
// w8: LCS Adjustment 1 - This weight adjusts the score based on the longest common subsequence (LCS) to account for sequences that are mostly similar but have minor differences.
// w9: LCS Adjustment 2 - This is another adjustment related to the LCS, used for cases that need further differentiation based on sequence patterns.

// These weights are carefully tuned through statistical and evolutionary methods to match specific applications, such as comparing multilingual names or minimizing errors in medical contexts where name similarity is critical. The tuning process ensures that the algorithm adapts to the patterns present in the data it analyzes.

// Weights struct for Soft-Bigram configuration
type Weights struct {
	Delete        float64 // wt2
	Insert        float64 // wt4
	Match         float64 // wt1
	Replace       float64 // wt5
	Transpose     float64 // wt3
	PartialTrans1 float64 // wt6
	PartialTrans2 float64 // wt7
	LCS1          float64 // wt8
	LCS2          float64 // wt9
}

// Popular weight configuration inspired by LASA error minimization research
var DefaultWeights = Weights{
	Delete:        1.0,
	Insert:        1.0,
	Match:         0.0,
	Replace:       1.0,
	Transpose:     0.5,
	PartialTrans1: 0.7,
	PartialTrans2: 0.7,
	LCS1:          0.3,
	LCS2:          0.3,
}

// Table4Weights are the weights from Table-4 of ./docs/papers/Soft Bigram distance for names matching.pdf
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

// Another popular configuration focusing on reducing false positives in name matching
var OptimizedWeights = Weights{
	Delete:        0.8,
	Insert:        0.8,
	Match:         0.0,
	Replace:       1.0,
	Transpose:     0.4,
	PartialTrans1: 0.5,
	PartialTrans2: 0.5,
	LCS1:          0.5,
	LCS2:          0.4,
}

// Configuration for LASA Drug Name Matching (FDA inspired)
var LASAWeights = Weights{
	Delete:        0.9,
	Insert:        0.9,
	Match:         0.0,
	Replace:       1.0,
	Transpose:     0.6,
	PartialTrans1: 0.7,
	PartialTrans2: 0.7,
	LCS1:          0.4,
	LCS2:          0.4,
}

// Configuration optimized for multilingual name comparison
var MultilingualWeights = Weights{
	Delete:        1.0,
	Insert:        1.0,
	Match:         0.0,
	Replace:       1.2, // Higher weight for replacements due to language variations
	Transpose:     0.3,
	PartialTrans1: 0.6,
	PartialTrans2: 0.6,
	LCS1:          0.5, // Emphasis on sequence similarity in different languages
	LCS2:          0.4,
}

// Configuration for High Precision Matching (focuses on minimizing false positives)
var HighPrecisionWeights = Weights{
	Delete:        0.8,
	Insert:        0.8,
	Match:         0.0,
	Replace:       1.5, // Higher penalty for replacements
	Transpose:     0.4,
	PartialTrans1: 0.6,
	PartialTrans2: 0.6,
	LCS1:          0.6, // Strong influence of LCS to maintain structure
	LCS2:          0.5,
}

// Configuration emphasizing phonetic similarity for medical applications
var PhoneticWeights = Weights{
	Delete:        1.0,
	Insert:        1.0,
	Match:         0.0,
	Replace:       0.7, // Lower weight for replacements due to phonetic leniency
	Transpose:     0.3, // Slight penalty for transpositions
	PartialTrans1: 0.5,
	PartialTrans2: 0.5,
	LCS1:          0.2, // Minimal LCS influence as phonetics are prioritized
	LCS2:          0.2,
}
