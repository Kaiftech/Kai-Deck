package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

// ASCII art functions for different moods
func showHappyArt() {
	fmt.Println(`
		\(^o^)/
		You did it! Keep going! You're awesome! 🎉
	`)
}

func showAnnoyingArt() {
	fmt.Println(`
		(ಠ_ಠ)
		Come on, you should have learned this by now! 😠
	`)
}

func showExitArt() {
	fmt.Println(`
		┌( ಠ‿ಠ)┘
		Thank you for using Progressive Flashcards! Bye!
	`)
}

func showAddFlashcardArt() {
	fmt.Println(`
		(ノ◕ヮ◕)ノ*:・゚✧
		Adding a new flashcard to boost your learning!
	`)
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

// Add a new flashcard
func addFlashcard(flashcards []Flashcard) []Flashcard {
	showAddFlashcardArt() // Show happy art when adding

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter question: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	fmt.Print("Enter answer: ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	card := Flashcard{
		Question:      question,
		Answer:        answer,
		Learned:       false,
		LastReviewed:  time.Now(),
		ReviewInterval: 1,
	}

	return append(flashcards, card)
}

// Review flashcards
func reviewFlashcards(flashcards []Flashcard) []Flashcard {
	reader := bufio.NewReader(os.Stdin)
	for i, card := range flashcards {
		if card.Learned {
			continue
		}

		fmt.Println("Question:", card.Question)
		fmt.Print("Did you learn this card? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "yes" {
			flashcards[i].Learned = true
			showHappyArt() // Encourage user for success
		} else {
			showAnnoyingArt() // Annoy the user for not learning
		}

		flashcards[i].LastReviewed = time.Now()
	}

	return flashcards
}

// Show main menu with ASCII art
func showMenu() {
	fmt.Println(`

	██╗░░██╗░█████╗░██╗  ██████╗░███████╗░█████╗░██╗░░██╗░██████╗
	██║░██╔╝██╔══██╗██║  ██╔══██╗██╔════╝██╔══██╗██║░██╔╝██╔════╝
	█████═╝░███████║██║  ██║░░██║█████╗░░██║░░╚═╝█████═╝░╚█████╗░
	██╔═██╗░██╔══██║██║  ██║░░██║██╔══╝░░██║░░██╗██╔═██╗░░╚═══██╗
	██║░╚██╗██║░░██║██║  ██████╔╝███████╗╚█████╔╝██║░╚██╗██████╔╝
	╚═╝░░╚═╝╚═╝░░╚═╝╚═╝  ╚═════╝░╚══════╝░╚════╝░╚═╝░░╚═╝╚═════╝░
	`)
	fmt.Println("1. Add Flashcard")
	fmt.Println("2. Review Flashcards")
	fmt.Println("3. Exit")
}

func main() {
	const filename = "flashcards.json"

	flashcards, err := loadFlashcards(filename)
	if err != nil {
		fmt.Println("No flashcards found, starting fresh.")
	}

	for {
		showMenu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			flashcards = addFlashcard(flashcards)
			saveFlashcards(filename, flashcards)
		case 2:
			flashcards = reviewFlashcards(flashcards)
			saveFlashcards(filename, flashcards)
		case 3:
			showExitArt() // Show thank you art for exiting
			return
		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}
