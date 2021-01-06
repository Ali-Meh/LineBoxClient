package mcts

import (
	"math/rand"
	"time"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//RollOut calculates end state of the game randomly
func (n *Node) RollOut(count int) float64 {

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
	// if n.depth == 1 {
	// actions := []Action{n.causingAction}
	// p := n.parentNode
	// for p.parentNode != nil {
	// 	actions = append(actions, p.causingAction)
	// 	p = p.parentNode
	// }
	// fmt.Printf("at %d  > %v evaled as : %f\n", n.depth, actions, maxRes)
	// }
	return maxRes + n.Eval() //hestoric outcome + current state seggested score
}

func extractRemainingMoves(gmap *gamemap.Map) []Action {
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

	availableMoves := []Action{}
	for _, v := range availableMovesMap {
		availableMoves = append(availableMoves, []int8{v.X, v.Y})
	}

	return availableMoves
}

func prioritiseActions(gmap gamemap.Map, availableMoves []Action) map[int][]Action {
	choisePriority := make(map[int][]Action, 0)

	max := 0
	for _, v := range availableMoves {
		for _, c := range gmap.GetEdgeCell(int(v[0]), int(v[1])) {
			if c.FilledEdgeCount > max {
				max = c.FilledEdgeCount
			}
		}
		choisePriority[max] = append(choisePriority[max], v)
		max = 0
	}
	return choisePriority
}

//evaluate
func evaluateRollOut(gmap gamemap.Map, turn bool, availableMoves []Action, resChan chan float64) {
	/*select moves randomly*/
	//shuffle the moves
	rand.Seed(time.Now().Local().UnixNano())

	choisePriority := prioritiseActions(gmap, availableMoves)

	for k := 0; k < 4; k++ {
		rand.Shuffle(len(choisePriority[k]), func(i, j int) {
			choisePriority[k][i], choisePriority[k][j] = choisePriority[k][j], choisePriority[k][i]
		})
	}

	//move the actions cousing cell to become 3 to least priority
	choisePriority[0], choisePriority[2] = choisePriority[2], choisePriority[0]

	//select 4s
	var edgestate string
	for k := 3; k >= 0; k-- {
		for _, v := range choisePriority[k] {
			if turn {
				edgestate = minimizerSambol
			} else {
				edgestate = maximizerSambol
			}
			if !gmap.SetEdgeState(int(v[0]), int(v[1]), gamemap.EdgeState(edgestate)) {
				turn = !turn
			}
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

//Eval will evaluate the score of the current terminal
func (n *Node) Eval() float64 {
	score := 0.0
	for _, raw := range n.gmap.Cells {
		for _, cell := range raw {
			switch cell.FilledEdgeCount {
			case 3:
				if n.turn {
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
