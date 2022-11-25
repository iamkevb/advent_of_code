package main

import (
	"fmt"
)

type Target struct {
	top, left, bottom, right int
}

var testTarget Target = Target{
	top:    -5,
	left:   20,
	bottom: -10,
	right:  30,
}

var realTarget Target = Target{
	top:    -108,
	bottom: -150,
	left:   81,
	right:  129,
}

func main() {
	target := realTarget

	// assuming positive vy, 2*vy steps gets you back to y zero.
	// if vx >= target bottom, your going to miss,
	// x   	y   vx  vy
	// 0   	0   6   10
	// 6   	10  5   9
	// 11  	19  4   8
	// 15  	27  3   7
	// 18  	34  2   6
	// 20  	40  1   5
	// 21  	45  0   4
	// 21  	49  0   3
	// 21  	52  0   2
	// 21  	54  0   1
	// 21  	55  0   0
	// 		55      -1
	// 		54      -2
	// 		52      -3
	// 		49      -4
	// 		45      -5
	// 		40      -6
	// 		34      -7
	// 		27      -8
	// 		19      -9
	// 		10      -10
	// 		0       -11
	// 		-11     -12  -- missed -5..-10

	// vy must be less than abs(target.bottom)
	vy := target.bottom*-1 - 1

	// because the vy increment is -1, the max height is the sum of digits between 0 and vy
	maxHeight := vy * (vy + 1) / 2

	fmt.Println(maxHeight)
}
