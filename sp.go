package sp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type YamlDataSet struct {
	Input io.Reader
}

func (yml YamlDataSet) String() string {
	s, err := yml.convert()
	if err != nil {
		return fmt.Sprintln(err)
	}
	return s
}

func (yml YamlDataSet) convert() (string, error) {
	var header []string
	var s string

	input := bufio.NewScanner(yml.Input)
	for input.Scan() {
		if err := input.Err(); err != nil {
			return "", err
		}
		columns := strings.Split(input.Text(), "\t")
		if len(header) == 0 || header[0] == "" {
			header = columns
			continue
		}
		s += fmt.Sprintln("-")
		for i, v := range columns {
			s += fmt.Sprintf("  %s: %s\n", header[i], quote(v))
		}
	}

	return s, nil
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
