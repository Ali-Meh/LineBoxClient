package mcts_test

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestRollout(t *testing.T) {
	testmap := []struct {
		tmap   string
		result []float64
	}{

		{
			tmap: `2-1
0-2
@A@A@
B#A#-
@B@-@
A#A#-
@B@B@`,
			result: []float64{0, 4},
		},
		{
			tmap: `2-1
0-2
@A@A@
B#A#A
@B@-@
A#A#-
@B@B@`,
			result: []float64{0, 4},
		},
		{
			tmap: `2-1
0-2
@A@A@
B#A#A
@B@-@
A#A#-
@B@B@`,
			result: []float64{0, 4},
		},
		{
			tmap: `2-1
3-0
@A@A@
B#A#A
@B@B@
A#A#-
@B@B@`,
			result: []float64{-2},
		},
		{
			tmap: `2-1
4-0
@A@A@
B#A#A
@B@B@
A#A#A
@B@B@`,
			result: []float64{-4},
		},
	}

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			//create map
			gmap := gamemap.NewMapSquare(2)
			minimizerSambol := "A"
			if test.tmap[0] == '2' {
				minimizerSambol = "B"
			}
			gmap.Update(test.tmap, minimizerSambol)

			rootNode := mcts.NewNode([]int8{}, true, gmap)

			res := rootNode.RollOut(6)
			assert.Contains(t, test.result, res)
		})
	}
}
