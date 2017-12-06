package advent

//import "fmt"

func EscapeMaze(filename string) int {

//    instructionList := list.New()

    records := ReadTsvRecordsToInt(filename)
    numInstructions := len(records)
    instructionList := make([]int, numInstructions)
    for i := 0; i < len(records); i++ {
        instructionList[i] = records[i][0]
    }

    done := 1
    steps := 0
    pos := 0;
    for done < 2 {
        currentVal := instructionList[pos]
        nextPos := pos + currentVal
        if (currentVal >= 3) {
            instructionList[pos] = currentVal -1
        } else {
            instructionList[pos] = currentVal + 1
        }

        pos = nextPos
        steps++
        if (pos < 0 || pos > len(instructionList) -1) {
            break
        }
    }

    return steps;
}
