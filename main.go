package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var count = 0
var commit_array_size [15]string
var provided_args []string

func main() {
	provided_args := os.Args[1:]
	automater_mode(provided_args)
}

func automater_mode(provided_args []string) {
	if "find" == provided_args[0] {
		find_commit_order(provided_args[1], provided_args[2])
	} else if "merge" == provided_args[0] {
		git_cherry_automater(provided_args[1])
	}

}
func git_cherry_automater(provided_commit string) {
	fmt.Println(" merge " + provided_commit)
}

func find_commit_order(commit_id, file_name string) {
	fmt.Println("  ############# Commit order############# ")
	fmt.Println(file_name)
	data, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		//find_commit(scanner.Text())
		s := strings.Split(scanner.Text(), " ")
		if s[0] == commit_id {
			fmt.Println(s[0], "bingo !!!, postion:", count)
		} else {
			fmt.Println(s[0])
		}

		commit_array_size[count] = scanner.Text()
		count++

		if count > 14 {
			break
		}
	}
	fmt.Println("  ############# End ############# ")

}
