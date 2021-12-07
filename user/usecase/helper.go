package usecase

import "strconv"

const (
	IINLength = 12
)

func InvalidField(field string) bool {
	return field == ""
}

func InvalidIIN(IIN string) bool {
	if isNotNumber(IIN) {
		return true
	}
	if len(IIN) != IINLength {
		return true
	}
	return false
}

func isNotNumber(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return true
	}
	return false
}
