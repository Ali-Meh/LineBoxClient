package mcts_test

import (
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestSelectMove(t *testing.T) {
	testmap := `2-1
0-0
@A@A@
-#B#-
@B@-@
-#A#A
@-@B@`

	//create map
	gmap := gamemap.NewMapSquare(2)
	gmap.Update(testmap)

	res := mcts.SelectMove(*gmap)

	assert.Equal(t, []int8{0, 1}, res)
}
