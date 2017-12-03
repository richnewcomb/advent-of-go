package advent

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDiffLargeSmallRows(t *testing.T) {
	assert.Equal(t, 8, diffLargeSmall([]int{5, 1, 9, 5}))
	assert.Equal(t, 4, diffLargeSmall([]int{7,5,3}))
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
	fmt.Printf("%d", checksum)
}


