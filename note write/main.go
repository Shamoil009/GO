package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	noteTitle, noteContent := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	var userNote *note.Note
	userNote, err = note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed")
		fmt.Println(err)
		return err
	}
	fmt.Println("Saving note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note Title: ")
	noteContent := getUserInput("Note Content: ")

	return noteTitle, noteContent
}

func getUserInput(promtText string) string {
	fmt.Print(promtText)

	// Scanln() can't handle long text input
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
