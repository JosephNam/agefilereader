package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	lineBytes, err := reader.ReadBytes('\n')

	ageMap := make(map[int]int)
	for len(lineBytes) > 0 && (err == nil || err.Error() == "EOF") {
		lineString := string(lineBytes)
		ageString := strings.Trim(strings.Trim(strings.Split(lineString, ",")[1], "\n"), " ")

		age, convertErr := strconv.Atoi(ageString)

		if convertErr != nil {
			fmt.Println("invalid input")
		}

		ageMap[age]++

		lineBytes, err = reader.ReadBytes('\n')
	}

	closeErr := file.Close()
	if closeErr != nil {
		panic(closeErr)
	}

	for key, value := range ageMap {
		fmt.Println("Age:", key, "Count:", value)
	}
}
