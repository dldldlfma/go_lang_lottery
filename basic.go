package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, _ := os.Open("./go_event.csv")

	rdr := csv.NewReader(bufio.NewReader(file))

	Participant, _ := rdr.ReadAll()

	fmt.Println("Number of Participant : ", len(Participant[0]))

	fmt.Println("Total Participant : ", Participant[0])

	fmt.Println("추첨을 시작합니다!")

	rand.Seed(time.Now().UnixNano())

	lottery_time := time.Date(2018, 11, 26, 8, 00, 0, 0, time.UTC)

	for {
		now := time.Now()
		diff := now.Sub(lottery_time)
		fmt.Println(diff.Seconds())
		if diff.Seconds() > 0 {
			for i := 0; i < 3; i++ {
				fmt.Println("축하드립니다! ", Participant[0][(rand.Intn(len(Participant[0])))], "님!")
				time.Sleep(1 * time.Second)
			}
			break
		}
	}

}
