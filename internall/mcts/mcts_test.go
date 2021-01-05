package mcts_test

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/ali-meh/LineBoxClient/internall/mcts"
	"github.com/stretchr/testify/assert"
)

func TestSelectMove(t *testing.T) {
	testmap := []struct {
		tmap   string
		result [][]int8
	}{
/* 	{
		tmap: `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`,
		result: [][]int8{{4, 1},},
	},
	{
		tmap: `2-1
0-0
@A@A@
-#B#-
@B@-@
-#A#A
@-@B@`,
		result: [][]int8{{0, 1},},
	}, {
		tmap: `2-1
0-0
@A@-@
-#-#-
@B@-@
-#A#-
@-@B@`,
		result: [][]int8{{3, 0},},//BUG
	},
	{
		tmap: `2-1
0-0
@A@A@
B#-#-
@B@-@
-#A#-
@-@B@`,
		result: [][]int8{{2, 1},},
	},*/
	{
		tmap: `1-2
0-1
@A@A@
B#A#-
@B@-@
-#A#-
@-@B@`,
		result: [][]int8{{1,4},{0,3}},//BUG
	},
// 	{
// 		tmap: `2-1
// 0-1
// @A@A@
// B#A#-
// @B@-@
// A#A#-
// @-@B@`,
// 		result: [][]int8{{1, 4},},
// 	},
// 	{
// 		tmap: `2-1
// 0-0
// @A@A@
// B#A#-
// @B@-@
// A#A#-
// @B@B@`,
// 		result: [][]int8{{4, 3},{4,1}},
// 	}, 
// 	{
// 		tmap: `2-1
// 0-0
// @A@A@
// -#-#B
// @B@-@
// -#A#-
// @-@B@`,
// 		result: [][]int8{{0, 3},{1,4}},//BUG
// 	},
// 		{
// 			tmap: `2-1
// 0-0
// @A@B@
// -#-#B
// @A@-@
// A#-#B
// @-@A@`,
// 	result: [][]int8{{0, 1}, {2, 1}, {3, 2}, {2, 3}, {1, 4}}, //it should find value failing!!
// },
// 		{
// 			tmap: `2-1
// 0-0
// @A@B@
// A#-#B
// @A@-@
// A#-#B
// @-@A@`,
// 			result: [][]int8{{2, 1}},
// 		},
// 		{
// 			tmap: `2-1
// 0-0
// @A@B@
// A#-#B
// @A@-@
// A#-#B
// @B@A@`,
// 			result: [][]int8{{2, 1}, {2, 3}},
// 		},
	}

	for i, test := range testmap {
		t.Run("test map select #"+fmt.Sprintf("%d", i), func(t *testing.T) {
			//create map
			gmap := gamemap.NewMapSquare(2)
			maximizer := "A"
			if test.tmap[0] == '2' {
				maximizer = "B"
			}
			gmap.Update(test.tmap, maximizer)

			for i := 0; i < 5; i++ {
				move := mcts.SelectMove(*gmap, maximizer)
				// move := mcts.SelectMoveRecursive(*gmap, maximizer)
				//assert the evaluation
				// fmt.Println(move)
				assert.Contains(t, test.result, move)
			}
		})
	}
}

func TestSelect1(t *testing.T) {
	testmap := `2-1
0-0
@A@-@B@-@
-#-#-#-#A
@-@-@-@A@
A#-#B#-#-
@A@A@-@-@
A#-#B#-#B
@-@-@-@-@
B#A#-#B#-
@-@B@B@A@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	move := mcts.SelectMove(*gmap, minimizerSambol)
	assert.Equal(t, []int8{0, 1}, move)
}
func TestSelect2(t *testing.T) {
	testmap := `2-1
0-0
@B@-@-@-@
-#-#-#-#A
@A@-@A@B@
-#-#-#-#B
@-@-@A@-@
-#-#-#-#-
@B@-@-@-@
-#-#A#-#-
@-@-@-@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	maximizerSambol := "A"
	if testmap[0] == '2' {
		maximizerSambol = "B"
	}
	gmap.Update(testmap, maximizerSambol)

	move := mcts.SelectMove(*gmap, maximizerSambol)
	assert.Equal(t, []int8{2, 1}, move)
}

func TestSelect3(t *testing.T) {
	testmap := `2-1
0-0
@-@-@-@-@
-#-#-#-#-
@B@-@-@-@
A#-#A#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@A@-@
-#B#-#-#-
@-@-@-@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)

	move := mcts.SelectMove(*gmap, minimizerSambol)
	assert.Equal(t, []int8{1, 0}, move)
}
