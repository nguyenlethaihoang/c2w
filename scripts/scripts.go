package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	inputString := string(inputBytes)
	inputString = strings.ReplaceAll(inputString, "--request", "-X")
	inputString = strings.ReplaceAll(inputString, "--data", "-d")
	inputString = strings.ReplaceAll(inputString, "--header", "-H")
	inputString = strings.ReplaceAll(inputString, "'", "\"")
	inputString = strings.ReplaceAll(inputString, "\\", "")
	inputString = strings.ReplaceAll(inputString, "\t", "")
	inputString = strings.ReplaceAll(inputString, "\n", "")
	inputString = strings.ReplaceAll(inputString, "\r", "")

	index := 0
	newString := ""
	check := false
	for {
		if index >= len(inputString) {
			break
		}
		char := rune(inputString[index])
		if char == '{' {
			check = true
		}
		if char == '}' {
			check = false
		}
		if check == true && char == '"' {
			newString += `\"`
		} else {
			newString += string(char)
		}
		index++
	}

	err = ioutil.WriteFile("output.txt", []byte(newString), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Curl command converted and saved to output.txt")
}
