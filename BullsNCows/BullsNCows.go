package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomNum() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(9)
	return n
}

func randomNums() []int {
	for {
		count := 0
		nums := []int{randomNum(), randomNum(), randomNum(), randomNum()}

		if nums[0] != 0 {
			count++
			for g := range nums {
				for i := range nums {
					if nums[g] != nums[i] {
						count++
					}
				}
			}
		}
		if count == 13 {
			return nums
		}
	}
}

func readerCnsl() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Write in your guess number -> ")
		text, _ := reader.ReadString('\n')
		textR := text
		if len(textR) < 6 {
			text = ""
			if textR != "" {
				return textR
			}
		} else {
			fmt.Println("You put in too big input")
		}
	}
}
func decompGuessNum(s string) []int {
	splited := strings.Split(s, "")
	var numI []int
	numI = make([]int, 6)

	for i, r := range splited {
		numI[i], _ = strconv.Atoi(r)
	}
	return numI[0:4]
}

func compare(nums, guess []int) (cow, bull int) {
	for i, r := range guess {
		if guess[i] == nums[i] {
			bull++
		}
		for f := range nums {
			if r == nums[f] {
				cow++
			}
		}
	}
	cow = cow - bull
	return bull, cow
}

func conditions(bull, cow, turns int) {
	if bull == 4 {
		fmt.Println("")
		fmt.Printf("Cowgratulations, you won, it takes %d turns.", turns)
		fmt.Println("")
		os.Exit(1)
	} else {
		fmt.Println("")
		fmt.Printf("In your number %d cows and %d bulls", cow, bull)
		fmt.Println("")
	}
}
func main() {
	turns := 1
	fmt.Println("Number game Bools and Cows")
	nums := randomNums()
	fmt.Println("Nums are ready!")
	for {
		guessStr := readerCnsl()
		guess := decompGuessNum(guessStr)
		bull, cow := compare(nums, guess)
		conditions(bull, cow, turns)
		turns++
	}
}
