package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func New(content string) (*Todo, error) {

	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo) // marshal require public fields

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // writeFile return err also
}
