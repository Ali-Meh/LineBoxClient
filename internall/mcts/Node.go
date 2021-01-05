package mcts

import (
	"math"
	"math/rand"
	"time"

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
	depth      int
	turn       bool
}

//UCB1 Calculates
func (n Node) UCB1() float64 {

	if n.visits == 0 || n.depth == 0 || n.visits == 1 && n.depth != 0 && n.parentNode.depth == 0 {
		return math.Inf(1)
	}
	// if len(extractRemainingMoves(n.gmap)) == 0 && n.value != 0 {
	// 	return math.Inf(-1)
	// }
	// return n.value/(n.visits+math.SmallestNonzeroFloat64) + uctk*math.Sqrt(2.*math.Log(t)/(n.visits+math.SmallestNonzeroFloat64))
	playerValue := 1.0
	// if !n.turn {
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
						if clonedMap.SetEdgeState(int(edge.X), int(edge.Y), gamemap.EdgeState(edgestate)) {
							turn = !turn
						}
						//generate node &&add to the children of existing node
						n.childNodes = append(n.childNodes, n.NewChild([]int8{edge.X, edge.Y}, turn, &clonedMap))
					}
				}
			}
		}
	}

	if len(n.childNodes) > 0 {
		rand.Seed(time.Now().UnixNano())
		return n.childNodes[rand.Intn(len(n.childNodes))]
		// return n.childNodes[0]
	}
	return nil
}

//NewNode next move based on base state of the game
func NewNode(move []int8, turn bool, gmap *gamemap.Map) *Node {
	return &Node{move: move, value: 0, visits: 0, depth: 0, turn: turn, gmap: gmap}
}

//NewChild will Create a new child node for existing node
func (n *Node) NewChild(move []int8, turn bool, gmap *gamemap.Map) *Node {
	return &Node{move: move, value: 0, visits: 0, turn: turn, gmap: gmap, parentNode: n, depth: n.depth + 1}
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

	// maxUCB := math.Inf(-1)
	// maxima := n
	// maximaList := make([]*Node, 0)
	// if n.childNodes == nil {
	// 	return n
	// }
	// epsilon := 0.000001
	// //find the highest upper-confidence-bound in this node's children
	// for _, n := range n.childNodes {
	// 	ucb := n.UCB1()
	// 	// add selection bias for nodes containing states that specifiy it
	// 	bias := float64(0)
	// 	// if n.State != nil {
	// 	// 	bias = n.State.Bias()
	// 	// }
	// 	if ucb+bias >= maxUCB {
	// 		//compare floats within range of epsilon
	// 		if (maxUCB - ucb) <= epsilon {
	// 			maxUCB = ucb
	// 			maxima = n
	// 			maximaList = make([]*Node, 0)
	// 		}
	// 		maximaList = append(maximaList, n)
	// 	}
	// }
	// //if there is no true maximum, pick a random one
	// if len(maximaList) > 1 {
	// 	n := len(maximaList)
	// 	rand.Seed(time.Now().UTC().UnixNano())
	// 	i := rand.Intn(n)
	// 	return maximaList[i]
	// }
	// return maxima

	/******************************************************************/

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
	if chossenUcb == math.Inf(-1) {
		return n.parentNode.getBestChild()
	}
	return chossenNode
}


// IsLeaf returns true if the called-upon node is a leaf node in the tree false
// otherwise.
func (n Node) IsLeaf() bool {
	return n.childNodes == nil || len(n.childNodes) == 0
}

// IsTerminal conceptually differs from IsLeaf in that a node will be called
// "terminal" if it's domain state is terminal (end of the game), whereas IsLeaf
// returns true if it is merely the node's position in the tree that is terminal.
func (n Node) IsTerminal() bool {
	if len(n.extractRemainingMoves()) > 0 {
		return false
	}
	return true
}

// IsRoot returns true if the called-upon node has no parent (and is in fact a
// root), false otherwise.
func (n Node) IsRoot() bool {
	return n.parentNode == nil
}

func (n Node) extractRemainingMoves() [][]int8 {
	availableMovesMap := map[int]gamemap.Coordinates{}
	for _, raw := range n.gmap.Cells {
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
