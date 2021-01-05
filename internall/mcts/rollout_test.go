package mcts

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/stretchr/testify/assert"
)

func TestRollout(t *testing.T) {

	testmap := []struct {
		action Action
		result []float64
		eval   []float64
	}{
		{
			action: Action{0, 1},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{0, 3},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{3, 2},
			result: []float64{0},
			eval:   []float64{0},
		},
	}

	tmap := `1-2
0-0
@A@A@
-#A#-
@B@-@
-#A#-
@B@B@`

	gmap := gamemap.NewMapSquare(2)
	maximizerSambol = "A"
	minimizerSambol = "B"
	if tmap[0] == '2' {
		maximizerSambol, minimizerSambol = minimizerSambol, maximizerSambol
	}
	gmap.Update(tmap, maximizerSambol)
	node := NewNode(nil, false, gmap)

	// assert.Equal(t, 2, node.Eval())

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			node.Expand()
			for _, v := range node.GetChildren() {
				if v.causingAction[0] == test.action[0] && v.causingAction[1] == test.action[1] {
					node = v
					break
				}
			}

			assert.Contains(t, test.result, node.RollOut(5))
			assert.Contains(t, test.eval, node.Eval())

		})
	}
}
func TestRolloutReverse(t *testing.T) {

	testmap := []struct {
		action Action
		result []float64
		eval   []float64
	}{
		{
			action: Action{0, 1},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{0, 3},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{3, 2},
			result: []float64{0},
			eval:   []float64{0},
		},
	}

	tmap := `2-1
0-0
@A@A@
-#A#-
@B@-@
-#A#-
@B@B@`

	gmap := gamemap.NewMapSquare(2)
	maximizerSambol = "A"
	minimizerSambol = "B"
	if tmap[0] == '2' {
		maximizerSambol, minimizerSambol = minimizerSambol, maximizerSambol
	}
	gmap.Update(tmap, maximizerSambol)
	node := NewNode(nil, false, gmap)

	// assert.Equal(t, 2, node.Eval())

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			node.Expand()
			for _, v := range node.GetChildren() {
				if v.causingAction[0] == test.action[0] && v.causingAction[1] == test.action[1] {
					node = v
					break
				}
			}

			assert.Contains(t, test.result, node.RollOut(5))
			assert.Contains(t, test.eval, node.Eval())

		})
	}
}

func TestRolloutOpenMap(t *testing.T) {

	testmap := []struct {
		action Action
		result []float64
		eval   []float64
	}{
		{
			action: Action{0, 1},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{0, 3},
			result: []float64{0, 6},
			eval:   []float64{2},
		},
		{
			action: Action{3, 2},
			result: []float64{0},
			eval:   []float64{0},
		},
	}

	tmap := `2-1
0-0
@-@A@
-#A#-
@B@-@
-#A#-
@B@B@`

	gmap := gamemap.NewMapSquare(2)
	maximizerSambol = "A"
	minimizerSambol = "B"
	if tmap[0] == '2' {
		maximizerSambol, minimizerSambol = minimizerSambol, maximizerSambol
	}
	gmap.Update(tmap, maximizerSambol)
	node := NewNode(nil, false, gmap)

	// assert.Equal(t, 2, node.Eval())

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			node.Expand()
			for _, v := range node.GetChildren() {
				if v.causingAction[0] == test.action[0] && v.causingAction[1] == test.action[1] {
					node = v
					break
				}
			}

			assert.Contains(t, test.result, node.RollOut(5))
			assert.Contains(t, test.eval, node.Eval())

		})
	}
}
