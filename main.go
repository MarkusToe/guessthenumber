package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess")

	secretNumber := generateRandomInteger(1, 100)
	fmt.Println("The secret number is", secretNumber)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}

	fmt.Println("Your guess is", guess)
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
