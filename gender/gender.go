package gender

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
