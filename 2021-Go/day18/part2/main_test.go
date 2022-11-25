package main

import (
	"fmt"
	"testing"
)

func assertNode(t *testing.T, node Node, value, depth int) {
	if node.value != value {
		t.Errorf("Expected value %d, got %d\n", value, node.value)
	}
	if node.depth != depth {
		t.Errorf("Expected depth %d, got %d\n", depth, node.depth)
	}
}
func TestParseNumber(t *testing.T) {
	sfn := parseNode("[0,[1,[2,3],[4,5],6],7]")
	if len(sfn) != 8 {
		t.Fatal("expected 8 nodes")
	}
	assertNode(t, sfn[0], 0, 1)
	assertNode(t, sfn[1], 1, 2)
	assertNode(t, sfn[2], 2, 3)
	assertNode(t, sfn[3], 3, 3)
	assertNode(t, sfn[4], 4, 3)
	assertNode(t, sfn[5], 5, 3)
	assertNode(t, sfn[6], 6, 2)
	assertNode(t, sfn[7], 7, 1)
}

func TestExplode981234(t *testing.T) {
	sfn := parseNode("[[[[[9,8],1],2],3],4]")
	r, _ := explode(sfn)

	//becomes [[[[0,9],2],3],4]
	if len(r) != 5 {
		t.Fatalf("expected len 5, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 0, 4)
	assertNode(t, r[1], 9, 4)
	assertNode(t, r[2], 2, 3)
	assertNode(t, r[3], 3, 2)
	assertNode(t, r[4], 4, 1)
}

func TestExplode654321(t *testing.T) {
	sfn := parseNode("[[6,[5,[4,[3,2]]]],1]")
	r, _ := explode(sfn)
	//becomes [[6,[5,[7,0]]],3]
	if len(r) != 5 {
		t.Fatalf("expected len 5, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 6, 2)
	assertNode(t, r[1], 5, 3)
	assertNode(t, r[2], 7, 4)
	assertNode(t, r[3], 0, 4)
	assertNode(t, r[4], 3, 1)
}

func TestExplode328095432(t *testing.T) {
	sfn := parseNode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	r, _ := explode(sfn)
	//becomes [[3,[2,[8,0]]],[9,[5,[7,0]]]]
	if len(r) != 8 {
		t.Fatalf("expected len 8, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 3, 2)
	assertNode(t, r[1], 2, 3)
	assertNode(t, r[2], 8, 4)
	assertNode(t, r[3], 0, 4)
	assertNode(t, r[4], 9, 2)
	assertNode(t, r[5], 5, 3)
	assertNode(t, r[6], 7, 4)
	assertNode(t, r[7], 0, 4)
}

func TestSplitEven(t *testing.T) {
	sfn := []Node{
		{value: 10, depth: 1},
		{value: 0, depth: 1},
	}
	r, _ := split(sfn)
	if len(r) != 3 {
		t.Fatalf("expected len 3, got %d %+v", len(r), r)
	}

	assertNode(t, r[0], 5, 2)
	assertNode(t, r[1], 5, 2)
	assertNode(t, r[2], 0, 1)
}
func TestSplitOdd(t *testing.T) {
	sfn := []Node{
		{value: 11, depth: 1},
		{value: 0, depth: 1},
	}
	r, _ := split(sfn)
	if len(r) != 3 {
		t.Fatalf("expected len 3, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 5, 2)
	assertNode(t, r[1], 6, 2)
	assertNode(t, r[2], 0, 1)
}

func TestAddSimple(t *testing.T) {
	sfn1 := parseNode("[0,1]")
	sfn2 := parseNode("[2,3]")
	sum := add(sfn1, sfn2)
	//becomes [[0,1],[2,3]]
	if len(sum) != 4 {
		t.Fatal("expected 4 nodes in sum.")
	}
	for i, v := range sum {
		if v.value != i {
			t.Errorf("expected %d, got %d", i, v.value)
		}
		if v.depth != 2 {
			t.Errorf("expected 2 depth, got %d", v.depth)
		}
	}
}

func TestAdd(t *testing.T) {
	sfn1 := parseNode("[[[[4,3],4],4],[7,[[8,4],9]]]")
	sfn2 := parseNode("[1,1]")
	r := add(sfn1, sfn2)

	//becomes [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
	if len(r) != 9 {
		t.Fatalf("expected len 9, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 0, 4)
	assertNode(t, r[1], 7, 4)
	assertNode(t, r[2], 4, 3)
	assertNode(t, r[3], 7, 4)
	assertNode(t, r[4], 8, 4)
	assertNode(t, r[5], 6, 4)
	assertNode(t, r[6], 0, 4)
	assertNode(t, r[7], 8, 2)
	assertNode(t, r[8], 1, 2)
}

func TestAddMultiple3(t *testing.T) {
	sfns := make([][]Node, 6)
	sfns[0] = parseNode("[1,1]")
	sfns[1] = parseNode("[2,2]")
	sfns[2] = parseNode("[3,3]")

	sum := add(sfns[0], sfns[1])
	sum = add(sum, sfns[2])
	//becomes [[[1,1],[2,2]],[3,3]]
	if len(sum) != 6 {
		t.Fatalf("expected len 6, got %d %+v", len(sum), sum)
	}
	assertNode(t, sum[0], 1, 3)
	assertNode(t, sum[1], 1, 3)
	assertNode(t, sum[2], 2, 3)
	assertNode(t, sum[3], 2, 3)
	assertNode(t, sum[4], 3, 2)
	assertNode(t, sum[5], 3, 2)
}

func TestAddMultiple4(t *testing.T) {
	sfns := make([][]Node, 4)
	sfns[0] = parseNode("[1,1]")
	sfns[1] = parseNode("[2,2]")
	sfns[2] = parseNode("[3,3]")
	sfns[3] = parseNode("[4,4]")
	var sum []Node = sfns[0]
	for i := 1; i < len(sfns); i++ {
		sum = add(sum, sfns[i])
	}
	//becomes [[[[1,1],[2,2]],[3,3]],[4,4]]
	if len(sum) != 8 {
		t.Fatalf("expected len 8, got %d %+v", len(sum), sum)
	}
	assertNode(t, sum[0], 1, 4)
	assertNode(t, sum[1], 1, 4)
	assertNode(t, sum[2], 2, 4)
	assertNode(t, sum[3], 2, 4)
	assertNode(t, sum[4], 3, 3)
	assertNode(t, sum[5], 3, 3)
	assertNode(t, sum[6], 4, 2)
	assertNode(t, sum[7], 4, 2)
}

func TestAddMultiple6(t *testing.T) {
	sfns := make([][]Node, 6)
	sfns[0] = parseNode("[1,1]")
	sfns[1] = parseNode("[2,2]")
	sfns[2] = parseNode("[3,3]")
	sfns[3] = parseNode("[4,4]")
	sfns[4] = parseNode("[5,5]")
	sfns[5] = parseNode("[6,6]")

	var sum []Node = sfns[0]
	for i := 1; i < len(sfns); i++ {
		sum = add(sum, sfns[i])
	}

	//becomes [[[[5,0],[7,4]],[5,5]],[6,6]]
	if len(sum) != 8 {
		t.Fatalf("expected len 8, got %d %+v", len(sum), sum)
	}
	assertNode(t, sum[0], 5, 4)
	assertNode(t, sum[1], 0, 4)
	assertNode(t, sum[2], 7, 4)
	assertNode(t, sum[3], 4, 4)
	assertNode(t, sum[4], 5, 3)
	assertNode(t, sum[5], 5, 3)
	assertNode(t, sum[6], 6, 2)
	assertNode(t, sum[7], 6, 2)
}

func TestReduceRecursive(t *testing.T) {
	sfn := parseNode("[[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]],[6,6]]")
	r := reduce(sfn)
	for _, n := range r {
		if n.depth > 4 {
			t.Errorf("Depth should never be more than 4, depth=%d", n.depth)
		}
		if n.value > 9 {
			t.Errorf("Value should never be more than 9, depth=%d", n.value)
		}
	}
}

func TestMagnitude(t *testing.T) {
	n := parseNode("[[1,2],[[3,4],5]]")
	mag := magnitude(n)
	if mag != 143 {
		t.Errorf("Expected 143, got %d", mag)
	}

	n = parseNode("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	mag = magnitude(n)
	if mag != 1384 {
		t.Errorf("Expected 1384, got %d", mag)
	}

	n = parseNode("[[[[1,1],[2,2]],[3,3]],[4,4]]")
	mag = magnitude(n)
	if mag != 445 {
		t.Errorf("Expected 445, got %d", mag)
	}

	n = parseNode("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	mag = magnitude(n)
	if mag != 3488 {
		t.Errorf("Expected 3488, got %d", mag)
	}
}

func TestLargerAdd(t *testing.T) {
	n := make([][]Node, 10)
	n[0] = parseNode("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
	n[1] = parseNode("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]")
	n[2] = parseNode("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]")
	n[3] = parseNode("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]")
	n[4] = parseNode("[7,[5,[[3,8],[1,4]]]]")
	n[5] = parseNode("[[2,[2,2]],[8,[8,1]]]")
	n[6] = parseNode("[2,9]")
	n[7] = parseNode("[1,[[[9,3],9],[[9,0],[0,7]]]]")
	n[8] = parseNode("[[[5,[7,4]],7],1]")
	n[9] = parseNode("[[[[4,2],2],6],[8,7]]")

	var sum []Node = n[0]
	for i := 1; i < len(n); i++ {
		sum = add(sum, n[i])
	}
	// becomes [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]
	// if len(sum) != 14 {
	// 	t.Fatalf("Expected 14 nodes, got %d", len(sum))
	// }
	for _, n := range sum {
		if n.depth > 4 {
			t.Errorf("Depth should never be more than 4, depth=%d", n.depth)
		}
		if n.value > 9 {
			t.Errorf("Value should never be more than 9, depth=%d", n.value)
		}
	}
}

func TestExplode4344784911(t *testing.T) {
	sfn := parseNode("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	r, _ := explode(sfn)
	// becomes [[[[0,7],4],[7,[[8,4],9]]],[1,1]]
	if len(r) != 9 {
		t.Fatalf("expected len 9, got %d %+v", len(r), r)
	}

	assertNode(t, r[0], 0, 4)
	assertNode(t, r[1], 7, 4)
	assertNode(t, r[2], 4, 3)
	assertNode(t, r[3], 7, 3)
	assertNode(t, r[4], 8, 5)
	assertNode(t, r[5], 4, 5)
	assertNode(t, r[6], 9, 4)
	assertNode(t, r[7], 1, 2)
	assertNode(t, r[8], 1, 2)
}

func TestExplode074784911(t *testing.T) {
	sfn := parseNode("[[[[0,7],4],[7,[[8,4],9]]],[1,1]]")
	r, _ := explode(sfn)

	// becomes [[[[0,7],4],[15,[0,13]]],[1,1]]
	if len(r) != 8 {
		t.Fatalf("expected len 8, got %d %+v", len(r), r)
	}
	assertNode(t, r[0], 0, 4)
	assertNode(t, r[1], 7, 4)
	assertNode(t, r[2], 4, 3)
	assertNode(t, r[3], 15, 3)
	assertNode(t, r[4], 0, 4)
	assertNode(t, r[5], 13, 4)
	assertNode(t, r[6], 1, 2)
	assertNode(t, r[7], 1, 2)
}

func TestSplit0741501311(t *testing.T) {
	sfn := []Node{
		{value: 0, depth: 4},
		{value: 7, depth: 4},
		{value: 4, depth: 3},
		{value: 15, depth: 3},
		{value: 0, depth: 4},
		{value: 13, depth: 4},
		{value: 1, depth: 2},
		{value: 1, depth: 2},
	}
	r, _ := split(sfn)
	//becomes [[[[0,7],4],[[7,8],[0,13]]],[1,1]]
	if len(r) != 9 {
		t.Fatalf("expected len 9, got %d", len(r))
	}
	assertNode(t, r[0], 0, 4)
	assertNode(t, r[1], 7, 4)
	assertNode(t, r[2], 4, 3)
	assertNode(t, r[3], 7, 4)
	assertNode(t, r[4], 8, 4)
	assertNode(t, r[5], 0, 4)
	assertNode(t, r[6], 13, 4)
	assertNode(t, r[7], 1, 2)
	assertNode(t, r[8], 1, 2)
}

func TestMagnitudeWasPanic(t *testing.T) {
	n := []Node{
		{value: 7, depth: 4},
		{value: 7, depth: 4},
		{value: 6, depth: 4},
		{value: 7, depth: 4},
		{value: 0, depth: 4},
		{value: 7, depth: 4},
		{value: 7, depth: 4},
		{value: 7, depth: 4},
		{value: 7, depth: 4},
		{value: 7, depth: 4},
		{value: 8, depth: 4},
		{value: 8, depth: 4},
		{value: 9, depth: 4},
	}
	m := magnitude(n)
	if m != 100 {
		t.Errorf("expected 100, got %d", m)
	}
}

func TestAddCausingPanic(t *testing.T) {
	n1 := parseNode("[[9,3],[[9,9],[6,[4,9]]]]")
	n2 := parseNode("[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]")
	s := add(n1, n2)
	if len(s) != 15 {
		t.Fatalf("expected len 15, got %d", len(s))
	}
	fmt.Printf("%+v\n", s)
	magnitude(s)
}
