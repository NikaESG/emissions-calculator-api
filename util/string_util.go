package util

import (
	"bufio"
	"strings"
)

func ReadMultiLineString(str string) (stringList []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return stringList, err
	}
	return stringList, nil
}
