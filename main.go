package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var perSalary, c int

func init() {
	flag.IntVar(&perSalary, "sal", 0, "時給")
	flag.IntVar(&c, "c", 1, "bool flag")
}

func weekCalc() (int, error) {
	var sum int
	week := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i := 0; i < 7; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(week[i], "の働いた時間(単位:分)を入力してください：")
		scanner.Scan()
		intInput := scanner.Text()
		intInput = strings.TrimSpace(intInput)
		workTime, err := strconv.Atoi(intInput)
		if err != nil {
			return 0, err
		} else {
			sum += workTime
		}
	}
	fmt.Println("今週のあなたの働いた時間(単位:分)は、", sum, "分です。")
	return sum, nil
}

func salaCalc(sumTime int) (int, error) {
	var salary int
	sumHour := sumTime / 60
	sumMinutes := sumTime - sumHour*60
	salary = sumHour*perSalary + int((float64(sumMinutes)/60)*float64(perSalary))
	return salary, nil
}

func main() {
	flag.Parse()
	if perSalary == 0 {
		fmt.Println("Usage: ./salaCalc -sal=[あなたの時給] (-c=[週の回数])")
	}
	sumTime, err := weekCalc()
	if err != nil {
		log.Fatal(err)
	}
	salary, err := salaCalc(sumTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("あなたの", c, "週間のお給料は", c*salary, "円です。")
}
