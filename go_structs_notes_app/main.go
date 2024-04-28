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

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething("Welcome")
	printSomething(1)
	printSomething(2.5)

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	// We can pass our todo struct value to this function but it won't do anything since we
	// don't hande that Todo struct type inside this printSomething()
	printSomething(todo)

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

// interface{} can also be replaced with "any" since it's an alias for it.
func printSomething(value interface{}) {
	// ok will indicate if value is of that type specified inside parenthesis.
	// intVal will be that value but now Go will know that it's of that type mentioned in ()
	// Alternative to Type Switch for executing different code block based on a specific type
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// Type Switch
	// value.(type) works only with switch statement as it's a special use case in Go
	// switch value.(type) {
	// case int:
	// 	fmt.Printf("Integer: %v\n", value)
	// case float64:
	// 	fmt.Printf("Float: %v\n", value)
	// case string:
	// 	fmt.Printf("String: %v\n", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
