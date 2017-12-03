package advent

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

const INPUT_FILE = "./data/dayTwoTestInput.tsv"

func TestReadRecords(t *testing.T) {
    records := ReadTsvRecords(INPUT_FILE)
    assert.Equal(t, 3, len(records))
    assert.Equal(t, 4, len(records[0]))
    assert.Equal(t, 3, len(records[1]))
    assert.Equal(t, "8", records[2][3])
}

func TestReadTsvRecordsToInt(t *testing.T) {
    records := ReadTsvRecordsToInt(INPUT_FILE)
    assert.Equal(t, 3, len(records))
    assert.Equal(t, 4, len(records[0]))
    assert.Equal(t, 3, len(records[1]))
    assert.Equal(t, 8, records[2][3])
}
