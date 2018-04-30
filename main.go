package main

import (
	"bufio"
	"fmt"
	"os"
)

func compareTwoStrings(s1 string, s2 string) int {
	score := 0;

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			score += 1;
		}
	}

	return score;
}

func findClosestString(in []string) string {
	minIdxScore := 0;
	minScore := 0;

	for i := 0; i < len(in); i++ {
		score := 0
		for j := 0; j < len(in); j++ {
			if i != j {
				score += compareTwoStrings(in[i], in[j])
			}
		}

		if minScore == 0 || score < minScore {
			minScore = score;
			minIdxScore = i;
		}
	}

	return in[minIdxScore];
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, _ := readLines("foo.in.txt")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	// numberOfShit := lines[0]
	actualLines := lines[1:]

	fmt.Println(findClosestString(actualLines));

	// fmt.Println(actualLines, numberOfShit);
}
