package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadCodesFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var codes []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		code := strings.TrimSpace(scanner.Text())
		if code != "" {
			codes = append(codes, code)
		}
	}
	return codes, scanner.Err()
}
