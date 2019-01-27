//Based on How to Write a Spelling Corrector - http://norvig.com/spell-correct.html
package spellchecker

import (
	"bufio"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	corpusLocation = "http://norvig.com/big.txt"
	letters        = "abcdefghijklmnopqrstuvwxyz"
)

func CheckWord(word string) (string, bool) {

	dict := getDict()
	if _, found := dict[word]; found {
		return word, true
	}

	return mostProbable(validWords(candidateWords(word), dict), dict), false

}

func candidateWords(word string) []string {

	editDist1Words := editDist1(word)
	editDist2Words := make([]string, len(editDist1Words))
	for _, modifiedWord := range editDist1Words {
		editDist2Words = append(editDist2Words, editDist1(modifiedWord)...)
	}

	return editDist1Words

}

func validWords(wordSet []string, dict map[string]int) []string {

	var words []string
	for _, word := range wordSet {
		if _, present := dict[word]; present {
			words = append(words, word)

		}
	}

	return words

}

func getDict() map[string]int {

	//Load dictionary if already saved
	dict := make(map[string]int)
	err := readGob("./dict.gob", &dict)
	if err == nil {
		return dict
	}

	dict = getWordCounts()
	err = writeGob("./dict.gob", dict)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return dict
}

func editDist1(word string) []string {

	deletes, transposes, replaces, inserts := make([]string, len(word)), make([]string, len(word)), make([]string, len(word)), make([]string, len(word))
	for i := 0; i <= len(word); i++ {

		if i < len(word) {
			deletes = append(deletes, word[:i]+word[i+1:])
		}

		if len(word)-i > 2 {
			transposes = append(transposes, word[:i]+word[i+1:i+2]+word[i:i+1]+word[i+2:])
		}

		for _, c := range letters {
			if i < len(word) {
				replaces = append(replaces, word[:i]+string(c)+word[i+1:])
			}
			inserts = append(inserts, word[:i]+string(c)+word[i:])
		}
	}

	return set(append(append(append(set(deletes), set(transposes)...), set(replaces)...), set(inserts)...))

}

func getWordCounts() map[string]int {

	resp, err := http.Get(corpusLocation)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	wordCounts := make(map[string]int)
	s := bufio.NewScanner(resp.Body)

	for tok := s.Scan(); tok != false; tok = s.Scan() {
		toks := strings.Fields(s.Text())
		for _, t := range toks {
			wordCounts[strings.ToLower(t)]++
		}
	}

	return wordCounts

}
func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(object)
	}
	err = file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	err = file.Close()
	return err
}

func set(words []string) []string {

	wordSet := make(map[string]bool)
	for _, word := range words {
		if _, ok := wordSet[word]; !ok {
			wordSet[word] = true
		}
	}

	retval := make([]string, len(wordSet))
	for key := range wordSet {
		if len(key) > 0 {
			retval = append(retval, key)
		}
	}

	return retval

}

func mostProbable(words []string, dict map[string]int) string {

	var mostProbableWord string
	maxCount := -1

	for _, word := range words {
		if count, found := dict[word]; found {
			if count > maxCount {
				maxCount = count
				mostProbableWord = word
			}
		}
	}

	return mostProbableWord

}
