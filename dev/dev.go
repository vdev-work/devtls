package dev

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// (INCLUSIVE)
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func Input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	line, err := reader.ReadString('\n')
	if err != nil {
		Error(err.Error())
	}

	return line
}
