package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var elfs [][]int
var elf []int
var largest int

func main() {
	file_path := os.Args[1]
	readFile, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file!!!")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	var elf_count int
  var elf_total int

	for _, line := range fileLines {
		// fmt.Println(line)

		if strings.TrimSpace(line) == "" {

			// fmt.Println("Elf", elf_count, get_elf_total(elf))
      elf_total = get_elf_total(elf)
      if elf_total > largest {
        largest = elf_total
      }
			// fmt.Println("Elf", elf_count, elf)

			elfs = append(elfs, elf)

			elf = nil
			elf_count += 1

		} else {

			clean_line, errr := strconv.Atoi(line)
			if errr != nil {
				fmt.Println("Error opening file!!!")
			}

			elf = append(elf, clean_line)
		}

		// fmt.Println(line, total)
	}
	// fmt.Println("Elfs", elfs)
  fmt.Println("Largest Elf", largest)
}

func get_elf_total(elf []int) int {
	var total int

	for _, x := range elf {
		total += x
	}

	return total
}
