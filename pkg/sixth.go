package pkg

import (
	"errors"
	"fmt"
	"log"
)

type Card struct {
	Number  int
	Owner   string
	Balance float64
}
type Transaction struct {
	Sum        float64
	CardNumber int
}
type Hero struct {
	On    bool
	Ammo  int
	Power int
}

func SixthTaskFirst() {
	var res int
	var firstPointer, secondPointer *int
	*firstPointer = 3
	*secondPointer = 5
	multiplication(firstPointer, secondPointer, &res)
	log.Println(res)
}
func SixthTaskSecond() {
	first, second := 123, 321
	swap(&first, &second)
	log.Println(first, second)
}
func SixthTaskThird() {
	personCard := Card{Number: 888800001111, Owner: "Kurush", Balance: 100_000}
	var purchase Transaction
	fmt.Println("Enter the number of your card:")
	fmt.Scan(&purchase.CardNumber)
	fmt.Println("Enter the cost of purchase:")
	fmt.Scan(&purchase.Sum)
	err := purchasing(&personCard, &purchase)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Purchase has been successful. Your state is:", personCard.Balance)
}
func SixTaskForth() {
	Thor := Hero{true, 100, 100}
	shoot := Thor.Shoot()
	ride := Thor.RideBike()
	fmt.Println("Shoot:", shoot)
	fmt.Println("Ride:", ride)
}
func (h *Hero) Shoot() bool {
	if !h.On {
		return false
	}
	if h.Ammo > 0 {
		h.Ammo--
		return true
	}
	return false
}
func (h *Hero) RideBike() bool {
	if !h.On {
		return false
	}
	if h.Power > 0 {
		h.Power--
		return true
	}
	return false
}

//First task
func multiplication(first *int, second *int, res *int) {
	*res = (*first) * (*second)
}
func swap(first *int, second *int) {
	*first, *second = *second, *first
}
func purchasing(personCard *Card, purchase *Transaction) error {
	if personCard.Number != purchase.CardNumber {
		return errors.New("You entered the wrong Card number!")
	}
	if personCard.Balance < purchase.Sum {
		return errors.New("You have not enough money!")
	} else {
		personCard.Balance -= purchase.Sum
	}
	return nil
}
