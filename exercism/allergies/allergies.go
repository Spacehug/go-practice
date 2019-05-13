package allergies

var substances = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(input uint) []string {
	ret := make([]string, 0, 8)
	bytes := byte(input)
	for i := 0; bytes > 0; i++ {
		if bytes&1 == 1 {
			ret = append(ret, substances[i])
		}
		bytes >>= 1
	}
	return ret
}

func AllergicTo(score uint, substance string) bool {
	for index, allergen := range substances {
		if substance == allergen {
			return score&(1<<uint(index)) != 0
		}
	}
	return false
}
