package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an integer value: ")
	input, _ := reader.ReadString('\n')

	converted, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	fmt.Println("The value + 1 is: ", converted+1)

}
