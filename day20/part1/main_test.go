package main

import (
	"testing"
)

var testAlgo []byte

func fillAlgo() {
	algo = make([]rune, 512)
	algo[0] = '.'
	for i := range algo {
		copy(algo[i:], algo[:i])
	}
}

func TestMinMaxXY(t *testing.T) {
	img := map[Pixel]bool{
		{x: 0, y: 0}:   true,
		{x: 100, y: 3}: true,
		{x: 2, y: -11}: true,
		{x: 4, y: 2}:   true,
		{x: 1, y: 0}:   true,
		{x: 2, y: 0}:   true,
		{x: -20, y: 4}: true,
		{x: 3, y: 2}:   true,
		{x: 4, y: 3}:   true,
		{x: 4, y: 400}: true,
	}
	min, max := minMaxXY(img)
	if min.x != -20 || min.y != -11 {
		t.Errorf("expected {-20,-11}, got %v", min)
	}
	if max.x != 100 || max.y != 400 {
		t.Errorf("expected {100,400}, got %v", max)
	}
}

func TestEnhancePixel(t *testing.T) {
	fillAlgo()
	algo[16] = '#'
	img := map[Pixel]bool{
		{1, 1}: true,
	}
	//binary string 000010000  16

	if !enhancePixel(1, 1, img) {
		t.Error("expected 1,1 to return true")
	}
}

func TestEnhancePixel273(t *testing.T) {
	fillAlgo()
	algo[280] = '#'
	img := map[Pixel]bool{
		{0, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
	}
	//binary string 100 011 000  280

	if !enhancePixel(1, 1, img) {
		t.Error("expected 1,1 to return true")
	}
}

func TestEnhance(t *testing.T) {
	fillAlgo()
	algo[0] = '#'
	img := map[Pixel]bool{
		{0, 0}: true,
		{0, 1}: true,
		{1, 0}: true,
		{1, 1}: true,
		{2, 0}: true,
		{2, 2}: true,
	}
	enhanced := enhance(img)
	enhanced = enhance(enhanced)
	writeImg(enhanced)
	for _, v := range enhanced {
		if v {
			t.Error("fail")
		}
	}
}
