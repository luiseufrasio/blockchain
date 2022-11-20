package main

import "time"

func main() {
	go genABC()

	time.Sleep(60 * time.Millisecond)
}

func genABC() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}