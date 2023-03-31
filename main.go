package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	// import the person module
	"sandbox.com/person"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else if utf8.RuneCountInString(name) == 0 {
		log.Fatal("Name was empty.")
	}

	return strings.TrimSpace(name)
}

func main() {

	fmt.Println("Hello there! Welcome to the world of Pokémon!")
	fmt.Println("My name is Oak! People call me the Pokémon Prof!")
	fmt.Println("This world is inhabited by creatures called Pokémon!")
	fmt.Println("For some people, Pokémon are pets. Others use them for fights. Myself...I study Pokémon as a profession.")

	fmt.Println("First, what is your name?")
	name := ReadInput()

	fmt.Println("Next, how old are you?")
	age, err := strconv.Atoi(ReadInput())
	if err != nil {
		log.Fatal(err)
	}

	canBecomeTrainer := func(age int) bool {
		return age >= 10
	}

	if !canBecomeTrainer(age) {
		fmt.Printf("Sorry, you're too young to become a Pokémon trainer, come back in %d years.", 10-age)
		os.Exit(0)
	}

	fmt.Println("Lastly, are you a boy (1), girl (2), or non-binary (3)?")
	genderCode, err := strconv.Atoi(ReadInput())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("So, to confirm...")
	profile := person.CreatePerson(name, age, genderCode)

	fmt.Printf("Your name is %s and you're a %d old %s", profile.Name, profile.Age, profile.Gender)
}
