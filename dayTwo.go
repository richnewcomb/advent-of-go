package advent

func CalcChecksum(filename string) int {
	records := ReadTsvRecordsToInt(filename)
	checksum := 0
	for i := 0; i < len(records); i++ {
		checksum += diffLargeSmall(records[i])
	}
	return checksum
}

func CalcChecksumTwo(filename string) int {
	records := ReadTsvRecordsToInt(filename)
	checksum := 0
	for i := 0; i < len(records); i++ {
		checksum += resultOfEventlyDivisible(records[i])
	}
	return checksum
}

// Return the difference between the largest and smallest converted integers in a record
func diffLargeSmall(record []int) int {
	smallest := 65535 // max uint16
	largest := 0      // assuming positive integers
	position := 0
	currentValue := 0

	for position < len(record) {
		currentValue = record[position]
		if currentValue < smallest {
			smallest = currentValue
		}
		if currentValue > largest {
			largest = currentValue
		}
		position++
	}
	return largest - smallest
}

func resultOfEventlyDivisible(record []int) int {
	for i := 0; i < len(record); i++ {
		for j := 0; j < len(record); j++ {
			if i != j {
				numerator := record[i]
				denominator := record[j]

				if numerator%denominator == 0 {
					return numerator / denominator
				}
			}
		}
	}
	return -1
}


