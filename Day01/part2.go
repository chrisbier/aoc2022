package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var elfs []int
var elf []int

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

	for _, line := range fileLines {
		// fmt.Println(line)

		if strings.TrimSpace(line) == "" {

			// fmt.Println("Elf", elf_count, get_elf_total(elf))
			// fmt.Println("Elf", elf_count, elf)

			elfs = append(elfs, get_elf_total(elf))

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
	sort.Ints(elfs[:])
	elfs_size := len(elfs)

	fmt.Println("Top 3 total:", get_elf_total(elfs[elfs_size-3:]))
}

func get_elf_total(elf []int) int {
	var total int

	for _, x := range elf {
		total += x
	}

	return total
}
