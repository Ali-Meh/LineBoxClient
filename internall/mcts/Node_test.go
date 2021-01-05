package mcts

import (
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	// "github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	tmap := `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`

	gmap := gamemap.NewMapSquare(2)
	maximizer := "A"
	if tmap[0] == '2' {
		maximizer = "B"
	}
	gmap.Update(tmap, maximizer)
	node := mcts.NewNode(nil, true, gmap)

	assert.Equal(t, 0, len(node.GetChildren()))

	child:=node.Expand()

	assert.Equal(t, 7, len(node.GetChildren()))


	child.

}
