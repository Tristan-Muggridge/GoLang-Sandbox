package person

type Gender int

const (
	Male Gender = iota
	Female
	NonBinary
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case NonBinary:
		return "Non-Binary"
	default:
		return "Unknown"
	}
}

func IntToGender(code int) Gender {
	switch code {
	case 1:
		return Male
	case 2:
		return Female
	default:
		return NonBinary
	}
}

type Person struct {
	Name   string
	Age    int
	Gender Gender
}

func CreatePerson(name string, age int, genderCode int) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		Gender: IntToGender(genderCode),
	}
}
