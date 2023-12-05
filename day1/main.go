package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	part222()

}

// ___________________________________________________________________________________________________________________________________________________________??

var (
	digitList []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func part222() {
	var (
		ans int
		err error
	)

	fData, err := getFileData("input.txt")
	if err != nil {
		fmt.Printf("ERROR reading file:\n    %v", err)
	}

	ans, err = processData(fData)
	if err != nil {
		fmt.Printf("ERROR processing file data:\n    %v", err)
	}

	fmt.Println("=====================")
	fmt.Println(ans)
}

// processData -
func processData(fData []string) (ans int, err error) {
	var (
		numList []int
	)
	for i := 0; i < len(fData); i++ {

		tmpLn := fData[i]
		fNum := findFirstNum(tmpLn)
		lNum := findLastNum(tmpLn)
		cNumStr := fNum + lNum

		comboNumInt, err := strconv.Atoi(cNumStr)
		if err != nil {
			return ans, err
		}

		numList = append(numList, comboNumInt)
	}

	ans = addNumList(numList)

	return ans, err
}

func convDigitStr(digitStr string) (dStr string) {
	if len(digitStr) == 1 {
		return digitStr
	}
	switch digitStr {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	fmt.Printf("  -- ERROR: Did not find dStr ( %v ) --  \n", digitStr)
	return dStr
}

func findFirstNum(ln string) (fNum string) {
	var (
		fNumIdx int = len(ln) + 1
	)

	for d := 0; d < len(digitList); d++ {
		tmpIdx := strings.Index(ln, digitList[d])
		if tmpIdx != -1 {
			if tmpIdx < fNumIdx {
				fNumIdx = tmpIdx
				fNum = digitList[d]
			}
		}
	}

	return convDigitStr(fNum)
}

func findLastNum(ln string) (lNum string) {
	var (
		lNumIdx int = -1
	)

	for d := 0; d < len(digitList); d++ {
		tmpIdx := strings.LastIndex(ln, digitList[d])
		if tmpIdx != -1 {
			if tmpIdx > lNumIdx {
				lNumIdx = tmpIdx
				lNum = digitList[d]
			}
		}
	}

	if len(lNum) < 1 {
		fmt.Println(ln)
		fmt.Println(lNumIdx)
		fmt.Println(lNum)
		fmt.Println("something wrong happend to lNum ...")
	}
	return convDigitStr(lNum)
}

func addNumList(numList []int) (ttl int) {
	for i := 0; i < len(numList); i++ {
		//fmt.Println(numList[i])
		ttl = ttl + numList[i]
	}

	return ttl
}

// getFileData - read the data from the file and retun some sort of data structure...
func getFileData(file2read string) (fileData []string, err error) {
	inputF, err := os.Open(file2read)
	if err != nil {
		return fileData, err
	}
	defer inputF.Close()

	scn := *bufio.NewScanner(inputF)

	for scn.Scan() {
		ln := strings.TrimSpace(scn.Text())
		fileData = append(fileData, ln)
	}

	return fileData, err
}

// ___________________________________________________________________________________________________________________________________________________________??

func part22() {
	// Open the file
	file, err := os.Open("inputTest.txt")
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

		modifiedLine := replaceWordsWithNumbers(unmodifiedline)

		twoDigitNum, err := calculateTotal(modifiedLine)
		if err != nil {
			fmt.Println("Error calculating total:", err)
			return
		}

		// Update the total
		total += twoDigitNum

		// Print the new string and its corresponding number
		fmt.Println(modifiedLine+":", twoDigitNum)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the final total
	fmt.Println("Total:", total)
}

func getOccurrences(input string, replacements map[string]string) map[int]string {
	occurrences := make(map[int]string)

	for old := range replacements {
		index := strings.Index(input, old)
		for index != -1 {
			occurrences[index] = old

			// Find the next occurrence starting from the position after the last match
			nextIndex := strings.Index(input[index+len(old):], old)
			if nextIndex != -1 {
				// Adjust the index to the absolute position in the string
				index += len(old) + nextIndex
			} else {
				// No more occurrences
				break
			}
		}
	}

	return occurrences
}
func replaceWordsWithNumbers(input string) string {
	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	indexWordList := getOccurrences(input, replacements)

	// Extract the keys (indices) into a slice
	var keys []int
	for key := range indexWordList {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Ints(keys)

	// Print the sorted key-value pairs
	for _, key := range keys {
		input = strings.Replace(input, indexWordList[key], replacements[indexWordList[key]], -1)
		fmt.Printf("Index: %d, Word: %s\n", key, indexWordList[key])
	}

	return input
}

func calculateTotal(input string) (int, error) {
	var builder strings.Builder

	// Loop through each character in the line
	for _, char := range input {
		if unicode.IsDigit(char) {
			// Use strings.Builder for efficient string concatenation
			builder.WriteRune(char)
		}
	}

	// Get the final string
	numString := builder.String()

	// Check if the string is empty
	if numString == "" {
		return 0, nil
	}

	// Extract the first and last characters
	firstNum := numString[0]
	lastNum := numString[len(numString)-1]

	// Concatenate the characters to form a two-digit string
	twoDigitNumString := string(firstNum) + string(lastNum)

	// Convert the two-digit string to an integer
	twoDigitNum, err := strconv.Atoi(twoDigitNumString)
	if err != nil {
		return 0, err
	}

	return twoDigitNum, nil
}

func part2() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the file when the function exits
	defer file.Close()

	// Create a buffered reader
	scanner := bufio.NewScanner(file)
	total := 0
	// Read each line from the file
	for scanner.Scan() {
		unmodifiedline := scanner.Text()
		var builder strings.Builder
		modifiedLine1 := strings.Replace(unmodifiedline, "one", "1", -1)
		modifiedLine2 := strings.Replace(modifiedLine1, "two", "2", -1)
		modifiedLine3 := strings.Replace(modifiedLine2, "three", "3", -1)
		modifiedLine4 := strings.Replace(modifiedLine3, "four", "4", -1)
		modifiedLine5 := strings.Replace(modifiedLine4, "five", "5", -1)
		modifiedLine6 := strings.Replace(modifiedLine5, "six", "6", -1)
		modifiedLine7 := strings.Replace(modifiedLine6, "seven", "7", -1)
		modifiedLine8 := strings.Replace(modifiedLine7, "eight", "8", -1)
		modifiedLine9 := strings.Replace(modifiedLine8, "nine", "9", -1)
		// Loop through each character in the line
		for _, char := range modifiedLine9 {
			if unicode.IsDigit(char) {
				// Use strings.Builder for efficient string concatenation

				builder.WriteRune(char)
			}
		}
		// Get the final string
		numString := builder.String()

		firstNum := numString[0]
		lastNum := numString[len(numString)-1]
		// Concatenate the characters to form a two-digit string
		twoDigitNumString := string(firstNum) + string(lastNum)

		// Convert the two-digit string to an integer
		twoDigitNum, err := strconv.Atoi(twoDigitNumString)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		total += twoDigitNum
		// Print the new string
		fmt.Println(numString+":", twoDigitNum)

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Total:", total)
}

func part1() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the file when the function exits
	defer file.Close()

	// Create a buffered reader
	scanner := bufio.NewScanner(file)
	total := 0
	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		var builder strings.Builder
		// Loop through each character in the line
		for _, char := range line {
			if unicode.IsDigit(char) {
				// Use strings.Builder for efficient string concatenation

				builder.WriteRune(char)
			}
		}
		// Get the final string
		numString := builder.String()

		firstNum := numString[0]
		lastNum := numString[len(numString)-1]
		// Concatenate the characters to form a two-digit string
		twoDigitNumString := string(firstNum) + string(lastNum)

		// Convert the two-digit string to an integer
		twoDigitNum, err := strconv.Atoi(twoDigitNumString)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		total += twoDigitNum
		// Print the new string
		fmt.Println(numString+":", twoDigitNum)

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Total:", total)
}
