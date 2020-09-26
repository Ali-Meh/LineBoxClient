package ai

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//Evaluate will evaluate the score of the current map
//will return for +10 for every 3 filled
func Evaluate(gmap gamemap.Map, maximizingTurn bool, filler string) int {
	score := 0
	//TODO for those are 3 dim find if the others
	for _, raw := range gmap.Cells {
		for _, cell := range raw {
			if cell.FilledEdgeCount == 3 {
				score += 10
			}
			if cell.FilledEdgeCount == 4 {
				if filler == string(cell.OwnedBy) {
					score += 20
				} else {
					score -= 20
				}
			}
		}
	}
	fmt.Println(gmap)
	// if maximizingTurn {
	return score
	// }
	// return -score
}

//MiniMax Algo Implementation
func MiniMax(gmap gamemap.Map, depth int, maximizingTurn bool, alpha, beta int) int {
	if depth == 0 || !gmap.HasFreeEdge() {
		if maximizingTurn {
			return Evaluate(gmap, maximizingTurn, "A")
		}
		return Evaluate(gmap, maximizingTurn, "B")
	}
	turn := !maximizingTurn
	if maximizingTurn {
		bestVal := -99999999
		for i := range gmap.Cells {
			for _, cell := range gmap.Cells[i] {
				if cell.FilledEdgeCount < 4 {
					for _, edge := range cell.Edges {
						if edge.State == gamemap.IsFreeEdge {
							edge.State = gamemap.IsAEdge
							cell.FilledEdgeCount++
							if cell.FilledEdgeCount == 4 {
								cell.OwnedBy = edge.State
								turn = !turn
							}
							score := MiniMax(gmap, depth-1, turn, alpha, beta)
							bestVal = max(score, bestVal)
							alpha = max(alpha, bestVal)
							if beta <= alpha {
								break
							}
						}
					}
				}
			}
		}
		return bestVal
	} else {
		bestVal := 99999999
		for i := range gmap.Cells {
			for _, cell := range gmap.Cells[i] {
				if cell.FilledEdgeCount < 4 {
					for _, edge := range cell.Edges {
						if edge.State == gamemap.IsFreeEdge {
							edge.State = gamemap.IsBEdge
							cell.FilledEdgeCount++
							if cell.FilledEdgeCount == 4 {
								cell.OwnedBy = edge.State
								turn = !turn
							}
							score := MiniMax(gmap, depth-1, turn, alpha, beta)
							bestVal = min(score, bestVal)
							beta = min(beta, bestVal)
							if beta <= alpha {
								break
							}
						}
					}
				}
			}
		}
		return bestVal
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
