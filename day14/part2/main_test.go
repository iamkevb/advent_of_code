package main

import (
	"testing"
)

func setup() {
	elements = map[rune]int{}
	polymer = "NNCB"
	rules = map[string]rune{
		"CH": 'B',
		"HH": 'N',
		"CB": 'H',
		"NH": 'C',
		"HB": 'C',
		"HC": 'B',
		"HN": 'C',
		"NN": 'C',
		"BH": 'H',
		"NC": 'B',
		"NB": 'B',
		"BN": 'B',
		"BB": 'N',
		"BC": 'B',
		"CC": 'N',
		"CN": 'C',
	}
}

func TestMakePairs(t *testing.T) {
	setup()
	pairs := makePairs(polymer)
	if len(pairs) != 3 {
		t.Errorf("Expected 3 pairs, got %d", len(pairs))
	}
	if pairs["NN"] != 1 ||
		pairs["NC"] != 1 ||
		pairs["CB"] != 1 {
		t.Error("Expected NN, NC, CB pairs")
	}
	if elements['N'] != 2 || elements['C'] != 1 || elements['B'] != 1 {
		t.Error("elements is wrong")
	}
}

func TestGrowPairs(t *testing.T) {
	setup()
	pairs := makePairs(polymer)

	pairs = growPairs(pairs)

	if pairs["NC"] != 1 ||
		pairs["CN"] != 1 ||
		pairs["NB"] != 1 ||
		pairs["BC"] != 1 ||
		pairs["CH"] != 1 ||
		pairs["HB"] != 1 {
		t.Error("Expected NC CN NB BC CH HB pairs")
	}

	if elements['N'] != 2 || elements['C'] != 2 || elements['B'] != 2 || elements['H'] != 1 {
		t.Error("elements is wrong")
		t.Error(elements)
	}

	pairs = growPairs(pairs)

	if pairs["NB"] != 2 ||
		pairs["BC"] != 2 ||
		pairs["CC"] != 1 ||
		pairs["CN"] != 1 ||
		pairs["BB"] != 2 ||
		pairs["CB"] != 2 ||
		pairs["BH"] != 1 ||
		pairs["HC"] != 1 {
		t.Error("Expected NB BC CC CN NB BB BB BC CB BH HC CB pairs")
	}

	if elements['N'] != 2 || elements['C'] != 4 || elements['B'] != 6 || elements['H'] != 1 {
		t.Error("elements is wrong")
		t.Error(elements)
	}
}
