package Utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func ContainsInStringArray(str string, s []string) bool {
	for _, str := range s {
		if str == str {
			return true
		}
	}
	return false
}
func RemoveFromStringArray(s string, slice []string) []string {
    var newSlice []string
    for _, str := range slice {
        if str != s {
            newSlice = append(newSlice, str)
        }
    }
    return newSlice
}
func RandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	s := base64.URLEncoding.EncodeToString(b)[:n]
	return s, nil
}
func ToHex(bin []byte) string {
	return hex.EncodeToString(bin)
}

func PrintlnAsHex(bin []byte, prefix string) {
	fmt.Println(prefix + ToHex(bin)+"\n")
}
