package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName string
	Age      int
	FavNum   float64
	OwnADog  bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavNum = readFloat("What is you Favourite Number?")

	// fmt.Println("your name is "+userName+". You are", age, "years old")

	// fmt.Printf(fmt.Sprintf("Your name is %s and you are %d years old ", userName, age))

	fmt.Printf("Your name is %s and you are %d years old and your favourite is %.4f.",
		user.UserName,
		user.Age,
		user.FavNum)
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
