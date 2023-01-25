package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println("Welcome to Betrayal CommandLine Game!")
	//_ = chooseCharcters()
	game := newGame()
	game.showBoard()
}

const minNumCharacters = 3

func getRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func chooseCharacter(availableCharacters []Player) (Player, []Player) {
	fmt.Println("What is your name?")
	playerName := readInput("player name")

	fmt.Println("Your available characters are:")
	for _, character := range availableCharacters {
		fmt.Printf("%s\nMight: %d. Speed: %d. Sanity: %d, Knowledge: %d\n", character.CharacterName, character.Might, character.Speed, character.Sanity, character.Knowledge)
	}

	fmt.Println("Who do you want to be?")
	characterName := readInput("character name")

	var selectedCharacter Player

	remainingCharacters := make([]Player, 0)
	for _, character := range availableCharacters {
		if character.CharacterName == characterName {
			character.PlayerName = playerName
			selectedCharacter = character
		} else {
			remainingCharacters = append(remainingCharacters, character)
		}
	}

	return selectedCharacter, remainingCharacters
}

func getRandomCharacter(availableCharacters []Player) (Player, []Player) {
	randomNumber := getRandomNumber(0, len(availableCharacters))
	availableCharacters[randomNumber] = availableCharacters[len(availableCharacters)-1]
	return availableCharacters[randomNumber], availableCharacters[:len(availableCharacters)-1]
}

func readInput(whatAreYouReading string) string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading %s: %v", whatAreYouReading, err)
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

func chooseCharcters() []Player {
	fmt.Println("How many players?")
	input := readInput("number of players")
	numPlayers, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("number of players is in invalid format: %v", err)
	}

	allAvailableCharacters := getAllCharacters()
	players := make([]Player, 0)
	var newPlayer Player
	for i := 0; i < numPlayers; i++ {
		newPlayer, allAvailableCharacters = chooseCharacter(allAvailableCharacters)
		players = append(players, newPlayer)
	}

	if numPlayers < minNumCharacters {
		for i := numPlayers; i < minNumCharacters; i++ {
			newPlayer, allAvailableCharacters = getRandomCharacter(allAvailableCharacters)
			newPlayer.PlayerName = "Computer"
			players = append(players, newPlayer)
		}
	}

	fmt.Println("Great! Our players today are:")
	for _, player := range players {
		fmt.Printf("%s, played by: %s\n", player.CharacterName, player.PlayerName)
	}

	return players
}
