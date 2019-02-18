package main

import (
	"fmt"
	"math"
	"math/rand"
)

const NumbersPerDie = 6
const NumDiePerSet = 5
const NumTurnsForYahtzee = 3

type Die struct {
	number int
}

func (self *Die) Roll() {
	// rand.Seed(time.Now().UnixNano())
	self.number = 1 + rand.Intn(NumbersPerDie)
}

type Dice []*Die

func newDice() Dice {
	dice := []*Die{}
	for i := 0; i < NumDiePerSet; i++ {
		dice = append(dice, &Die{})
	}
	return dice
}

func (self Dice) RollAll() {
	for _, die := range self {
		die.Roll()
	}
}

func (self Dice) RerollDiceNotMatchingMostCommonNumber() {
	maxNumber, _ := self.CountNumbers()
	for _, die := range self {
		if die.number != maxNumber {
			die.Roll()
		}
	}
}

func (self Dice) Print() {
	fmt.Println("printing dice")
	for _, die := range self {
		fmt.Printf("  -> %v\n", die.number)
	}
	mostCommonNumber, mostCommonCount := self.CountNumbers()
	fmt.Printf("  mostCommonNumber %v\n", mostCommonNumber)
	fmt.Printf("  mostCommonCount  %v\n", mostCommonCount)
}

func (self Dice) CountNumbers() (mostCommonNumber int, mostCommonCount int) {
	numbersByCount := map[int]int{}
	for _, die := range self {
		numbersByCount[die.number] += 1
	}
	mostCommonNumber = 0
	mostCommonCount = 0
	for number, count := range numbersByCount {
		if count >= mostCommonCount {
			mostCommonCount = count
			mostCommonNumber = number
		}
	}
	return
}

func (self Dice) HasYahtzee() bool {
	_, maxCount := self.CountNumbers()
	return maxCount == len(self)
}

func TryForYahtzee(dice Dice) bool {
	dice.RollAll()
	// dice.Print()
	turnsRemaining := NumTurnsForYahtzee - 1
	for i := 0; i < turnsRemaining; i++ {
		dice.RerollDiceNotMatchingMostCommonNumber()
		// dice.Print()
	}
	// fmt.Println("has a yahtzee??? ", dice.HasYahtzee())
	return dice.HasYahtzee()
}

func main() {
	dice := newDice()

	numAttempts := 1000000
	numYahtzees := 0
	for i := 0; i < numAttempts; i++ {
		if TryForYahtzee(dice) {
			numYahtzees++
		}
	}
	accuracyFactor := float64(10000)
	yahtzeePercent := math.Floor(100*accuracyFactor*(float64(numYahtzees)/float64(numAttempts))) / accuracyFactor
	fmt.Printf("with %v tries to get a yahtzee with %v dice, we have a yahtzee rate of %v%v (sample size %v)\n", NumTurnsForYahtzee, NumDiePerSet, yahtzeePercent, "%", numAttempts)
}
