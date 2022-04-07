package repeater

import (
	"fmt"
	"os"
	"strings"
)

//This is a exercise function written in go
func Repeater() {
	msg := os.Args[1]
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	fmt.Println(s)
}

func RepeaterInFrontOfString() {
	msg := os.Args[1]
	l := len(msg)
	n := strings.Repeat("!", l)
	s := n + msg + n
	fmt.Println(s)
}

func RepeaterWithUppercase() {
	msg := os.Args[1]
	l := len(msg)
	n := strings.Repeat("!", l)
	s := strings.ToUpper(n + msg + n)
	fmt.Println(s)
}
