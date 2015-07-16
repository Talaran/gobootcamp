package main
import (
	"fmt"
	"strings"
	"unicode/utf8"
)
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)
func main() {

	firstRune := func(str string) rune {

		var result rune

		result, _ = utf8.DecodeRune([]byte(str))

		return result
	}

	coinsForUser := func(user string) int {

		lowerUser := strings.ToLower(user)

		var userCoins int

		for _, character := range lowerUser {
			switch character {
			case firstRune("a"), firstRune("e"):
				userCoins+=1
			case firstRune("i"):
				userCoins+=2
			case firstRune("o"):
				userCoins+=3
			case firstRune("u"):
				userCoins+=4
			}
		}

		coins-= userCoins

		return userCoins
	}

	for _, user := range users {
		distribution[user] = coinsForUser(user)
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}