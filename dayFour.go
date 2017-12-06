package advent

import (
    "strings"
    "sort"
)

// var gridMap = make(map[string]int)

func CountValidPassphrases(filename string) int {
    records := ReadTsvRecordsWithDelimiter(filename, ' ')
    goodPhraseCount := 0;
    for i := 0; i < len(records); i++ {
        phrase := records[i]
        strMap := make(map[string]int)
        goodPhrase := true;
        for j:=0; j < len(phrase); j++ {
            if strMap[phrase[j]] == 1 {
                goodPhrase = false;
                break;
            }
            strMap[phrase[j]] = 1
        }
        if goodPhrase {
            goodPhraseCount ++;
        }
    }

    return goodPhraseCount;

}

func CountValidPassphraseAnagrams(filename string) int {

    records := ReadTsvRecordsWithDelimiter(filename, ' ')
    goodPhraseCount := 0;
    for i := 0; i < len(records); i++ {
        phrase := records[i]
        strMap := make(map[string]int)
        goodPhrase := true;
        for j:=0; j < len(phrase); j++ {
            splitPhrase := strings.Split(phrase[j], "")
            sort.Strings(splitPhrase)
            alphaPhrase := strings.Join(splitPhrase, "")
            if strMap[alphaPhrase] == 1 {
                goodPhrase = false;
                break;
            }
            strMap[alphaPhrase] = 1
        }
        if goodPhrase {
            goodPhraseCount ++;
        }
    }

    return goodPhraseCount;

}
