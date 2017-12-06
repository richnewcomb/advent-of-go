package advent

import (
    "testing"
    "fmt"
)

func TestCountValidPassphrases(t *testing.T) {
    validPhrases := CountValidPassphrases("./data/dayFourTestPassphrases")
    fmt.Printf("%d", validPhrases)
}

func TestCountValidAdventPassphrases(t *testing.T) {
    validPhrases := CountValidPassphrases("./data/dayFourPassphrases")
    fmt.Printf("%d", validPhrases)
}

func TestCountValidAdventPassphraseAnagrams(t *testing.T) {
    validPhrases := CountValidPassphraseAnagrams("./data/dayFourPassphrases")
    fmt.Printf("%d", validPhrases)
}