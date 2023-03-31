package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	var name string
	var age int

	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	// Get user input for name
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err != nil || utf8.RuneCountInString(name) == 0 {
		log.Fatal("Empty string input received")
	}

	fmt.Printf("Welcome %s! \n", name)

	fmt.Print("Please enter your age: ")
	fmt.Scanln(&age)

	canAccess := func(age int) bool {
		return age >= 18
	}

	if canAccess(age) {
		fmt.Println("Welcome in bud.")
	} else {
		fmt.Println("Sorry, can't let you in.")
	}
}
