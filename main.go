package main

import "fmt"

var gameType int
var keywords []string
var nextKeywordIndex int

func main() {
	gameType = selectGame()
	var puzzle []string = generatePuzzle(gameType)
	printPuzzle(puzzle)
}

func savePuzzleToFile(puzzle string) {
	panic("unimplemented")
}

func printPuzzle(puzzle []string) {
	fmt.Println(puzzle)
}

func generatePuzzle(gameType int) []string {
	var puzzle []string

	//TODO: How to declare constant?
	var MAX_PUZZLE_WIDTH int = 50
	var MAX_PUZZLE_HEIGHT int = 50

	for height := MAX_PUZZLE_HEIGHT; height >= 0; height-- {
		keyword := getNextkeyword(gameType)

		rowStartLetters := getStartingRandomLetters()

		rowEndLettersCount := MAX_PUZZLE_WIDTH - len(keyword) - len(rowStartLetters)
		rowEndLetters := getEndingRandomLetters(rowEndLettersCount)

		puzzle = append(puzzle, rowStartLetters, keyword, rowEndLetters)

	}
	return puzzle
}

func getNextkeyword(gameType int) string {
	if keywords == nil {
		keywords = getKeywords(gameType)
	}

	if nextKeywordIndex < len(keywords) {
		keyword := keywords[nextKeywordIndex]
		nextKeywordIndex++
		return keyword

	} else {
		return ""
	}
}

func getStartingRandomLetters() string {
	return "SSSS"
}

func getEndingRandomLetters(letterCount int) string {
	return "E"
}

func getKeywords(gameType int) []string {
	return []string{"Sans Souci", "Ingleburn", "Kogarah", "Circual Quay", "Tempe", "Redfurn", "Wynyard"}
}

func selectGame() int {
	var gameType int

	displayGames("Suburbs", "Capitals", "Scientists")
	fmt.Print("\nSelect a game: ")

	//TODO: Handle error if user does not input int.
	fmt.Scan(&gameType)

	return gameType
}

func displayGames(gameNames ...string) {
	counter := 0
	for _, name := range gameNames {
		counter++
		fmt.Println(counter, ") ", name)
	}
}
