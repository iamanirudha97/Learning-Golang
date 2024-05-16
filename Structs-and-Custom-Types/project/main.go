package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	note1, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	note1.Display()

	err = note1.SaveNote()
	if err != nil {
		fmt.Println("FAILED SAVING THE NOTE", err)
		return
	}
	fmt.Println("NOTE SAVED")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSpace(text)
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Title : ")
	desc := getUserInput("Enter Description : ")
	return title, desc
}
