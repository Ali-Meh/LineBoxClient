package mcts

import (
	"math/rand"
	"time"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//RollOut calculates end state of the game randomly
func (n *Node) RollOut(count int) int {
	resChan := make(chan int)
	defer close(resChan)

	for i := 0; i < count; i++ {
		go evaluateRollOut(n.gmap.Clone(), n.turn, extractRemainingMoves(n.gmap), resChan)
	}
	value := 0
	for i := 0; i < count; i++ {
		t := <-resChan
		if value < t || i==0 {
			value = t
		}
	}
	return value
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
func evaluateRollOut(gmap gamemap.Map, turn bool, availableMoves [][]int8, resChan chan int) {
	/*select moves randomly*/
	//shuffle the moves
	rand.Seed(time.Now().UnixNano())
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
func evaluate(gmap gamemap.Map) int {
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
