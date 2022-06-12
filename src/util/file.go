package util

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func Read(inputFile string) string {
	s := ""

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	i := 0
	for scanner.Scan() {
		if i > 0 {
			s += "\n"
		}
		s += scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}
