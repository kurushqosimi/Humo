package pkg

import (
	"errors"
	"fmt"
	"math"
)

func SeventhTaskOne() {
	var n int
	fmt.Scan(&n)
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	myMap := makeMap(slice)
	fmt.Println(myMap)
}
func makeMap(slice []float64) map[int]float64 {
	MyMap := make(map[int]float64, len(slice))
	for i := 0; i < len(slice); i++ {
		MyMap[i] = slice[i]
	}
	return MyMap
}
func SeventhTaskTwo() {
	myMap := make(map[string]int)
	var n, number int
	var name string
	fmt.Println("Enter the number of students: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the Name: ")
		fmt.Scan(&name)
		fmt.Println("Enter the phone number: ")
		fmt.Scan(&number)
		myMap[name] = number
	}
	fmt.Println(myMap)
}
func SeventhTaskThree() {
	CyrillicToLatin := map[string]string{"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e", "ё": "yo", "ж": "zh", "з": "z",
		"и": "i", "й": "y", "к": "k", "л": "l", "м": "m", "н": "n", "о": "o", "п": "p", "р": "r", "с": "s", "т": "t", "у": "u",
		"ф": "f", "х": "kh", "ц": "ts", "ч": "ch", "ш": "sh", "щ": "sch", "ъ": "'", "ы": "y", "ь": "", "э": "e", "ю": "yu", "я": "ya"}
	fmt.Println(CyrillicToLatin)
	var cyrillic string
	fmt.Scan(&cyrillic)
	value := CyrillicToLatin[cyrillic]
	fmt.Println(value)
}
func SeventhTaskFour() {
	var person Person
	Die(&person)
}

type Person struct {
	Name string
}

func (p *Person) Eat() {
	fmt.Printf("%v, is eating.", p.Name)
}
func (p *Person) Sleep() {
	fmt.Printf("%v, is sleeping.", p.Name)
}
func (p *Person) Walk() {
	fmt.Printf("%v, is walking.", p.Name)
}

type Human interface {
	Eat()
	Sleep()
	Walk()
}

func Die(h Human) {
	fmt.Println("died!...")
}
func SeventhTaskFive() {
	myCalc := SiemensCalculator{Company: "Sony", Year: 2012, Isworking: true}
	fmt.Println(myCalc)
}

type SiemensCalculator struct {
	Company   string
	Year      int
	Isworking bool
}

func SeventhTaskSix() {
	var calc SiemensCalculator
	calc.Company = "Sony"
	calc.Year = 2012
	calc.Isworking = true
	var fnum, snum int
	fmt.Scan(&fnum, &snum)
	if calc.Isworking == true {
		res, err := calc.Divide(fnum, snum)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	} else {
		fmt.Println("Your calculator is not working!")
	}
}
func SeventhTaskSeven() {
	calc := SiemensCalculator{
		Company:   "Siemens",
		Year:      2022,
		Isworking: true,
	}
	calc.Multiply(2, 3)
	calc.Sum(1, 2)
	calc.Divide(2, 3)
}

type Calculator interface {
	Divide(a int, b int) (int, error)
	Multiply(a int, b int) int
	Sum(a int, b int) int
}

func (s SiemensCalculator) Divide(a int, b int) (int, error) {
	var result int
	switch {
	case a > b:
		errors.New("Your first number must be greater, other way your result always would be zero!")
	case b == 0:
		errors.New("Your second number should not be zero!")
	default:
		result = a / b
	}
	return result, nil
}
func (s SiemensCalculator) Multiply(a int, b int) int {
	result := a * b
	return result
}
func (s SiemensCalculator) Sum(a int, b int) int {
	result := a + b
	return result
}
func SeventhTaskEight() {
	var b bool
	_, err := fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	fmt.Println("Вы ввели:", b)
}

type rectangle struct {
	width  float64
	height float64
}

func SeventhTaskNine() {
	var rect rectangle
	fmt.Println("enter the height of rectangle: ")
	fmt.Scan(&rect.height)
	fmt.Println("enter the height of width: ")
	fmt.Scan(&rect.width)
	fmt.Println(rect.Square())
}
func (r rectangle) Square() float64 {
	if r.width == 0 || r.height == 0 {
		panic("Oops, you entered zero as an argument, if its correct you answer is zero, other way restart the program!")
	} else {
		res := r.width * r.height
		return res
	}
}

func SeventhTaskTen() {
	var fcatet, scatet float64
	fmt.Scan(&fcatet, &scatet)
	res := findHypotenuse(fcatet, scatet)
	fmt.Printf("%.2f", res)
}
func findHypotenuse(a, b float64) float64 {
	if a == 0 || b == 0 {
		panic("a or b cannot be zero")
	}
	return math.Sqrt(a*a + b*b)
}
