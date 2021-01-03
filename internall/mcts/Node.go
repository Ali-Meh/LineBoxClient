package mcts

import (
	"math"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//Node will keep track of the game state
type Node struct {
	move       []int8
	parentNode *Node
	childNodes []*Node
	gmap       *gamemap.Map
	value      float64
	visits     float64
	turn       bool
}

//UCB1 Calculates
func (n *Node) UCB1(UCTK float64) float64 {
	return n.value/(n.visits+math.SmallestNonzeroFloat64) + UCTK*math.Sqrt(math.Log(n.parentNode.visits)/(n.visits+math.SmallestNonzeroFloat64))
}

// //RollOut calculates end state of the game randomly
// func (n *Node) RollOut() int {

// 	return 0
// }

//Expand will expand node with appropriate children
func (n *Node) Expand() {
	for i := range n.gmap.Cells {
		for _, cell := range n.gmap.Cells[i] {
			if cell.FilledEdgeCount < 4 {
				for _, edge := range cell.Edges {
					if edge.State == gamemap.IsFreeEdge && !n.hasChild([]int8{edge.X, edge.Y}) {
						clonedMap := n.gmap.Clone()
						turn := !n.turn

						var edgestate string
						if turn {
							edgestate = maximizerSambol
						} else {
							edgestate = minimizerSambol
						}

						//check if the board is filled then dont change turn//TODO test turn
						if clonedMap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(edgestate)) {
							turn = !turn
						}
						//generate node &&add to the children of existing node
						n.childNodes = append(n.childNodes, NewNode([]int8{edge.X, edge.Y}, turn, &clonedMap))
					}
				}
			}
		}
	}
}

//NewNode next move based on base state of the game
func NewNode(move []int8, turn bool, gmap *gamemap.Map) *Node {
	return &Node{move: move, value: 0, visits: 0, turn: turn, gmap: gmap}
}

/***********Private Methods**********/

func (n *Node) hasChild(move []int8) bool {
	for _, c := range n.childNodes {
		if len(move) == len(c.move) && n.move[0] == c.move[0] && n.move[1] == c.move[1] {
			return true
		}
	}
	return false
}
