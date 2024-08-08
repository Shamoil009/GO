package main

import "fmt"

func main() {

	noteTitle:=getNoteData("Note Title: ")
	noteContent:=getNoteData("Note Content: ")


}

func getNoteData(promtText string) string {
	fmt.Println(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}
