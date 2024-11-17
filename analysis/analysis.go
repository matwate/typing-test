package analysis

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hbollon/go-edlib"
)

// What we'll gonna do is analyze an input string of english words
// And we'll return how accurate the typing was.

var Allowed_words = []string{}

func Distance(input string, target string) int {
	return edlib.LevenshteinDistance(input, target)
}

func load_allowed_words() {
	// This will load the allowed words from a file

	file, err := os.Open("analysis/allowed_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Allowed_words = append(Allowed_words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func find_proper_target(word string) string {
	// This will fzf the target word that most likely the user was trying to type
	if len(Allowed_words) == 0 {
		panic("Allowed words not loaded")
	}

	res, err := edlib.FuzzySearchSet(word, Allowed_words, 1, edlib.Levenshtein)
	if err != nil {
		return word
	}
	if len(res) == 0 {
		return word
	}
	fmt.Println(res)
	return res[0]
}

func construct_target_sentence(sentence string) string {
	load_allowed_words()
	// This will construct the target sentence from the input sentence
	// We'll remove all the punctuations and convert the sentence to lowercase
	words := strings.ReplaceAll(sentence, ".", "")
	words = strings.ReplaceAll(words, ",", "")
	words = strings.ReplaceAll(words, "!", "")
	words = strings.ReplaceAll(words, "?", "")
	words = strings.ReplaceAll(words, ":", "")
	words = strings.ReplaceAll(words, ";", "")
	words = strings.ReplaceAll(words, "(", "")
	words = strings.ReplaceAll(words, ")", "")
	words = strings.TrimSpace(words)
	words = strings.ToLower(words)
	tokens := strings.Split(words, " ")
	target_sentence := ""
	for _, word := range tokens {
		target_word := find_proper_target(word)
		target_sentence += target_word + " "
	}
	return target_sentence
}

func Accuracy(input string) float64 {
	// This will calculate the accuracy of the input sentence based on what
	// the algorithm thinks the user was trying to type
	target_sentence := construct_target_sentence(input)
	distance := Distance(input, target_sentence)
	return 1 - (float64(distance) / float64(len(target_sentence)))
}
