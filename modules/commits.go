package modules

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

var count = 0
var commit_array_size [15]string

func Position(commit_id, file_name string) {
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

		s := strings.Split(scanner.Text(), " ")
		if s[0] == commit_id {
			fmt.Println(s[0], "!! Bingo !!, Postion:", count)
		} else {
			//fmt.Println(strings.Join(s, " ")) implement if verbose is messages are needed
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