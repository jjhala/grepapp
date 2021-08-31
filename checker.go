package main

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	checkFunc func(keyWord, line string) (string, error)
)

func checkFull(content []string, keyWord string, checkFunc checkFunc, formatter formatter) ([]string, error) {
	// TODO: add test
	var detected []string
	for _, line := range content {
		res, err := checkFunc(keyWord, line)
		if err != nil {
			return nil, fmt.Errorf("error while check the line:%w: %s", err, line)
		}

		detected = append(detected, formatter(keyWord, res))
	}

	return detected, nil
}

func containsCheck(keyWord, line string) (string, error) {
	// TODO: add test

	if strings.Contains(line, keyWord) {
		return line, nil
	}

	// usually "else" keyword of what we don't really need
	return "", nil
}

/* func caseCheck(keyWord, line string) (string, error) {
	// TODO: add test
	if strings.Contains(strings.ToLower(line), strings.ToLower(keyWord)) {
		return line, nil
	}
	return "", nil
} */

func caseCheck(keyWord, line string) (string, error) {
	// TODO: add test
	j := 0
	lineArray := strings.Fields(line)
	for i, v := range lineArray {
		//println(v)
		//if v == keyWord {
		if strings.EqualFold(v, keyWord) {
			j += 1
			lineArray[i] = keyWord
		}
	}
	if j != 0 {
		return strings.Join(lineArray, " "), nil
	}
	return "", nil

}

func regCheck(keyWord, line string) (string, error) {

	jk, err := regexp.MatchString(keyWord, line)
	errorCheck(err)
	if jk {
		return line, nil
	}
	return "", nil

}
