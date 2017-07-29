package sp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type YamlDataSet struct {
	Input *os.File
}

func (r YamlDataSet) String() string {
	yml, err := convert(r)
	if err != nil {
		return fmt.Sprintln(err)
	}
	return yml
}

func convert(r YamlDataSet) (string, error) {
	input := bufio.NewScanner(r.Input)
	input.Scan()
	header := strings.Split(input.Text(), "\t")

	var yml string

	for input.Scan() {
		if err := input.Err(); err != nil {
			return "", err
		}
		yml += fmt.Sprintln("-")
		cols := strings.Split(input.Text(), "\t")
		for i, v := range cols {
			yml += fmt.Sprintf("  %s: %s\n", header[i], quote(v))
		}
	}

	return yml, nil
}

func quote(s string) string {
	if _, err := time.Parse("20060102030405", s); err == nil {
		return strconv.Quote(s)
	}
	if _, err := time.Parse("20060102", s); err == nil {
		return strconv.Quote(s)
	}
	if _, err := strconv.Atoi(s); err == nil {
		return s
	}
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return s
	}
	return strconv.Quote(s)
}
