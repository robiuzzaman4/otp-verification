package main

import (
	"fmt"
	"math/rand"
)

// generate otp function
func generateOTP(length int) string {
	const numbers = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(otp)
}

func main() {
	otp := generateOTP(6)
	fmt.Println("Your OTP is: ", otp)
}
