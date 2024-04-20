package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		displayHelp()
		return
	}
	collectionName := os.Args[1]
	notes, err := loadCollection(collectionName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		displayMenu()
		choice, err := getUserChoice(reader)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		handleChoice(choice, collectionName, &notes)
		if choice == 4 {
			break
		}
	}
}

func displayHelp() {
	fmt.Println("Usage: ./notestool collection_name")
	fmt.Println("  collection_name: The name of the note collection to manage.")
	fmt.Println("  help: Display this help message.")
}

func loadCollection(collectionName string) ([]string, error) {
	var notes []string
	file, err := os.Open(collectionName)
	if err != nil {
		file, err = os.Create(collectionName)
		if err != nil {
			return nil, err
		}
		return notes, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes, scanner.Err()
}

func getUserChoice(reader *bufio.Reader) (int, error) {
	fmt.Println("Enter your choice:")
	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
	return choice, err
}

func handleChoice(choice int, collectionName string, notes *[]string) {
	switch choice {
	case 1:
		displayNotes(*notes)
	case 2:
		addNewNote(collectionName, notes)
	case 3:
		removeNote(collectionName, notes)
	case 4:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func displayMenu() {
	fmt.Println("1. Display notes")
	fmt.Println("2. Add a note")
	fmt.Println("3. Remove a note")
	fmt.Println("4. Exit")
}

func displayNotes(notes []string) {
	if len(notes) == 0 {
		fmt.Println("No notes in the collection.")
		return
	}
	for index, note := range notes {
		fmt.Printf("%03d - %s\n", index+1, note)
	}
}

func addNewNote(collectionName string, notes *[]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the note text:")
	noteText, _ := reader.ReadString('\n')
	*notes = append(*notes, strings.TrimSpace(noteText))
	saveNotes(collectionName, *notes)
}

func removeNote(collectionName string, notes *[]string) {
	fmt.Println("Enter the number of the note to remove:")
	reader := bufio.NewReader(os.Stdin)
	numberStr, _ := reader.ReadString('\n')
	number, _ := strconv.Atoi(strings.TrimSpace(numberStr))
	number -= 1

	if number < 0 || number >= len(*notes) {
		fmt.Println("Invalid note number.")
		return
	}

	*notes = append((*notes)[:number], (*notes)[number+1:]...)
	saveNotes(collectionName, *notes)
}

func saveNotes(collectionName string, notes []string) error {
	file, err := os.Create(collectionName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note := range notes {
		_, err := writer.WriteString(note + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
