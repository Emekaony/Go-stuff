package third_party

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func getRandomProverb() string {
	return proverbs.Random().Saying
}

func Run() {
	fmt.Println(getRandomProverb())
}
