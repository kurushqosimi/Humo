package pkg

import (
	"fmt"
	"log"
	"strconv"
)

func FifthTaskFirst() {
	var n int
	fmt.Scan(&n)
	if n <= 4 {
		log.Println("You need more than 4 elements.")
		return
	}
	First(n)
}
func FifthTaskSecond() {
	var num, n int
	fmt.Scan(&num, &n)
	Second(num, n)
}
func FifthTaskThird() {
	var a, b int
	fmt.Scan(&a, &b)
	max := b - b%7
	if max < a {
		fmt.Println("-1")
	} else {
		fmt.Println(max)
	}
}
func FifthTaskForth() {
	forth()
}
func FifthTaskFifth() {
	fifth()
}

func First(n int) {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	log.Print(slice[3:4])
}
func Second(num, n int) {
	var slice []int
	for {
		if num < 10 {
			slice = append(slice, num)
			break
		} else {
			slice = append(slice, num%10)
			num /= 10
		}
	}
	var str string
	for i := 0; i < len(slice); i++ {
		if slice[i] == n {
			continue
		} else {
			str += strconv.Itoa(slice[i])
		}
	}
	str = reverse(str)
	log.Println(str)
}
func reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func forth() {
	var num [4]int
	for i := 0; i < 4; i++ {
		fmt.Scan(&num[i])
	}
	for i := 0; i < len(num)-1; i++ {
		for j := 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	fmt.Println(num[0])
}
func fifth() {
	slice := make([]int, 10)
	var a int
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		slice[i] = a
	}
	fmt.Println(slice)
}
