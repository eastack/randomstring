package randomstring

import (
	"math/rand"
)

const (
	Alphabet     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals     = "0123456789"
	Alphanumeric = Alphabet + Numerals
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

func Generate(length int, charset string) string {
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		buffer[i] = charset[rand.Intn(len(charset))]
	}
	return string(buffer)
}
