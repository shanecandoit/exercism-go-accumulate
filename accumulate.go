package accumulate

import (
	"fmt"
	"strings"
)

// Accumulate given a collection and
// an operation to perform on each element of the collection,
// returns a new collection containing the results
// of applying that operation to each element of the input collection.
func Accumulate(given []string, converter func(string) string) (results []string) {

	for _, s := range given {
		newS := converter(s)
		results = append(results, newS)
	}

	return results
}

func main() {
	given := []string{"cat", "Dog", "b4t", "gO"}
	op := strings.ToUpper
	expect := []string{"CAT", "DOG", "B4T", "GO"}
	got := Accumulate(given, op)
	fmt.Println("expect", expect)
	fmt.Println("got", got)
}
