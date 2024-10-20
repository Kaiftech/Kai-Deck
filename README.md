# KaifCards

**KaifCards** is a Command-Line Interface (CLI) application built in Go, designed to help users learn through flashcards. It uses a simple spaced repetition system to make learning efficient by adjusting review frequency based on user progress. The flashcards are stored in a JSON file, allowing users to retain and manage their flashcards between sessions.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Add Flashcards](#add-flashcards)
  - [Review Flashcards](#review-flashcards)
  - [View Progress](#view-progress)
  - [Exit](#exit)
- [Data Storage](#data-storage)
- [Future Enhancements](#future-enhancements)

## Features

- **Add Flashcards**: Users can create new flashcards by inputting a question and its corresponding answer.
- **Review Flashcards**: Users can review flashcards in the application and indicate if they've learned the answer.
- **Track Progress**: The application tracks flashcard review history, adjusting the review intervals to focus on unlearned material.
- **JSON Storage**: Flashcards are stored in a JSON file, allowing easy access and persistence across sessions.
- **Interactive Feedback**: ASCII art and fun messages keep users motivated and engaged.

## Installation

To install and run **KaifCards**, you will need to have Go installed on your system. If you haven't installed Go yet, [follow this guide](https://golang.org/doc/install).

### Steps:

1. Clone this repository:
   ```bash
   git clone https://github.com/YourUsername/KaifCards.git
   ```

2. Navigate into the project directory:
   ```bash
   cd KaifCards
   ```

3. Build the application:
   ```bash
   go build main.go
   ```

4. Run the application:
   ```bash
   ./main
   ```

That's it! Now you can start adding and reviewing your flashcards.

## Usage

Once you run the application, you will be presented with a simple menu:

```

██╗░░██╗░█████╗░██╗  ██████╗░███████╗░█████╗░██╗░░██╗░██████╗
██║░██╔╝██╔══██╗██║  ██╔══██╗██╔════╝██╔══██╗██║░██╔╝██╔════╝
█████═╝░███████║██║  ██║░░██║█████╗░░██║░░╚═╝█████═╝░╚█████╗░
██╔═██╗░██╔══██║██║  ██║░░██║██╔══╝░░██║░░██╗██╔═██╗░░╚═══██╗
██║░╚██╗██║░░██║██║  ██████╔╝███████╗╚█████╔╝██║░╚██╗██████╔╝
╚═╝░░╚═╝╚═╝░░╚═╝╚═╝  ╚═════╝░╚══════╝░╚════╝░╚═╝░░╚═╝╚═════╝░

1. Add Flashcard
2. Review Flashcards
3. Exit
```

### Add Flashcards

To add a flashcard:

1. Select the **"Add Flashcard"** option from the menu.
2. The application will prompt you to input a question.
3. Next, you will be prompted to provide the answer for the question.
4. The flashcard will be saved and you will be returned to the main menu.

### Review Flashcards

To review your flashcards:

1. Select the **"Review Flashcards"** option.
2. The application will present flashcards that have not been marked as "Learned."
3. For each flashcard, you will be asked if you've learned the card:
   - If you answer **yes**, the flashcard will be marked as learned, and the app will cheer you on with ASCII art!
   - If you answer **no**, the app will slightly "annoy" you, prompting you to try harder.

### View Progress

The review system is built on spaced repetition, meaning that your review sessions will focus on unlearned cards or cards that haven't been reviewed in a while. Progress tracking is based on whether you've successfully learned a flashcard or not.

### Exit

Select **"Exit"** to quit the application. The app will thank you with some ASCII art and save your progress automatically.

## Data Storage

All flashcards are saved in a JSON file (`flashcards.json`) in the same directory as the application. This file ensures that all your data is persistent and can be accessed when you run the application again.

### Flashcard Structure

Each flashcard is stored with the following properties:

```json
{
  "question": "What is the capital of France?",
  "answer": "Paris",
  "learned": false,
  "last_reviewed": "2024-10-19T15:00:00Z",
  "review_interval": 1
}
```

- **Question**: The question for the flashcard.
- **Answer**: The answer for the flashcard.
- **Learned**: A boolean indicating whether the user has learned the flashcard.
- **LastReviewed**: A timestamp of when the flashcard was last reviewed.
- **ReviewInterval**: The interval in days for the next review (spaced repetition).

## Future Enhancements

- **Spaced Repetition Algorithm**: Implement a more advanced algorithm to intelligently space out the reviews based on user performance.
- **Flashcard Editing**: Allow users to edit or delete existing flashcards.
- **Categories/Tags**: Add support for categories or tags to organize flashcards.
- **Statistics**: Display learning statistics, such as the number of flashcards learned over time.


Let me know if you need any changes or adjustments!
