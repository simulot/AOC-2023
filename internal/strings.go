package internal

import (
	"bufio"
	"strings"
)

func ReadString(value string) []string {
	f := strings.NewReader(value)
	o := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) > 0 {
			o = append(o, l)
		}
	}
	return o
}
