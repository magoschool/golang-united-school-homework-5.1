package square

import (
	"testing"
)

func newSquare(aX int, aY int, aA uint) Square {
	return Square{
		start: Point{x: aX, y: aY},
		a:     aA,
	}
}

func TestEnd(t *testing.T) {
	lSquare := newSquare(2, 3, 4)
	lTest := lSquare.End()
	lExpected := Point{x: 6, y: 7}

	if lTest.x != lExpected.x || lTest.y != lExpected.y {
		t.Errorf("Square.End() failed: %v expected but %v found", lExpected, lTest)
	}
}

func TestArea(t *testing.T) {
	lSquare := newSquare(2, 3, 4)
	lTest := lSquare.Area()
	var lExpected uint = 16

	if lTest != lExpected {
		t.Errorf("Square.Area() failed: %v expected but %v found", lExpected, lTest)
	}
}

func TestPerimeter(t *testing.T) {
	lSquare := newSquare(3, 4, 5)
	lTest := lSquare.Perimeter()
	var lExpected uint = 20

	if lTest != lExpected {
		t.Errorf("Square.Perimeter() failed: %v expected but %v found", lExpected, lTest)
	}
}
