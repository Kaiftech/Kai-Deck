package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Flashcard struct
type Flashcard struct {
	Question      string    `json:"question"`
	Answer        string    `json:"answer"`
	Learned       bool      `json:"learned"`
	LastReviewed  time.Time `json:"last_reviewed"`
	ReviewInterval int      `json:"review_interval"`
}

// Load flashcards from a JSON file
func loadFlashcards(filename string) ([]Flashcard, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Flashcard{}, err
	}

	var flashcards []Flashcard
	err = json.Unmarshal(file, &flashcards)
	if err != nil {
		return []Flashcard{}, err
	}
	return flashcards, nil
}

// Save flashcards to a JSON file
func saveFlashcards(filename string, flashcards []Flashcard) error {
	data, err := json.MarshalIndent(flashcards, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
