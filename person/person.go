package person

import "sandbox.com/person/gender"

type Person struct {
	Name   string
	Age    int
	Gender gender.Gender
}

func CreatePerson(name string, age int, genderCode int) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		Gender: gender.IntToGender(genderCode),
	}
}
