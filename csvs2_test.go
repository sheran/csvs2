package csvs2

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	place1 := NewPointFromLL(25.2048493, 55.2707828)
	place2 := NewPointFromLL(25.1648925, 55.4084034)
	q := place1.DistanceFrom(*place2)
	fmt.Println(q)
}
