package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	strArr := strings.Split(s, "")
	strArr = strArr[0 : len(strArr)-1]

	for i := 0; i < len(strArr); i++ {
		j := len(strArr) - 1 - i
		if strArr[i] != strArr[j] {
			return false
		}
	}

	return true
}

func main() {
	var isPalindrome bool

	fmt.Print("Masukkan teks: ")
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')

	isPalindrome = CheckPalindrome(input)
	fmt.Println(isPalindrome)
}
