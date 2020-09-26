package ai

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//Evaluate will evaluate the score of the current map
//will return for +10 for every 3
func Evaluate(gamemap gamemap.Map, maximizingTurn bool) int {
	score := 0
	fmt.Println(gamemap)
	return score
}
