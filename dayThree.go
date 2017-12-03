package advent

import (
	"strconv"
)

var gridMap = make(map[string]int)

func CalcManhattanDistance(location int) int {
	i := 1;
	ring := 0;
	for i < location {
		ring++;
		for side := 0; side < 4; side++ {
			sideLen := ring * 2;
		    for j := sideLen; j > sideLen / 2; j-- {
		    	i++
				if i == location { return j -1 }
			}
			for k := 1; k <= sideLen / 2; k++ {
				i++
				if i == location { return k + ring }
			}
		}
	}
	return -1;
}

func SumAdjacentLocations(minTotal int) int {

	gridMap[strconv.Itoa(0) + "," + strconv.Itoa(0)] = 1

	value := 0;
	x := 0;
	y := -1;

	for value < minTotal {
		x++;
		sideLen := x * 2;
		for i := 0; i < sideLen; i++ {
			y++;
			value = sumAdjacentLocations(x, y);
			if value > minTotal {return value}
		}
		for i := 0; i < sideLen; i++ {
			x--;
			value = sumAdjacentLocations(x, y);
			if value > minTotal {return value}
		}
		for i := 0; i < sideLen; i++ {
			y--;
			value = sumAdjacentLocations(x, y);
			if value > minTotal {return value}
		}
		for i := 0; i < sideLen; i++ {
			x++;
			value = sumAdjacentLocations(x, y);
			if value > minTotal {return value}
		}
		y--;
	}

	return -1;
}

func sumAdjacentLocations(x int, y int) int {
	sum :=
		gridMap[strconv.Itoa(x) + "," + strconv.Itoa(y + 1)] +
		gridMap[strconv.Itoa(x) + "," + strconv.Itoa(y - 1)] +
		gridMap[strconv.Itoa(x + 1) + "," + strconv.Itoa(y)] +
		gridMap[strconv.Itoa(x - 1) + "," + strconv.Itoa(y)] +
	    gridMap[strconv.Itoa(x + 1) + "," + strconv.Itoa(y + 1)] +
	    gridMap[strconv.Itoa(x - 1) + "," + strconv.Itoa(y - 1)] +
	    gridMap[strconv.Itoa(x + 1) + "," + strconv.Itoa(y - 1)] +
		gridMap[strconv.Itoa(x - 1) + "," + strconv.Itoa(y + 1)]

	gridMap[strconv.Itoa(x) + "," + strconv.Itoa(y)] = sum

	return sum;
}
