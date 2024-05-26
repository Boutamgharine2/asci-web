package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Printing(s, r string) string {
	fmt.Println("==>", s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if Checkinput(s) {
		style := r + ".txt"

		Filestyle, _ := os.Open(style)
		v := strings.Split(s, "\n")
		h := PrintN(Filestyle, v)
		return h
	}
	return "you need to choose a character from the ascii table."
}
