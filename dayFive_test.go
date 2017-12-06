package advent

import (
    "testing"
    "fmt"
)

func TestEscapeMaze(t *testing.T) {
    steps := EscapeMaze("./data/dayFiveInstructions")
    fmt.Printf("%d", steps)
}
