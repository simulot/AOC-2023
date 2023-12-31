package internal

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadString(value string) []string {
	f := strings.NewReader(value)
	o := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		o = append(o, l)
	}
	return o
}

func ScanNumbers(l string) []int {
	ss := strings.Fields(l)
	numbers := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers[i] = n
	}
	return numbers
}

func FilterOutEmptyLines(lines []string) []string {
	r := []string{}
	for _, l := range lines {
		if len(l) > 0 {
			r = append(r, l)
		}
	}
	return r
}

func CSVSplit(s string) []string {
	return strings.Split(s, ",")
}
