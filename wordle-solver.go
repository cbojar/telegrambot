package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordleSolver struct {
	dictionary string
}

func CreateWordleSovler(dictionary string) *WordleSolver {
	return &WordleSolver{
		dictionary: dictionary,
	}
}

func (solver *WordleSolver) Solve(correct string, misplaced string, incorrect string) ([]string, error) {
	dictionary, err := os.Open(solver.dictionary)
	if err != nil {
		return nil, fmt.Errorf("failed to open dictionary at \"%s\": %w", solver.dictionary, err)
	}

	words := bufio.NewScanner(dictionary)
	matches := make([]string, 0)

	for words.Scan() {
		word := words.Text()
		if matchesLength(word, correct) &&
			!hasInvalidCharacters(word) &&
			matchesCorrectLetters(word, correct) &&
			matchesMisplacedLetters(word, misplaced) &&
			!matchesIncorrectLetters(word, incorrect) {
			matches = append(matches, word)
		}
	}

	if err = words.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func matchesLength(word string, correct string) bool {
	return len(word) == len(correct)
}

func hasInvalidCharacters(word string) bool {
	for _, char := range word {
		if char < 'a' || char > 'z' {
			return true
		}
	}

	return false
}

func matchesCorrectLetters(word string, correct string) bool {
	wordRunes := []rune(word)

	for position, char := range correct {
		if char == '_' {
			continue
		}

		if char != wordRunes[position] {
			return false
		}
	}

	return true
}

func matchesMisplacedLetters(word string, misplaced string) bool {
	for _, char := range misplaced {
		if !strings.ContainsRune(word, char) {
			return false
		}
	}

	return true
}

func matchesIncorrectLetters(word string, invalid string) bool {
	for _, char := range invalid {
		if strings.ContainsRune(word, char) {
			return true
		}
	}

	return false
}
