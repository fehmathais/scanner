package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func GetString(txt string) (string, error) {
	print(txt)
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Trim(str, "\n"), err
}

func GetInt(txt string) (int, error) {
	print(txt)

	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return 0, nil
	}

	if n := rune(str[0]); !unicode.IsDigit(n) {
		return 0, fmt.Errorf("invalid input: you must to ender a valid digit number")
	}

	return strconv.Atoi(strings.Trim(str, "\n"))
}

func GetFloat(txt string) (float64, error) {
	print(txt)

	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(strings.Trim(str, "\n"), 64)
}

func print(txt string) {
	if len(txt) > 0 {
		fmt.Printf("%s", txt)
	}
}
