package mcts_test

import (
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestRollout(t *testing.T) {
	testmap := `2-1
0-0
@A@A@
B#A#-
@B@-@
A#A#-
@B@B@`

	//create map
	gmap := gamemap.NewMapSquare(2)
	gmap.Update(testmap)

	rootNode := mcts.NewNode([]int8{}, true, gmap)
	rootNode.RollOut(3)

	assert.Equal(t, 40, 40)
}
