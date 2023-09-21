package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName string
	Age      int
	FavNum   float64
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavNum = readFloat("What is you Favourite Number?")
	user.OwnsADog = readBool("Do you own a dog? (y/n)")

	// fmt.Println("your name is "+userName+". You are", age, "years old")

	// fmt.Printf(fmt.Sprintf("Your name is %s and you are %d years old ", userName, age))

	fmt.Printf("Your name is %s and you are %d years old and your favourite is %.4f. Owns a dog: %t.",
		user.UserName,
		user.Age,
		user.FavNum,
		user.OwnsADog)
}

func prompt() {
	fmt.Print(("-> "))
}

func readString(s string) string {
	for {

		fmt.Println((s))
		prompt()

		userInput := formatString()

		if userInput == "" {
			fmt.Println("Please enter the valid value")
		} else {
			return userInput
		}

	}
}

func formatString() string {
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func readInt(s string) int {

	for {
		fmt.Println(s)
		prompt()

		userInput := formatString()

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter the whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {

	for {
		fmt.Println(s)
		prompt()

		userInput := formatString()

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}

}
