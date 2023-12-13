package main

import (
	"AOC2023/internal"
	"fmt"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("getSumOfDamageReports ", getSumOfDamageReports(readReports(lines)))
}

type report struct {
	status  string
	damaged []int
}

func readReports(lines []string) []report {
	reports := []report{}
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		s := strings.Split(l, " ")
		reports = append(reports, report{
			status:  s[0],
			damaged: internal.ScanNumbers(strings.Replace(s[1], ",", " ", -1)),
		})
	}
	return reports
}

func getSumOfDamageReports(reports []report) int {
	sum := 0
	for _, r := range reports {
		sum += len(getArrangements(r))
	}
	return sum
}

func getDamagedGroups(s string) []int {
	r := []int{}

	current := 0
	for i := range s {
		if s[i] == '.' {
			if current != 0 {
				r = append(r, current)
				current = 0
			}
		} else {
			current++
		}
	}
	if current > 0 {
		r = append(r, current)
	}
	return r
}

func getArrangements(report report) []string {
	stack := []string{report.status}
	results := []string{}

	for len(stack) > 0 {
		if strings.IndexByte(stack[0], '?') < 0 {
			d := getDamagedGroups(stack[0])
			if len(d) == len(report.damaged) {
				eq := true
				for i := range report.damaged {
					if d[i] != report.damaged[i] {
						eq = false
						break
					}
				}
				if eq {
					results = append(results, stack[0])
				}
			}
		} else {
			b := []byte(stack[0])
			for i := 0; i < len(b); i++ {
				if b[i] == '?' {
					b[i] = '#'
					stack = append(stack, string(b))
					b[i] = '.'
					stack = append(stack, string(b))
					break
				}
			}
		}
		stack = stack[1:]
	}
	return results
}
