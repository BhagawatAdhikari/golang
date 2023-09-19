package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready"

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)

}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {

	reader := bufio.NewReader(os.Stdin)

	// display instruction
	fmt.Println("Guess the number ")
	fmt.Println("--------------------------")
	fmt.Println("Think of a number between number from 1 to", prompt)
	reader.ReadString('\n')

	// take through the number

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the number by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by you orginally thought", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract the number by", substraction, prompt)
	reader.ReadString('\n')

	// give the answer

	fmt.Println("The answer is:", answer)
}
