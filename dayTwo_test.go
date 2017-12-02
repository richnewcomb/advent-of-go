package advent

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func testDiffLargeSmall(t *testing.T, record []int, expected int) {
	result := diffLargeSmall(record)
	assert.Equal(t, expected, result, "Result does not match expected.")
}

func TestDiffLargeSmallRows(t *testing.T) {
	testDiffLargeSmall(t, []int{5, 1, 9, 5}, 8)
	testDiffLargeSmall(t, []int{7, 5, 3}, 4)
	testDiffLargeSmall(t, []int{2, 4, 6, 8}, 6)
}

func TestCalcChecksum(t *testing.T) {
	checksum := CalcChecksum("./data/dayTwoTestInput.tsv")
	assert.Equal(t, 18, checksum, "Result does not match expected.")
}

func TestCalcChecksumTwo(t *testing.T) {
	checksum := CalcChecksumTwo("./data/dayTwoTestInputTwo.tsv")
	assert.Equal(t, 9, checksum, "Result does not match expected.")
}

func TestCalcAdventChecksum(t *testing.T) {
	checksum := CalcChecksum("./data/dayTwoInput.tsv")
	fmt.Printf("%d", checksum)
}

func TestCalcAdventChecksumTwo(t *testing.T) {
	checksum := CalcChecksumTwo("./data/dayTwoInput.tsv")
	assert.Equal(t, 226, checksum, "Result does not match expected.")
}


