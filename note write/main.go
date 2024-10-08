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
	Save() error // error is return value
}

type outputtable interface {
	saver
	Display()
}

func main() {
	
	printSomething(1)
	printSomething(1.5)
	printSomething("string print something")


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

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

// you can use interface{} or any to allow anything
func printSomething(value interface{}) {

	// way 1
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	// way 2
	// type switching
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
