package main

import "fmt"

func main() {
	var games map[int]string
	games = make(map[int]string)

	games[456] = "성기훈"
	games[218] = "박해수"
	games[067] = "강새벽"
	games[001] = "오일남"
	games[199] = "알리"
	games[414] = "장진석"

	for _, v := range games {
		fmt.Println(v)
	}

	games[414] = "장덕수"

	for k, v := range games {
		fmt.Println(k, v)
	}
}
