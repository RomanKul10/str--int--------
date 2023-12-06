package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("How old are you?")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n') //(`\n`) -гоурутина яка додає ще один рядок

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	age, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	result := ""

	if age >= 18 {
		result = "good"
	} else {
		result = "not good"
	}

	fmt.Print("You age - ", age, " years old. It`s  - ", result)

}
