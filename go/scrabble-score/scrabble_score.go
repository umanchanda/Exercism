package scrabble

var scrabbleScores = map[byte]int{
	'A': 1,
	'B': 3,
	'a': 1,
	'f': 4,
	't': 1,
	'z': 10,
	'o': 1,
	's': 1,
	'r': 1,
	'e': 1,
	'q': 10,
	'u': 1,
	'c': 3,
	'y': 4,
	'O': 1,
	'i': 1,
	'k': 5,
	'b': 3,
	'd': 2,
	'g': 2,
	'p': 3,
	'n': 1,
	'x': 8,
	'h': 4,
	'j': 8,
	'l': 1,
	'm': 3,
	'v': 4,
	'w': 4,
}

// Score is a function that finds the scrabble score for a given word
func Score(word string) int {
	score := 0
	for index := 0; index < len(word); index++ {
		score += scrabbleScores[word[index]]
	}
	return score
}
