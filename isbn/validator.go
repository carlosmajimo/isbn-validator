package isbn

import (
	"strings"
	"unicode"
)

type ValidateISBN struct {
	isbn string
}

func NewValidateISBN(isbn string) *ValidateISBN {
	return &ValidateISBN{
		isbn: isbn,
	}
}

func (v *ValidateISBN) IsValid() bool {
	cleanISBN := strings.ReplaceAll(v.isbn, "-", "")

	length := len(cleanISBN)
	if length != 10 && length != 13 {
		return false
	}

	if !v.hasValidCharacters(cleanISBN) {
		return false
	}

	if length == 10 {
		return v.isValidISBN10(cleanISBN)
	}

	return v.isValidISBN13(cleanISBN)
}

func (v *ValidateISBN) hasValidCharacters(isbn string) bool {
	for i, char := range isbn {
		if i == len(isbn)-1 && len(isbn) == 10 && char == 'X' {
			continue
		}

		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func (v *ValidateISBN) isValidISBN10(isbn string) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(isbn[i] - '0')
		sum += digit * (10 - i)
	}

	lastChar := isbn[9]
	if lastChar == 'X' {
		sum += 10
	} else {
		sum += int(lastChar - '0')
	}

	return sum%11 == 0
}

func (v *ValidateISBN) isValidISBN13(isbn string) bool {
	sum := 0
	for i := 0; i < 13; i++ {
		digit := int(isbn[i] - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	return sum%10 == 0
}
