package mcts

import (
	// "fmt"

	"math/rand"
	"time"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//RollOut calculates end state of the game randomly
func (n *Node) RollOut(count int) float64 {

	// return n.eval()
	// if n.depth == 1 && n.move[0]==3&&n.move[1]==2 {
	// 	fmt.Print()
	// }

	resChan := make(chan float64)
	defer close(resChan)

	for i := 0; i < count; i++ {
		go evaluateRollOut(n.gmap.Clone(), n.turn, extractRemainingMoves(n.gmap), resChan)
	}
	repeatCount := make(map[float64]int)
	value := <-resChan
	repeatCount[value]++
	maxRes := value
	for i := 0; i < count-1; i++ {
		value = <-resChan
		repeatCount[value]++
		if repeatCount[maxRes] < repeatCount[value] {
			maxRes = value
		}
	}
	return maxRes + n.eval()*float64(n.depth*2)
}

func extractRemainingMoves(gmap *gamemap.Map) [][]int8 {
	availableMovesMap := map[int]gamemap.Coordinates{}
	for _, raw := range gmap.Cells {
		for _, cell := range raw {
			for _, e := range cell.Edges {
				if e.State == gamemap.IsFreeEdge {
					availableMovesMap[int(e.X*10+e.Y)] = e.Coordinates
				}
			}
		}
	}

	availableMoves := [][]int8{}
	for _, v := range availableMovesMap {
		availableMoves = append(availableMoves, []int8{v.X, v.Y})
	}

	return availableMoves
}

//evaluate
func evaluateRollOut(gmap gamemap.Map, turn bool, availableMoves [][]int8, resChan chan float64) {
	/*select moves randomly*/
	//shuffle the moves
	rand.Seed(time.Now().Local().UnixNano())
	rand.Shuffle(len(availableMoves), func(i, j int) { availableMoves[i], availableMoves[j] = availableMoves[j], availableMoves[i] })
	//select moves
	var edgestate string
	for _, v := range availableMoves {
		if turn {
			edgestate = maximizerSambol
		} else {
			edgestate = minimizerSambol
		}
		if !gmap.SetEdgeState(int(v[0]), int(v[1]), gamemap.EdgeState(edgestate)) {
			turn = !turn
		}
	}

	resChan <- evaluate(gmap)
}

//Evaluate will evaluate the score of the current terminal
func evaluate(gmap gamemap.Map) float64 {
	score := 0.0
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
	// if score > 0 {
	// 	return 1
	// }
	// 	return 0

	return score
}

//Evaluate will evaluate the score of the current terminal
func (n *Node) eval() float64 {
	score := 0.0
	for _, raw := range n.gmap.Cells {
		for _, cell := range raw {
			switch cell.FilledEdgeCount {
			case 3:
				if n.parentNode.turn {
					score--
				} else {
					score++
				}
			case 4:
				if maximizerSambol == string(cell.OwnedBy) {
					score++
				} else {
					score--
				}
			}
		}
	}
	// if score > 0 {
	// 	return 1
	// }
	// 	return 0

	return score
}
