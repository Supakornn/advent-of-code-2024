package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList := readInput("input.txt")

	leftListSorted := bubbleSort(leftList)
	rightListSorted := bubbleSort(rightList)

	distance := 0
	for i := 0; i < len(leftListSorted); i++ {
		distance += abs(leftListSorted[i] - rightListSorted[i])
	}

	fmt.Println("The total distance between the lists is:", distance)

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Println("The similarity score between the lists is:", similarityScore)
}

func readInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			leftNum, _ := strconv.Atoi(numbers[0])
			rightNum, _ := strconv.Atoi(numbers[1])
			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return leftList, rightList
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateSimilarityScore(leftList, rightList []int) int {
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	return similarityScore
}
