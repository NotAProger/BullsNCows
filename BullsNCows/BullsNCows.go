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

func randomNum(a, n int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(n) + a
	return r
}

func isValid(nums []int) bool {
	for g := range nums {
		for i := range nums {
			if nums[g] == nums[i] {
				return false
			}
		}
	}
	return true
}
func randomNums() []int {
	b := false
	var nums []int
	for {
		nums = []int{randomNum(1, 8), randomNum(0, 9), randomNum(0, 9), randomNum(0, 9)}
		if b == true {
			break
		}
		b = isValid(nums)
	}
	fmt.Println(nums)
	return nums
}

func readerCnsl() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Write in your guess number -> ")
		text, _ := reader.ReadString('\n')
		textR := text
		text = ""
		if len(textR) < 7 && textR != "" {
			return textR
		} else if textR == "" {
			fmt.Println("Put in some number")
		} else {
			fmt.Println("You put in too big number")
		}
	}
}
func decompGuessNum(s string) []int {
	splited := strings.Split(s, "")
	var numI []int
	numI = make([]int, 7)

	for i, r := range splited {
		numI[i], _ = strconv.Atoi(r)
	}
	return numI[0:4]
}

func compare(nums, guess []int, t int) bool {
	bull := 0
	cow := 0

	for i, g := range guess {
		for j, n := range nums {
			if g == n {
				if i == j {
					bull++
				} else {
					cow++
				}
			}
		}
	}

	if bull == 4 {
		fmt.Println("")
		fmt.Printf("Cowgratulations, you won, it takes %d turns.", t)
		fmt.Println("")
		return true
	} else {
		fmt.Println("")
		fmt.Printf("In your number %d cows and %d bulls", cow, bull)
		fmt.Println("")
		return false
	}
}

func fin() {
	turns := 0
	fmt.Println("Number game Bools and Cows")
	nums := randomNums()
	fmt.Println("Nums are ready!")
	for {
		turns++
		guessStr := readerCnsl()
		guess := decompGuessNum(guessStr)
		b := compare(nums, guess, turns)
		if b == true {
			break
		}
	}
}
func main() {
	fin()
}
