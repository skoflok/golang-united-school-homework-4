package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	ss, err := getOperands(input)

	if err != nil {
		return "", err
	}

	ii, err := normalize(ss)

	if err != nil {
		return "", err
	}

	sum := ii[0] + ii[1]

	output = strconv.Itoa(sum)

	return output, nil
}

func isInt(r rune) bool {
	ints := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	return contains(ints, r)
}

func isOpt(r rune) bool {
	ops := []rune{'+', '-'}
	return contains(ops, r)
}
func isSpace(r rune) bool {
	spaces := []rune{'\t', '\n', '\v', '\f', '\r', ' '}
	return contains(spaces, r)
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func intContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getOperands(s string) (ss []string, e error) {
	s = strings.TrimSpace(s)
	if s == "" {
		e = fmt.Errorf("Empty input string: %w", errorEmptyInput)
		ss = make([]string, 0)
		return ss, e
	}
	e = nil
	r := []rune(s)
	need := false
	used := []int{}
	for i := 0; i < len(r); i++ {
		rr := []rune{}
		for i < len(r) && isInt(r[i]) == true {
			need = true
			if i > 0 && isOpt(r[i-1]) == true {
				rr = append(rr, r[i-1])
				used = append(used, i-1)
			}
			rr = append(rr, r[i])
			used = append(used, i)
			i++
		}
		if need == true {
			ss = append(ss, string(rr))
			need = false
		}
	}

	for i := 0; i < len(r); i++ {
		if intContains(used, i) {
			continue
		}

		if isSpace(r[i]) == true {
			continue
		}
		if isOpt(r[i]) == true {
			e = fmt.Errorf("Parse error:%d => %s", i, string(r[i]))
			ss = make([]string, 0)
			return ss, e
		}
	}

	if len(ss) != 2 {
		e = fmt.Errorf("Invalid operands count. Expected 2 operands. %w", errorNotTwoOperands)
		ss = make([]string, 0)
		return ss, e
	}
	return ss, e
}

func normalize(ss []string) (ii []int, e error) {
	first, e1 := strconv.Atoi(ss[0])
	second, e2 := strconv.Atoi(ss[1])

	if e1 != nil || e2 != nil {
		e = errors.New("Error normalize")
	}

	if e1 != nil {
		ii = append(ii, 0)
		ii = append(ii, 0)
		return ii, e1
	} else {
		ii = append(ii, first)
		ii = append(ii, second)
		return ii, e
	}
}
