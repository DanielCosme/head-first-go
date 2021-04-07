package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var voteCount map[string]int
	voteCount = make(map[string]int)

	lines, err := getStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	//names, count := sliceCount(lines)
	voteCount = mapCount(lines)
	//delete(voteCount, "Amber Graham")
	//delete(voteCount, "Daniel Cosme")

	var sortedVotes []string
	var sortedVotesInt []int
	reversedVotes := make(map[int]string)
	for name, count := range voteCount {
		reversedVotes[count] = name
		sortedVotesInt = append(sortedVotesInt, count)
		//sortedVotes = append(sortedVotes, name)
	}

	sort.Ints(sortedVotesInt)
	sort.Strings(sortedVotes)
	for _, name := range sortedVotesInt {
		fmt.Println(name, reversedVotes[name])
	}

	// for i, name := range names {
	// 	fmt.Printf("%s: %d\n", name, count[i])
	// }

}

func mapCount(votes []string) (count map[string]int) {
	count = make(map[string]int)

	for _, vote := range votes {
		count[vote]++
	}

	return count
}

func sliceCount(lines []string) (names []string, count []int) {
	match := false
	for _, line := range lines {
		for i, name := range names {
			if name == line {
				count[i]++
				match = true
			}
		}

		if match == false {
			names = append(names, line)
			count = append(count, 1)
		}
	}

	return names, count
}

func getStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, err
}
