package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count := 0
	data, err := os.Open("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}

	fmt.Printf("number of commits %d\n provided arguments\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
