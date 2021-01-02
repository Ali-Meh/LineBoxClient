package gamemap

import "fmt"

//Tree Keeps track of game map evaluation
type Tree struct {
	move     []int8
	state    EdgeState
	children []*Tree
	parent   *Tree
	score    int
}

//AddChild will add new child to node
func (t *Tree) AddChild(c *Tree, score int) {
	if !t.checkExist(c) {
		c.score = score
		c.parent = t
		t.children = append(t.children, c)
	}
}

func (t *Tree) checkExist(c *Tree) bool {
	for _, n := range t.children {
		if len(n.move) == len(c.move) {
			if n.move[0] == c.move[0] && n.move[1] == c.move[1] {
				return true
			}
		}
	}
	return false
}

//UpdateScore will update the score of path
func (t *Tree) UpdateScore(score int) {
	t.score = score
}

//NewNode will create new node for the tree
func NewNode(move []int8, state EdgeState) *Tree {
	tree := new(Tree)
	tree.move = move
	tree.state = state
	tree.parent = nil
	return tree
}

func (t *Tree) String() string {
	queue := []*Tree{}
	queue = append(queue, t, nil, nil)
	res := ""
	for len(queue) > 1 {
		lastnode := queue[0]
		queue = queue[1:]

		//todo
		if lastnode == nil {
			if queue[0] == nil && len(queue) > 1 {
				queue = queue[1:]
				res += "\n"
			} else {
				res += "\t||\t"
			}
			continue
		} else {
			res += fmt.Sprintf("(%v , %d, %s)\t", lastnode.move, lastnode.score, lastnode.state)
			queue = append(queue, lastnode.children...)
			queue = append(queue, nil)
			if queue[0] != nil && queue[0].parent != lastnode.parent {
			}
			if len(queue) > 3 && queue[2] != nil && queue[2].parent.parent == lastnode.parent {
				queue = append(queue, nil)
			}
		}

	}
	// res += "\n"
	return res
}
