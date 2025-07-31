package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("My first code on GitHub")
}

type Page struct {
	Title string
	Body  []byte
}

func (P *Page) save() error {
	filename := P.Title + ".txt"
	return os.WriteFile(filename, P.Body, 0644)
}
