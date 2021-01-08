package mcts

import (
	"math"
	"math/rand"
	"time"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

//Action to produce node
type Action []int8

// func (a Action) String() string {
// 	res := ""
// 	for _, v := range a {
// 		res += fmt.Sprint(v)
// 	}
// 	return res
// }

//Node will keep track of the game state
type Node struct {
	causingAction    Action
	remainingActions []Action
	parentNode       *Node
	children         []*Node
	value            float64
	visits           float64
	gmap             *gamemap.Map //state of the game
	depth            int          //depth of tree till this node
	turn             bool         //true if its maximzer turn
}

//UCB1 Calculates
func (n Node) UCB1(c float64) float64 {
	if n.visits == 0 || n.depth == 0 || n.visits <5 && n.depth != 0 && n.parentNode.depth == 0 {
		return math.Inf(1)
	}
	sign := 1.0
	if n.value < 0 {
		sign = -1.0
	}
	return sign * (((n.Eval()/float64(n.depth+1))/* *(c+1) */)*n.value/(n.visits) + c*math.Sqrt(math.Log(n.parentNode.visits)/(n.visits)))
}

//GetChildren Gets Node Children
func (n Node) GetChildren() []*Node {
	return n.children
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
						n.children = append(n.children, n.NewChild([]int8{edge.X, edge.Y}, turn, &clonedMap))
					}
				}
			}
		}
	}

	if len(n.children) > 0 {
		rand.Seed(time.Now().UnixNano())
		return n.children[rand.Intn(len(n.children))]
		// return n.childNodes[0]
	}
	return nil
}

//NewNode next move based on base state of the game
func NewNode(action Action, turn bool, gmap *gamemap.Map) *Node {
	return &Node{causingAction: action, value: 0, visits: 0, depth: 0, turn: turn, gmap: gmap, remainingActions: extractRemainingMoves(gmap)}
}

//NewChild will Create a new child node for existing node
func (n *Node) NewChild(action Action, turn bool, gmap *gamemap.Map) *Node {
	return &Node{causingAction: action, value: 0, visits: 0, turn: turn, gmap: gmap, parentNode: n, depth: n.depth + 1, remainingActions: extractRemainingMoves(gmap)}
}

/***********Private Methods**********/

func (n *Node) hasChild(action []int8) bool {
	for _, c := range n.children {
		if len(action) == len(c.causingAction) && action[0] == c.causingAction[0] && action[1] == c.causingAction[1] {
			return true
		}
	}
	return false
}

func (n *Node) getBestChild(c float64) *Node {
	chosenIndex := 0
	maxValue := -math.MaxFloat64
	for i, child := range n.children {
		if child.UCB1(c) > maxValue {
			maxValue = child.UCB1(c)
			chosenIndex = i
		}
	}
	return n.children[chosenIndex]
}

// IsLeaf returns true if the called-upon node is a leaf node in the tree false
// otherwise.
func (n Node) isLeaf() bool {
	return n.visits == 0 && (n.children == nil || len(n.children) == 0)
}

func (n *Node) isFullyExpanded() bool {
	return n.visits > 0 && n.children != nil
}

// IsTerminal conceptually differs from IsLeaf in that a node will be called
// "terminal" if it's domain state is terminal (end of the game), whereas IsLeaf
// returns true if it is merely the node's position in the tree that is terminal.
func (n Node) isTerminal() bool {
	return len(n.getRemainingMoves()) == 0
}

// IsRoot returns true if the called-upon node has no parent (and is in fact a
// root), false otherwise.
func (n Node) isRoot() bool {
	return n.parentNode == nil
}

func (n Node) getRemainingMoves() []Action {
	return extractRemainingMoves(n.gmap)
	// return n.remaining
}

func (n *Node) popAction() Action {
	action := n.remainingActions[0]
	n.remainingActions = n.remainingActions[1:]
	return action
}

func (n *Node) backpropagate(result float64) {
	for !n.isRoot() {
		n.value += float64(result) //* float64(n.parent.value.NextToMove())
		n.visits++
		n = n.parentNode
	}
	n.visits++
}
