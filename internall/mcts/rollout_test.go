package mcts

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/stretchr/testify/assert"
)

func TestPrioritisedChoises(t *testing.T) {
	testmap := []struct {
		tmap   string
		result map[string]int
	}{
		{
			tmap: `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`,
			result: map[string]int{
				"[0 1]": 0,
				"[2 1]": 0,
				"[4 1]": 2,
				"[2 2]": 0,
				"[3 2]": 0,
				"[0 3]": 0,
				"[1 4]": 0,
				"[4 3]": 0,
			},
					}, {
						tmap: `1-2
0-0
@A@A@
A#B#B
@B@-@
-#-#-
@A@-@`,
						result: map[string]int{
							"[3 2]": 3,
							"[0 3]": 0,
							"[2 3]": 0,
							"[4 3]": 1,
							"[3 4]": 1,
						},
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

			res := prioritiseActions(*gmap, extractRemainingMoves(gmap))

			for k, v := range res {
				for _, a := range v {
					t.Log(a, "==>", k)
					assert.Equal(t, test.result[fmt.Sprint(a)], k)
				}
			}

		})
	}

}
func TestRollout(t *testing.T) {

	testmap := []struct {
		action Action
		result []float64
		eval   []float64
	}{
		{
			action: Action{0, 1},
			result: []float64{6,2},
			eval:   []float64{2},
		},
		{
			action: Action{0, 3},
			result: []float64{ 6, 2},
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
			result: []float64{2, 6},
			eval:   []float64{2},
		},
		{
			action: Action{0, 3},
			result: []float64{2, 6},
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
			action: Action{1, 0},
			result: []float64{2, -4, 0, 4, -2},
			eval:   []float64{0},
		},
		{
			action: Action{1, 4},
			result: []float64{-4, 0, 4},
			eval:   []float64{0},
		},
		{
			action: Action{4, 3},
			result: []float64{-5,-1,1},
			eval:   []float64{-1},
		},
		{
			action: Action{3, 2},
			result: []float64{-2,-6},
			eval:   []float64{-2},
		},
		{
			action: Action{4, 1},
			result: []float64{-2,-6},
			eval:   []float64{-2},
		},
		{
			action: Action{1, 2},
			result: []float64{0},
			eval:   []float64{0},
		},
		// {
		// 	action: Action{3, 2},
		// 	result: []float64{0},
		// 	eval:   []float64{0},
		// },
	}

	tmap := `1-2
0-0
@-@A@
-#B#-
@-@-@
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
