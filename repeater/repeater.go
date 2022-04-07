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
