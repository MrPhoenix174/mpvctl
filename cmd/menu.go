package cmd

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(result)
}

func GetUserChoice() int {
	fmt.Print("> ")
	var choice int
	fmt.Scan(&choice)
	return choice
}
