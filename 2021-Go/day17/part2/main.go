package main

import (
	"fmt"
)

// sub is at 0,0
// x changes by 1, towards 0
// y decreases 1, no limit
// goal: get probe into target area with max y

type Point struct {
	x, y int
}

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

type Velocity struct {
	x, y int
}

func main() {
	target := realTarget

	// if you one step, this is the hardest you can go and still hit
	maxVx := target.right
	//goes up maxVy, down to zero, then to -maxVy. Any larger than bottom and you miss
	maxVy := target.bottom*-1 - 1

	// lowest x velocity to make it to target
	minVx := computeMinVx(target.left)
	// most direct route (lowest possible value)
	minVy := target.bottom

	hitCount := 0

	for x := minVx; x <= maxVx; x++ {
		for y := minVy; y <= maxVy; y++ {
			v := Velocity{x, y}
			if hitTest(v, target) {
				hitCount++
			}
		}
	}
	fmt.Printf("hits %d\n", hitCount)
}

func hitTest(v Velocity, target Target) bool {
	x, y := 0, 0
	for x <= target.right && y >= target.bottom {
		if x >= target.left && x <= target.right && y >= target.bottom && y <= target.top {
			return true
		}
		x, y, v = step(x, y, v)
	}
	return false
}

func step(x, y int, v Velocity) (int, int, Velocity) {
	x += v.x
	y += v.y
	if v.x > 0 {
		v.x = max(0, v.x-1)
	} else {
		v.x = min(0, v.x+1)
	}
	v.y -= 1
	return x, y, v
}

func computeMinVx(d int) int {
	// have to launch with enough to get to d
	// d = vx * steps - (steps(steps-1)/2)
	// maximum d when steps = vx (after that vx is 0 forever)
	// d = vx * vx - (vx(vx-1))/2
	vx := 1
	for vx < d {
		if vx*vx-(vx*(vx-1))/2 >= d {
			return vx
		}
		vx++
	}
	panic("math is hard")
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
