package advent

import (
	"testing"
	"fmt"
)

func testDiffLargeSmall(t *testing.T, record []string, expected int) {
	result := diffLargeSmall(record)
	if result != expected {
		t.Errorf("Fail : expected %d", expected)
	}
}

func TestRows(t *testing.T) {
	testDiffLargeSmall(t, []string{"5", "1", "9", "5"}, 8)
	testDiffLargeSmall(t, []string{"7", "5", "3"}, 4)
	testDiffLargeSmall(t, []string{"2", "4", "6", "8"}, 6)
}

func TestCalcChecksum(t *testing.T) {
	checksum := CalcChecksum("./dayTwoTestInput.tsv")
	if checksum != 18 {
		t.Errorf("Fail : expected %d", 18)
	}
}

func TestCalcChecksumTwo(t *testing.T) {
	checksum := CalcChecksumTwo("./dayTwoTestInputTwo.tsv")
	if checksum != 9 {
		t.Errorf("Fail : expected %d", 9)
	}
}

func TestCalcAdventChecksum(t *testing.T) {
	checksum := CalcChecksum("./dayTwoInput.tsv")
	fmt.Printf("%d", checksum)
	//if checksum != 18 {
	//	t.Errorf("Fail : expected %d", 18)
	//}
}

func TestCalcAdventChecksumTwo(t *testing.T) {
	checksum := CalcChecksumTwo("./dayTwoInput.tsv")
	fmt.Printf("%d", checksum)
	//if checksum != 18 {
	//	t.Errorf("Fail : expected %d", 18)
	//}
}

func TestReadRecords(t *testing.T) {
	records := readRecords("./dayTwoTestInput.tsv")
	for i := 0; i < len(records); i++ {
		fmt.Printf("%s : %d\n", records[i], len(records[i]))
	}

}