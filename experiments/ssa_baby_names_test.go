package experiments

import (
	"testing"

	"github.com/PhonoGrams/soft_bigram"
)

// TestSSABabyNames is a comparison of SSA.gov popular baby names from 2023
// Source: https://www.ssa.gov/oact/babynames/
func TestSSABabyNames(t *testing.T) {
	weights := make(map[string]soft_bigram.Weights)
	weights["default"] = soft_bigram.DefaultWeights
	weights["table4"] = soft_bigram.Table4Weights
	weights["optimized"] = soft_bigram.OptimizedWeights
	weights["high precision"] = soft_bigram.HighPrecisionWeights
	weights["phonetic"] = soft_bigram.PhoneticWeights

	t.Run("male names", func(t *testing.T) {
		pairs := [][]string{
			{"Liam", "Lyle", "Lian", "Lio"},
			{"Noah", "Nolan", "Noel", "Noe"},
			{"Oliver", "Olivier", "Ollie", "Olly"},
			{"James", "Jamie", "Jay", "Jim"},
			{"Elijah", "Elias", "Eli", "Elie"},
			{"Mateo", "Matias", "Matt", "Matteo"},
			{"Theodore", "Theo", "Teddy", "Ted"},
			{"Henry", "Harry", "Hank", "Henrik"},
			{"Lucas", "Luca", "Luke", "Lucian"},
			{"William", "Will", "Bill", "Liam"},
			{"Benjamin", "Ben", "Bennett", "Benny"},
			{"Levi", "Lev", "Levon", "Lee"},
			{"Sebastian", "Bastian", "Seb", "Basti"},
			{"Jack", "Jackson", "Jackie", "Jax"},
			{"Ezra", "Eli", "Ezer", "Azra"},
			{"Michael", "Micah", "Mick", "Mike"},
			{"Daniel", "Dan", "Danny", "Dani"},
			{"Leo", "Leon", "Leonard", "Leonardo"},
			{"Owen", "Ewan", "Owain", "Odin"},
			{"Samuel", "Sam", "Sammy", "Sami"},
		}
		runExperiments(t, weights, pairs)
	})

	t.Run("female names", func(t *testing.T) {
		pairs := [][]string{
			{"Olivia", "Olive", "Liv", "Livia"},
			{"Emma", "Emmy", "Emilia", "Em"},
			{"Charlotte", "Charlie", "Lottie", "Carlotta"},
			{"Amelia", "Amy", "Mia", "Emilia"},
			{"Sophia", "Sophie", "Sofie", "Sofia"},
			{"Mia", "Maya", "Amia", "Mina"},
			{"Isabella", "Isabelle", "Bella", "Izzy"},
			{"Ava", "Eva", "Avie", "Avia"},
			{"Evelyn", "Eve", "Evie", "Lyn"},
			{"Luna", "Lu", "Lunette", "Luz"},
			{"Harper", "Harley", "Harlow", "Perry"},
			{"Sofia", "Sophie", "Sophia", "Sofie"},
			{"Camila", "Cami", "Mila", "Millie"},
			{"Eleanor", "Ellie", "Nora", "Lenore"},
			{"Elizabeth", "Eliza", "Liz", "Beth"},
			{"Violet", "Vi", "Lettie", "Viola"},
			{"Scarlett", "Scar", "Letty", "Scarlet"},
			{"Emily", "Emma", "Emmy", "Millie"},
			{"Hazel", "Haze", "Zel", "Haven"},
			{"Lily", "Lila", "Lilith", "Lilian"},
		}
		runExperiments(t, weights, pairs)
	})
}
