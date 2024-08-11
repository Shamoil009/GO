package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	noteTitle, noteContent := getNoteData()

	var userNote *note.Note
	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving note failed")
		fmt.Println(err)
		return
	}
	fmt.Println("Saving note succeeded!")

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
