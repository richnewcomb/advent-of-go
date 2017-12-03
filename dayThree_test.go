package advent

import (
	"testing"
	"fmt"
)

func TestCalcManhattanDistance(t *testing.T) {
	distance := CalcManhattanDistance(265149)
	fmt.Printf("%d", distance)
}

func TestSumAdjacentLocations(t *testing.T) {
	sumOfAdjacentLocations := SumAdjacentLocations(265149)
	fmt.Printf("%d", sumOfAdjacentLocations)

}