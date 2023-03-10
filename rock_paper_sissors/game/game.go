package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	// Print to screen
	// Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	// g.DisplayChan <- "Rock, Paper & Scissors"
	// g.DisplayChan <- "----------------------"
	// g.DisplayChan <- "Game is played for three rounds, and best of three wins the game. Good luck!"
	// g.DisplayChan <- ""
	g.DisplayChan <- fmt.Sprintln(`
	Rock, Paper & Scissors
	----------------------
	Game is played for three rounds, and best of three wins the game. Good luck!
	
	`)
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	g.DisplayChan <- ""
	g.DisplayChan <- fmt.Sprintln("Round", g.Round.RoundNumber)
	g.DisplayChan <- "-------"

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- ""
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
	default:
	}

	g.DisplayChan <- ""

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			g.DisplayChan <- "Invalid choice!"
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- "Final score"
	g.DisplayChan <- "-----------"
	g.DisplayChan <- fmt.Sprintf("Player: %d/3, Computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	g.DisplayChan <- ""

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins game!"
	} else {
		g.DisplayChan <- "Computer wins game!"
	}
}
