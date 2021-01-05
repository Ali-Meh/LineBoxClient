package mcts

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	// "github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestExpandNode(t *testing.T) {
	testmap := []struct {
		tmap string
	}{
		{
			tmap: `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`,
		}, {
			tmap: `1-2
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`,
		},
	}

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			gmap := gamemap.NewMapSquare(2)
			maximizerSambol = "A"
			minimizerSambol = "B"
			if test.tmap[0] == '2' {
				maximizerSambol, minimizerSambol = minimizerSambol, maximizerSambol
			}
			gmap.Update(test.tmap, maximizerSambol)
			node := NewNode(nil, false, gmap)

			assert.Equal(t, 0, len(node.GetChildren()))

			child := node.Expand()

			assert.Equal(t, 7, len(node.GetChildren()))
			assert.Contains(t, node.remainingActions, child.causingAction)
			assert.NotEqual(t, node.turn, child.turn)

			for _, c := range node.children {
				assert.Equal(t, string(c.gmap.GetEdgeState(int(c.causingAction[0]), int(c.causingAction[1]))), maximizerSambol)
				assert.NotEqual(t, c.turn, node.turn)

			}
		})
	}

}
func TestExpandMapMoreLayers(t *testing.T) {
	tmap := `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`

	gmap := gamemap.NewMapSquare(2)
	maximizerSambol = "A"
	minimizerSambol = "B"
	if tmap[0] == '2' {
		maximizerSambol, minimizerSambol = minimizerSambol, maximizerSambol
	}
	gmap.Update(tmap, maximizerSambol)
	node := NewNode(nil, false, gmap)

	assert.Equal(t, 0, len(node.GetChildren()))

	// t.Log(node.gmap)

	node.Expand()
	child := node.children[0]
	assert.Equal(t, 7, len(node.GetChildren()))
	assert.Contains(t, node.remainingActions, child.causingAction)
	assert.Equal(t, string(child.gmap.GetEdgeState(int(child.causingAction[0]), int(child.causingAction[1]))), maximizerSambol)
	assert.NotEqual(t, node.turn, child.turn)

	// t.Log(child.gmap)

	child.Expand()
	newchild := child.children[0]
	assert.Equal(t, 6, len(child.GetChildren()))
	assert.Contains(t, child.remainingActions, newchild.causingAction)
	assert.Equal(t, string(newchild.gmap.GetEdgeState(int(newchild.causingAction[0]), int(newchild.causingAction[1]))), minimizerSambol)
	assert.Equal(t, string(newchild.gmap.Cells[0][0].OwnedBy), minimizerSambol)
	assert.NotEqual(t, node.turn, newchild.turn)
	child = newchild

	// t.Log(child.gmap)

	child.Expand()
	newchild = child.children[0]
	assert.Equal(t, 5, len(child.GetChildren()))
	assert.Contains(t, child.remainingActions, newchild.causingAction)
	assert.Equal(t, string(newchild.gmap.GetEdgeState(int(newchild.causingAction[0]), int(newchild.causingAction[1]))), minimizerSambol)
	assert.Equal(t, node.turn, newchild.turn)
	child = newchild
	// t.Log(child.gmap)

	child.Expand()
	newchild = child.children[0]
	assert.Equal(t, 4, len(child.GetChildren()))
	assert.Contains(t, child.remainingActions, newchild.causingAction)
	assert.Equal(t, string(newchild.gmap.GetEdgeState(int(newchild.causingAction[0]), int(newchild.causingAction[1]))), maximizerSambol)
	assert.Equal(t, node.turn, newchild.turn)
	child = newchild
	// t.Log(child.gmap)

}
