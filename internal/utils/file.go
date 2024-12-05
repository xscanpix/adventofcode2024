package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func SplitFileIntoColumnsFromFilename(filename string, separator string, columns int) *[][]int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return SplitFileIntoColumnsFromReader(file, separator, columns)
}

func SplitFileIntoColumnsFromReader(r io.Reader, separator string, columns int) *[][]int {
	scanner := bufio.NewScanner(r)

	result := make([][]int, columns)

	for scanner.Scan() {
		var values = strings.Split(scanner.Text(), separator)

		for i := range columns {
			val, err := strconv.Atoi(values[i])

			if err != nil {
				log.Fatal(err)
			}

			result[i] = append(result[i], val)
		}
	}

	return &result
}

func SplitFileIntoRowsFromFilename(filename string, separator string) *[][]int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return SplitFileIntoRowsFromReader(file, separator)
}

func SplitFileIntoRowsFromReader(r io.Reader, separator string) *[][]int {
	scanner := bufio.NewScanner(r)

	result := make([][]int, 0)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), separator)

		values := make([]int, 0)

		for i := range len(split) {
			val, err := strconv.Atoi(split[i])

			if err != nil {
				log.Fatal(err)
			}

			values = append(values, val)
		}

		result = append(result, values)
	}

	return &result
}
