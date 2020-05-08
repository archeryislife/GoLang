package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lower    = "abcdedfghijklmnopqrstuvwxyz"
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specials = "~!@#$%^&*()_+-=[]{}|:,.<>/\\\""
	digits   = "0123456789"
	all      = lower + upper + specials + digits
)

func main() {
	i := 1
	for i <= 100000 {
		time.Sleep(1 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		minSpecialChar := 1
		minNum := 1
		minUpperCase := 1
		minLowerCase := 1
		passwordLength := 60
		password := GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase, minLowerCase)
		//Uncomment next line for DLP tracking tag
		//const dlp = "ENTER DLP STRING HERE"
		//Make next line fmt.Println(dlp + password)
		fmt.Println(password)
		i++
	}
}

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase, minLowerCase int) string {
	var password strings.Builder
	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specials))
		password.WriteString(string(specials[random]))
	}
	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(digits))
		password.WriteString(string(digits[random]))
	}
	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upper))
		password.WriteString(string(upper[random]))
	}
	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(upper))
		password.WriteString(string(upper[random]))
	}
	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase - minLowerCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(all))
		password.WriteString(string(all[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
