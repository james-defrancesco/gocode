package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

var input int

func main() {
	fmt.Println("How many UUID's do you need? \nEnter any number:")
	fmt.Scanln(&input)
	displayKeys()
	fmt.Println("Here are your UUIDs! Enjoy :)")
}

// display the # of generated IDs
func displayKeys() int {
	var keys int

	for keys < input {
		uuid := generateUUID()
		fmt.Println(uuid)
		keys++
	}
	return keys
}

// generates UUID
func generateUUID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x", b[0:8])

	return uuid
}
