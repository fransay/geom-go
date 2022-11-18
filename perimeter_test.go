package shapelib

import (
	"math"
	functs "shapelib/functs/2D"
	"testing"
)

func TestPerimeterOfTriangle(t *testing.T) {
	res := functs.PerimOfRTriangle(10.0, 10.0, 10.0)
	exp := math.Trunc(30)
	if math.Trunc(res) != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}
}

func TestPerimeterOfSquare(t *testing.T) {
	res := functs.PerimOfRSquare(5.0)
	exp := math.Trunc(20.0)
	if res != exp {
		t.Errorf("expected %f, got %f", exp, res)
	}

}
