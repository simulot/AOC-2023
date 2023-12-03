package internal

import (
	"bufio"
	"os"
)

func ReadFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	o := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		o = append(o, s.Text())
	}
	return o, s.Err()
}
