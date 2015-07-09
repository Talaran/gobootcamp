package main
import (
	"github.com/golang/tour/wc"
	"strings"
)
func WordCount(s string) map[string]int {

	fields := strings.Fields(s)

	output := map[string]int{}

	for _, value := range fields {
		output[value]++
	}

	return output
}
func main() {
	wc.Test(WordCount)
}