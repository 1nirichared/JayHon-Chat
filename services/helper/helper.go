package helper

import (
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
)

func Bcrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}

		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}

	default:
		return false
	}
	return false
}
