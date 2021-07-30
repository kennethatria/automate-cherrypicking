package automate

func git_cherry_automater(provided_args []string) {

}

/*func cherry_automater(provided_args string[]) {
	count = 0
	var commits_no_limited [15]int
	//var expected_args = strings.Join(provided_args)
	data, err := os.Open(provided_args)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	for i := 0; i < 15; i++ {
		for scanner.Scan() {
			count++
			//commits_no_limited[count] = scanner.Text()
		}
	}

	//fmt.Printf("number of commits %d\nprovided arguments\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
} */
