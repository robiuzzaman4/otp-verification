package main

import (
	"fmt"
	"math/rand"
	"strings"
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

// verify OTP function
func verifyOTP(generatedOTP, enteredOTP string) bool {
	return strings.Compare(generatedOTP, enteredOTP) == 0
}

func main() {
	otp := generateOTP(6)
	fmt.Println("Your OTP is: ", otp)

	var enteredOTP string
	fmt.Print("Enter the OTP to verify: ")
	fmt.Scanln(&enteredOTP)

	// Verify the OTP
	if verifyOTP(otp, enteredOTP) {
		fmt.Println("OTP Verified Successfully!")
	} else {
		fmt.Println("Invalid OTP!")
	}
}
