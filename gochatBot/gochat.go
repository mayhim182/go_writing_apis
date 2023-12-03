package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Initize the NLP
func main() {
	fmt.Printf("Random print")
	model, err := sentiment.Restore()
	if err != nil {
		log.Fatalf("Error restoring sentimental model: %v", err)
	}

	//Set up terminal I/O
	t, err := tty.Open()
	if err != nil {
		log.Fatalf("Error opening tty: %v", err)
	}
	defer t.Close()

	//Set up reader for user input
	reader := bufio.NewReader(os.Stdin)

	for {

	}
}
