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
func (n Node) UCB1() float64 {
	// return n.value/(n.visits+math.SmallestNonzeroFloat64) + uctk*math.Sqrt(2.*math.Log(t)/(n.visits+math.SmallestNonzeroFloat64))
	playerValue := 1.0
	// if n.turn {
	// 	playerValue = -1.0
	// }
	return playerValue*n.value/(n.visits+math.SmallestNonzeroFloat64) + uctk*math.Sqrt(2*math.Log(n.parentNode.visits)/(n.visits+math.SmallestNonzeroFloat64))

	if n.visits == 0 {
		return math.MaxFloat64
	}
	return n.value/(n.visits+math.SmallestNonzeroFloat64) + uctk*math.Sqrt(math.Log(n.parentNode.visits)/(n.visits+math.SmallestNonzeroFloat64))
}

//Expand will expand node with appropriate children
func (n *Node) Expand() *Node {
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
						if !clonedMap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(edgestate)) {
							turn = !turn
						}
						//generate node &&add to the children of existing node
						n.childNodes = append(n.childNodes, n.NewChild([]int8{edge.X, edge.Y}, turn, &clonedMap))
					}
				}
			}
		}
	}
	return n.childNodes[0]
}

//NewNode next move based on base state of the game
func NewNode(move []int8, turn bool, gmap *gamemap.Map) *Node {
	return &Node{move: move, value: 0, visits: 0, turn: turn, gmap: gmap}
}

//NewChild will Create a new child node for existing node
func (n *Node) NewChild(move []int8, turn bool, gmap *gamemap.Map) *Node {
	return &Node{move: move, value: 0, visits: 0, turn: turn, gmap: gmap, parentNode: n}
}

/***********Private Methods**********/

func (n *Node) hasChild(move []int8) bool {
	for _, c := range n.childNodes {
		if len(move) == len(c.move) && move[0] == c.move[0] && move[1] == c.move[1] {
			return true
		}
	}
	return false
}

func (n *Node) getBestChild() *Node {
	chossenNode := n.childNodes[0]
	chossenUcb := chossenNode.UCB1()
	//go to leafnode
	for _, n := range n.childNodes[1:] {
		v := n.UCB1()
		if v > chossenUcb {
			chossenUcb = v
			chossenNode = n
		}
	}
	return chossenNode
}
