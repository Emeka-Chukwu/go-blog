package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

/// RandomInt: generate a random int between max and minimum

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

///RandomString; generate a a random string of length n

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	// fmt.Println(sb.String())
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomPassword() string {
	return "@123" + RandomString(6)
}

func role() string {
	return "user"
}

func RandomNumber() int64 {
	return RandomInt(0, 1000000000)
}
