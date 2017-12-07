package advent

import (
    "testing"
    "fmt"
    "github.com/stretchr/testify/assert"
)

func TestBalanceBlocks(t *testing.T) {
    steps := BalanceBlocks("./data/daySixTestInput.tsv")
    assert.Equal(t, 5, steps)
}

func TestBalanceBlocksLoop(t *testing.T) {
    steps := BalanceBlocks("./data/daySixTestInputTwo.tsv")
    assert.Equal(t, 4, steps)
}

func TestAdventBalanceBlocks(t *testing.T) {
    steps := BalanceBlocks("./data/daySixInput.tsv")
    fmt.Printf("%d", steps)
}
