package pkg

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const pi = 3.14

func ForthTaskFirst() {
	var numberofelements int
	log.Println("Enter the number of elements in the array:")
	fmt.Scan(&numberofelements)
	var arr []string
	for i := 0; i < numberofelements; i++ {
		fmt.Scan(&arr[i])
	}
	log.Print(arr[:numberofelements])
}
func ForthTaskSecond() {
	var size int
	fmt.Println("Please, enter the size of natural array: ")
	fmt.Scan(&size)
	biggest, err := compare(size)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println(biggest)
}
func ForthTaskThird() {
	var size int
	fmt.Println("Please, enter the size of natural array: ")
	fmt.Scan(&size)
	sum, err := arraysum(size)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)
}
func ForthTaskForth() {
	var str [5]string
	for i := 0; i < 5; i++ {
		fmt.Scan(&str[i])
	}
	result := concateniate(str)
	fmt.Println(result)
}
func ForthTaskFifth() {
	arr := [6]float64{}
	for i := 0; i < 6; i++ {
		fmt.Scan(&arr[i])
	}
	idx := [2]int{}
	for i := 0; i < 2; i++ {
		fmt.Scan(&idx[i])
	}
	swapelements(&arr, idx)
	fmt.Println(arr)
}
func ForthTaskSixth() {
	var r float64
	fmt.Scan(&r)
	length, area := calculate(r)
	fmt.Println(length, area)
}
func ForthTaskSeventh() {
	arr := [5]float64{263.332, 12.1456, 567.12, 55.90, 97.987654}
	rounding(&arr)
	fmt.Println(arr[1], arr[4])
}
func ForthTaskEigths() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	showingelementswithevenendexes(arr)
}
func ForthTaskNineth() {
	var arr [15]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	counter := countEvenElements(arr)
	fmt.Println(counter)
}
func ForthTaskTenth() {
	var number int64
	fmt.Scan(&number)
	result := convert(number)
	fmt.Println(result)
}
func ForthTaskEleventh() {
	var strOne, strTwo string
	fmt.Scan(&strOne, &strTwo)
	result := summing(strOne, strTwo)
	fmt.Println(result)
}
func summing(strOne, strTwo string) (result int64) {
	intOne, err := strconv.ParseInt(strOne, 10, 64)
	if err != nil {
		fmt.Println("You did not entered the numbers")
		return
	}
	intTwo, err := strconv.ParseInt(strTwo, 10, 64)
	if err != nil {
		fmt.Println("You did not entered the numbers")
		return
	}
	fmt.Println(err)
	result = intOne + intTwo
	return
}
func convert(number int64) (result uint16) {
	result = uint16(number)
	return
}
func countEvenElements(arr [15]int) (counter int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			counter++
		}
	}
	return
}
func concateniate(arr [5]string) (result string) {
	for _, str := range arr {
		result += str
	}
	return
}
func arraysum(size int) (sum int, err string) {
	var arr [100]int
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
		if arr[i] < 1 {
			err = "its not the natural number!"
			return
		}
		sum += arr[i]
	}
	return
}
func compare(size int) (biggest int, err string) {
	var arr [100]int
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
		if biggest < arr[i] {
			biggest = arr[i]
		}
		if arr[i] < 1 {
			err = "It is not a natural number"
			return
		}
	}
	return
}
func swapelements(arr *[6]float64, idx [2]int) {
	temp := arr[idx[0]]
	arr[idx[0]] = arr[idx[1]]
	arr[idx[1]] = temp
}
func calculate(r float64) (length float64, area float64) {
	length = 2 * pi * r
	area = pi * r * r
	return
}
func rounding(arr *[5]float64) {
	for i := 0; i < 5; i++ {
		arr[i] = math.Floor(arr[i]*100) / 100
	}
}
func showingelementswithevenendexes(arr [10]int) {
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Println(arr[i])
		}
	}
}
