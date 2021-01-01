package ai

import (
	"fmt"
	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string = "A"
	minimizerSambol string = "B"
)

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
	//TODO for those are 3 dim find if the others
	for _, raw := range gmap.Cells {
		for _, cell := range raw {
			switch cell.FilledEdgeCount {
			case 3:
				if maximizingTurn {
					score += 20
				} else {
					score -= 20
				}
			case 4:
				if maximizerSambol == string(cell.OwnedBy) {
					score += 20
				} else {
					score -= 20
				}
			}
			// if cell.FilledEdgeCount == 3 {
			// 	if maximizingTurn {
			// 		score += 20
			// 	} else {
			// 		score -= 20
			// 	}
			// }
			// if cell.FilledEdgeCount == 4 {
			// 	if maximizerSambol == string(cell.OwnedBy) {
			// 		score += 20
			// 	} else {
			// 		score -= 20
			// 	}
			// }
		}
	}
	// fmt.Println("EVALUATING MAP: ", score, maximizingTurn)
	// fmt.Println(gmap)
	// fmt.Println(score)
	// if maximizingTurn {
	return score
	// }
	// return -score
}

//MiniMax Algo Implementation
func MiniMax(gmap gamemap.Map, depth int, maximizingTurn bool, alpha, beta int) int {
	if depth == 0 || !gmap.HasFreeEdge() {
		eval := Evaluate(gmap, maximizingTurn)
		// fmt.Println("**************final map*************")
		// fmt.Println(gmap)
		// fmt.Println("depth: ", depth, "has free: ", gmap.HasFreeEdge())
		// fmt.Println("is A:", maximizingTurn)
		// fmt.Println("eval: ", eval)
		// fmt.Println("**********************************************")

		return eval
	}
	turn := !maximizingTurn
	if maximizingTurn {
		bestVal := -99999999
		for i := range gmap.Cells {
			for _, cell := range gmap.Cells[i] {
				if cell.FilledEdgeCount < 4 {
					for _, edge := range cell.Edges {
						if edge.State == gamemap.IsFreeEdge {
							// fmt.Println("depth:", depth, "MAX 🔼 Chose:", edge.Coordinates)
							clonedMap := gmap.Clone()
							// cell.FilledEdgeCount++
							if clonedMap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(maximizerSambol)) {
								// clonedMap.Cells[i][j].OwnedBy = gamemap.IsAEdge
								turn = !turn
							}
							score := MiniMax(clonedMap, depth-1, turn, alpha, beta)
							turn = !maximizingTurn
							bestVal = max(score, bestVal)
							alpha = max(alpha, score)
							if beta <= alpha {
								break
							}
						}
					}
				}
			}
		}
		// if depth > 1 {
		// 	fmt.Println(gmap)
		// 	fmt.Println("depth: ", depth)
		// 	fmt.Println("Maximizer🔼")
		// 	fmt.Println("returning", bestVal)
		// }
		return bestVal
	} else {
		bestVal := 99999999
		for i := range gmap.Cells {
			for _, cell := range gmap.Cells[i] {
				if cell.FilledEdgeCount < 4 {
					for _, edge := range cell.Edges {
						if edge.State == gamemap.IsFreeEdge {
							// fmt.Println("depth:", depth, "MIN 🔻 Chose:", edge.Coordinates)
							clonedMap := gmap.Clone()

							// cell.FilledEdgeCount++
							if clonedMap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(minimizerSambol)) {
								// clonedMap.Cells[i][j].OwnedBy = gamemap.IsBEdge
								turn = !turn
							}
							score := MiniMax(clonedMap, depth-1, turn, alpha, beta)
							turn = !maximizingTurn
							bestVal = min(score, bestVal)
							beta = min(beta, score)
							if beta <= alpha {
								break
							}
						}
					}
				}
			}
		}
		// if depth > 2 {
		// 	fmt.Println(gmap)
		// 	fmt.Println("depth: ", depth)
		// 	fmt.Println("MINIMIZER 🔻")
		// 	fmt.Println("returning", bestVal)
		// }
		return bestVal
	}
}

//SelectMove find and retrun best move
func SelectMove(gmap gamemap.Map, depth int, maximizer string) []int8 {
	maximizerSambol = maximizer
	if maximizerSambol == "A" {
		minimizerSambol = "B"
	} else {
		minimizerSambol = "A"
	}

	bestVal := -99999999
	turn := false
	move := []int8{0, 0}
	for i := range gmap.Cells {
		for _, cell := range gmap.Cells[i] {
			if cell.FilledEdgeCount < 4 {
				for _, edge := range cell.Edges {
					if edge.State == gamemap.IsFreeEdge {
						// fmt.Printf("========================>the edge is : %v \n", edge.Coordinates)
						clonedmap := gmap.Clone()

						// cell.FilledEdgeCount++
						if clonedmap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(maximizerSambol)) {
							// clonedmap.Cells[i][j].OwnedBy = gamemap.IsAEdge
							turn = !turn
						}
						score := MiniMax(clonedmap, depth, turn, -999999, 999999)
						turn = false
						if score > bestVal {
							bestVal = score
							move = []int8{edge.X, edge.Y}
						}
					}
				}
			}
		}
	}
	fmt.Println("selecting: [", move[0], ",", move[1], "] with score: ", bestVal)
	return move
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
