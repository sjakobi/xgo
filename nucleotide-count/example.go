package dna

import (
	"errors"
	"strings"
)

type Histogram map[byte]int

type DNA string

func (dna DNA) Count(nucleotide byte) (count int, err error) {
	validNucleotides := "ACGT"
	if strings.IndexByte(validNucleotides, nucleotide) < 0 {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}

	return strings.Count(string(dna), string(nucleotide)), nil
}

func (dna DNA) Counts() Histogram {
	a, _ := dna.Count('A')
	c, _ := dna.Count('C')
	t, _ := dna.Count('T')
	g, _ := dna.Count('G')

	return Histogram{'A': a, 'C': c, 'T': t, 'G': g}
}
