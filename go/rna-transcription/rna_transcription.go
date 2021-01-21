package strand

var transform = map[rune]rune {
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns a rna sequence given a DNA sequence
func ToRNA(dna string) string {
	rna := ""
	for _, c := range dna {
		rna += string(transform[c])
	}
	return string(rna)
}
