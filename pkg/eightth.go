package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"unicode"
)

type Student struct {
	Name    string `json:"first_name"`
	Surname string `json:"Last_name"`
	Age     int    `json:"age"`
	Number  int    `json:"phone_number"`
	Married string `json:"married"`
}

func EighthTaskOne() {
	var str string
	fmt.Scan(&str)
	res := countSymbols(str)
	fmt.Println(res)
}
func EighthTaskTwo() {
	var str string
	fmt.Scan(&str)
	res, err := IsPalindrome(str)
	if err != nil {
		fmt.Println(res, err)
		return
	}
	fmt.Println(res)
}
func EighthTaskThree() {
	var str string
	fmt.Scan(&str)
	res := Evenstrings(str)
	fmt.Println(res)
}
func EighthTaskFour() {
	var password string
	fmt.Scan(&password)
	res, err := IsValid(password)
	if err != nil {
		fmt.Println(res, err)
		return
	}
	fmt.Println(res)
}
func EighthTaskFive() {
	student := Student{
		Name:    "Bahodoor",
		Surname: "Negmatov",
		Age:     20,
		Number:  1234567890,
		Married: "yes",
	}
	bytes, err := json.Marshal(student)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
func EighthTaskSix() {
	requestStr := `{"first_name": "Bahodoor", "Last_name":"Negmatov", "age":23, "phone_number": 29328181, "married": "yes" }`
	request := []byte(requestStr)
	var someStudent Student
	err := json.Unmarshal(request, &someStudent)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(someStudent)
}
func countSymbols(str string) (res int) {
	runes := []rune(str)
	res = len(runes)
	return res
}
func IsPalindrome(str string) (res bool, err error) {
	runes := []rune(str)
	var count int
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] == runes[len(runes)-i-1] {
			count++
		}
	}
	if count == len(runes)/2 {
		res = true
		return
	} else {
		res = false
		err = errors.New("нолан недоволен")
		return
	}
}
func Evenstrings(str string) string {
	var strTwo string
	for i := 0; i < len(str); i += 2 {
		strTwo += string(str[i])
	}
	return strTwo
}
func IsValid(password string) (res bool, err error) {
	runes := []rune(password)
	regex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	b := regex.MatchString(password)
	if len(runes) < 8 {
		err = errors.New("your password is too short")
		res = false
		return
	}
	var counter int
	for _, r := range password {
		if unicode.Is(unicode.Latin, r) {
			res = true
		}
		if unicode.Is(unicode.Digit, r) {
			res = true
			counter++
		}
		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Digit, r) && !b {
			res = false
			err = errors.New("your word is not latin")
			return
		}
	}
	for _, r := range password {
		if unicode.Is(unicode.Digit, r) {
			res = true
			counter++
		}
	}
	if counter == 0 {
		res = false
		err = errors.New("you do not have digits in your code")
		return
	}
	if !b {
		res = false
		err = errors.New("your password does not contain special characters")
		return
	}
	hasLower := false
	hasUpper := false
	for _, r := range password {
		if unicode.IsLower(r) {
			hasLower = true
		} else if unicode.IsUpper(r) {
			hasUpper = true
		}
	}
	if hasLower && hasUpper {
		res = true
	} else {
		res = false
		err = errors.New("your password has to have uppercase letters or lowercase")
		return
	}
	return
}
