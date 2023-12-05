package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		unmodifiedline := scanner.Text()
		fmt.Println(unmodifiedline)

		wordNumberMap := extractWordsAndNumbersPart2(unmodifiedline)
		fmt.Println("red", wordNumberMap["red"], "green", wordNumberMap["green"], "blue", wordNumberMap["blue"], "total", wordNumberMap["red"]*wordNumberMap["green"]*wordNumberMap["blue"])
		total += (wordNumberMap["red"] * wordNumberMap["green"] * wordNumberMap["blue"])

	}

	fmt.Println("Total: ", total)
}

func extractWordsAndNumbersPart2(line string) map[string]int {
	result := make(map[string]int)

	patternGame := `Game (\d+): ([^;]+);`
	reGame := regexp.MustCompile(patternGame)
	matchGames := reGame.FindAllStringSubmatch(line, -1)
	matchGame := matchGames[0]
	numberGame, _ := strconv.Atoi(matchGame[1])
	result["Game"] = numberGame

	// Define a regular expression pattern for extracting words and numbers
	pattern := "\\b(\\d+)\\s*([a-zA-Z]+)\\b"
	re := regexp.MustCompile(pattern)

	// Find all matches in the input line
	matches := re.FindAllStringSubmatch(line, -1)

	// Process the matches and populate the map
	for _, match := range matches {
		if len(match) == 3 {
			// Convert the matched number to an integer
			number, err := strconv.Atoi(match[1])
			if err == nil {
				// Check if the word is already in the map
				if existingNumber, ok := result[match[2]]; ok {
					// Word is already in the map, update the number (e.g., add them)
					if number > existingNumber {
						result[match[2]] = number
					}
				} else {
					// Word is not in the map, add it with the new number
					result[match[2]] = number
				}
			}
		}
	}

	return result
}

func extractWordsAndNumbersPart1(line string) map[string]int {
	result := make(map[string]int)

	patternGame := `Game (\d+): ([^;]+);`
	reGame := regexp.MustCompile(patternGame)
	matchGames := reGame.FindAllStringSubmatch(line, -1)
	matchGame := matchGames[0]
	numberGame, _ := strconv.Atoi(matchGame[1])
	result["Game"] = numberGame

	// Define a regular expression pattern for extracting words and numbers
	pattern := "\\b(\\d+)\\s*([a-zA-Z]+)\\b"
	re := regexp.MustCompile(pattern)

	// Find all matches in the input line
	matches := re.FindAllStringSubmatch(line, -1)

	// Process the matches and populate the map
	for _, match := range matches {
		if len(match) == 3 {
			// Convert the matched number to an integer
			number, err := strconv.Atoi(match[1])
			if err == nil {
				// Check if the word is already in the map
				if existingNumber, ok := result[match[2]]; ok {
					// Word is already in the map, update the number (e.g., add them)
					if number > existingNumber {
						result[match[2]] = number
					}
				} else {
					// Word is not in the map, add it with the new number
					result[match[2]] = number
				}
			}
		}
	}

	return result
}
func checkPossible(gameMap map[string]int) bool {
	result := true
	if gameMap["red"] > 12 || gameMap["green"] > 13 || gameMap["blue"] > 14 {
		result = false
	}
	return result
}
