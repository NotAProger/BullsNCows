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

func info() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("")
	fmt.Println("This game (as you can guess) is number game")
	fmt.Println("Concept is simple - one person (or bot in our case) thinks hiden number, other person guess it.")
	fmt.Println("Rules: In number 4 (in our case) unick uint numerals, guessperson shuld guess it for minimum rounds")
	fmt.Println("if guess numeral equal to one of hiden numerals and on the same position - its Bull")
	fmt.Println("if guess numeral equal to one of hiden numerals but on other position - its cow.")
	time.Sleep(time.Second * 4)
	fmt.Println("")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("So.... lets start")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Genering random nums... usualy its take ~5 sec....")
}

func randomNum() int {
	rand.Seed(time.Now().UnixNano())
	n := 1 + rand.Intn(10)
	if n == 10 {
		n = 0
	}
	return n
}

func randomNums() []int {
	for {
		one := randomNum()
		two := randomNum()
		three := randomNum()
		four := randomNum()

		if one != 0 && one != two && one != three && one != four {
			if two != three && two != four {
				if three != four {
					return []int{one, two, three, four}
				}
			}
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
	temp := [2]int{}
	for i := range guess {
		if guess[i] == nums[i] {
			bull++
		}
	}

	for i, r := range guess {
		temp[0] = i
		for f, g := range nums {
			temp[1] = g
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
	info()
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
