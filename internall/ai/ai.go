package ai

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//Evaluate will evaluate the score of the current map
//will return for +10 for every 3 filled
func Evaluate(gamemap gamemap.Map, maximizingTurn bool) int {
	score := 0
	for _, raw := range gamemap.Cells {
		for _, cell := range raw {
			if cell.FilledEdgeCount == 3 {
				score += 10
			}
		}
	}
	fmt.Println(gamemap)
	return score
}
