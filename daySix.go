package advent

import (
    "strconv"
)

func BalanceBlocks(filename string) int {

    records := ReadTsvRecordsToInt(filename)
    blocks := records[0]
    numBanks := len(blocks)
    bankList := make([]int, numBanks)
    for i := 0; i < len(blocks); i++ {
        bankList[i] = blocks[i]
    }
    gridMap[flattenBlocks(blocks)] = 1;

    done := 0
    steps := 0
    collisionStep := 0
    loopSteps := 0
    loopBeginString := ""

    for done < 1 {
        steps++
        largestBlockIndex := -1
        largestBlock := 0
        for i := 0; i < len(blocks); i++ {
            if blocks[i] > largestBlock {
              largestBlock = blocks[i]
              largestBlockIndex = i
            }
        }

        blocksToAssign := largestBlock;
        blocks[largestBlockIndex] = 0
        for blocksToAssign > 0 {
            for j := largestBlockIndex + 1; j < len(blocks); j++ {
                blocks[j] = blocks[j] + 1
                blocksToAssign --;
                if blocksToAssign == 0 {break}
            }

            if blocksToAssign == 0 {break}

            for k := 0; k <= largestBlockIndex; k++ {
                blocks[k] = blocks[k] + 1
                blocksToAssign --;
                if blocksToAssign == 0 {break}
            }
        }

        flattened := flattenBlocks(blocks)
        if gridMap[flattened] == 1 {
            if (collisionStep == 0) {
                collisionStep = steps;
                loopBeginString = flattened;
            } else {
                if flattened == loopBeginString {
                    loopSteps = steps - collisionStep;
                    done = 1;
                }
            }
        } else {
            gridMap[flattened] = 1
        }

    }
    return loopSteps;

}

func flattenBlocks(blocks []int) string {
    allBlocks := ""
    for i := 0; i < len(blocks); i++ {
        allBlocks += strconv.Itoa(blocks[i])
        allBlocks += "-"
    }
    return allBlocks
}
