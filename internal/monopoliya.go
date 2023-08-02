package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Participants struct {
	name        string
	money       int
	location    int
	ownedFields []int
}

type Fields struct {
	Location int
	Amount   int
	owner    string
}

func main() {
	fmt.Println("Enter the number of participants: ")
	var n int
	fmt.Scan(&n)
	Players := make([]Participants, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the name of the %v player: ", i+1)
		fmt.Scan(&Players[i].name)
		Players[i].money = 10000
	}
	var fields []Fields
	fields = fillingFields(fields)
	for {
		for i := 0; i < n; i++ {
			fmt.Println("Your turn :", Players[i].name)
			roll1 := rollDice()
			roll2 := rollDice()
			fmt.Printf("Вы выбросили %v и %v.\n", roll1, roll2)
			Players[i].location = (Players[i].location + roll1 + roll2) % len(fields)
			fmt.Println(Players[i].location)
			if Players[i].location != 0 && Players[i].location%10 != 0 {
				if fields[Players[i].location-1].owner == "" {
					fields[Players[i].location-1].owner = Players[i].name
					Players[i].ownedFields = append(Players[i].ownedFields, Players[i].location-1)
					fmt.Println("Вы становитесь хозяином поля:", fields[Players[i].location-1].Location)
				}
				if fields[Players[i].location-1].owner != Players[i].name {
					fmt.Printf("Вы попали на поле %v, которое принадлежит игроку %v.\n", fields[Players[i].location-1].Location, fields[Players[i].location-1].owner)
					payRent(&Players[i], fields[Players[i].location-1])
					//for i := 0; i < n; i++ {
					//	if Players[i].name == fields[(Players[i].location)-1].owner {
					//		Players[i].money += fields[(Players[i].location)-1].Amount
					//		fmt.Printf("Общая сумма хозяина %v поля стало: %v \n", Players[i].name, Players[i].money)
					//	}
					//}
				}
			}
			if Players[i].money <= 0 {
				fmt.Printf("%s выбывает из игры!\n", Players[i].name)
				break
			}
		}
		var counter int
		var winner string
		for i := 0; i < len(Players); i++ {
			if Players[i].money != 0 {
				counter++
				winner = Players[i].name
			}
		}
		if counter == 1 {
			fmt.Println("Победитель игры:", winner)
			return
		}
	}
}

func fillingFields(fields []Fields) []Fields {
	for i := 0; i < 40; i++ {
		newField := Fields{Location: i + 1, Amount: (i + 1) * 75, owner: ""}
		fields = append(fields, newField)
	}
	return fields
}

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
func payRent(player *Participants, field Fields) {
	if player.money > field.Amount {
		fmt.Println("У вас было:", player.money)
		player.money -= field.Amount
		fmt.Printf("Вы заплатили игроку %s сумму %d.\n У вас осталось: %v\n", field.owner, field.Amount, player.money)
	} else {
		fmt.Printf("У вас недостаточно денег для оплаты аренды! Вы выбываете из игры.\n")
		remainingAmount := field.Amount - player.money
		player.money = 0
		fmt.Printf("Вы должны банку сумму %d.\n", remainingAmount)
	}
}
