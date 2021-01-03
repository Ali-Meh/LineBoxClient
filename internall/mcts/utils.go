package mcts

import "github.com/ali-meh/LineBoxClient/internall/gamemap"

//Evaluate will evaluate the score of the current map
//will return for +10 for every 3 filled
func Evaluate(gmap gamemap.Map, maximizingTurn bool, maximizer ...string) int {
	if maximizer != nil {
		maximizerSambol = maximizer[0]
		if maximizerSambol == "A" {
			minimizerSambol = "B"
		} else {
			minimizerSambol = "A"
		}
	}
	score := 0
	for _, raw := range gmap.Cells {
		for _, cell := range raw {
			switch cell.FilledEdgeCount {
			case 4:
				if maximizerSambol == string(cell.OwnedBy) {
					score++
				} else {
					score--
				}
			}
		}
	}
	return score
}
