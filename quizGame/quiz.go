// quiz.go
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// ... (rest of the code)

func Game() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to start the game.")
	reader.ReadString('\n')

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("[1]Error while reading the file.")
		fmt.Println("Error: ", err)
		file.Close()
		return
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("[2]Error while reading the file.")
		fmt.Println("Error: ", err)
		return
	}

	var cor int = 0
	var i int = 1
	var numQues = len(rows)

	timer := time.NewTimer(time.Second * 5) // Set timer
	done := make(chan bool)
	timeUp := make(chan bool)

	go func() {
		// this function waits for this variable to be set
		// normally this variable is nil, when the timer is finished, it is set
		// so then this function will receive the set value as a channel
		// we did not need to initialize it (so to say)
		<-timer.C
		fmt.Println("\nTime's up!")
		// when the timer finishes, the chan variable is set to true
		// this variable is concurrently being checked at the start of each for loop
		// in the other concurrent function and also after the user inputs his answers.
		timeUp <- true
	}()

	go func() {
		for _, v := range rows {
			select {
			case <-timeUp:
				done <- true
				return
			default:
				fmt.Println("\nQuestion ", i, ": ", v[0])

				var a string
				fmt.Scan(&a)
				if a == v[1] {
					select {
					case <-timeUp:
						fmt.Printf("Correct! But time is up, so point not added.")
					default:
						fmt.Printf("Correct\n")
						cor++
					}
				} else {
					fmt.Printf("Incorrect %s\n", v[1])
				}
				fmt.Printf("\nScore: %d/%d", cor, numQues)
				i++
			}
		}
		done <- true
	}()

	<-done
}
