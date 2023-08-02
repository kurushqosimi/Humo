package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func File() {
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
	file, err2 := os.Create("data.json")
	if err2 != nil {
		log.Println(err)
		return
	}

	defer file.Close()
	file.WriteString(string(bytes))
	fmt.Println("Done.")
}
