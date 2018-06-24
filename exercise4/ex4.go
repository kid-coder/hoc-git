package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func coinsForUser(user string) int {
	coin := 0
	for _, c := range user {
		switch c {
		case 'A', 'a', 'E', 'e':
			coin++
		case 'I', 'i':
			coin += 2
		case 'O', 'o':
			coin += 3
		case 'U', 'u':
			coin += 4
		}
	}
	return coin
}

func main() {
	for _, user := range users {
		c := coinsForUser(user)
		if c > 10 {
			c = 10
		}
		distribution[user] = c
		coins = coins - c
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
