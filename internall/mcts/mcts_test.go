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
		{
			tmap: `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#-
@-@B@`,
			result: [][]int8{{4,1}},
		},
		{
			tmap: `2-1
0-0
@A@A@
-#B#-
@B@-@
-#A#A
@-@B@`,
			result: [][]int8{{0, 1}, {3, 2}},
		},
		 {
			tmap: `2-1
0-0
@A@-@
-#-#-
@B@-@
-#A#-
@-@B@`,
			result: [][]int8{{3, 0}, {4, 1}},
		},
		{
			tmap: `2-1
0-0
@A@A@
B#-#-
@B@-@
-#A#-
@-@B@`,
			result: [][]int8{{2, 1}},
		},
		{
			tmap: `1-2
1-0
@A@A@
B#A#-
@B@-@
-#A#-
@-@B@`,
			result: [][]int8{{1, 4}, {0, 3}}, //BUG
		},
		{
			tmap: `2-1
0-1
@A@A@
B#A#-
@B@-@
A#A#-
@-@B@`,
			result: [][]int8{{1, 4}},
		},
		{
			tmap: `2-1
1-1
@A@A@
B#A#-
@B@-@
A#A#-
@B@B@`,
			result: [][]int8{{4, 3}, {4, 1}},
		},
		{
			tmap: `2-1
0-0
@A@A@
-#-#B
@B@-@
-#A#-
@-@B@`,
			result: [][]int8{{0, 3}, {1, 4}}, //BUG
		},
		{
			tmap: `2-1
0-0
@A@B@
-#-#B
@A@-@
A#-#B
@-@A@`,
			result: [][]int8{{0, 1}, {2, 1}, {3, 2}, {2, 3}, {1, 4}}, //it should find value failing!!
		},
		{
			tmap: `2-1
0-0
@A@B@
A#-#B
@A@-@
A#-#B
@-@A@`,
			result: [][]int8{{2, 1}},
		},
		{
			tmap: `2-1
0-0
@A@B@
A#-#B
@A@-@
A#-#B
@B@A@`,
			result: [][]int8{{2, 1}, {2, 3}},
		},
		{
			tmap: `2-1
0-0
@A@B@
-#-#A
@A@-@
-#B#B
@-@-@`,
			result: [][]int8{{1, 4}, {0, 3}},
		},
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

			for i := 0; i < 1; i++ {
				move := mcts.SelectMove(*gmap, maximizer)
				// move := mcts.SelectMoveRecursive(*gmap, maximizer)
				//assert the evaluation
				fmt.Println(move)
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
	assert.NotContains(t,[][]int8{{1, 2},{2, 3},{3, 1},{7, 0},{7, 6},{8, 7},{5, 6},{4, 7},{1, 8},{1, 6},{2, 5},{3, 6},},move)
	// assert.NotEqual(t, []int8{6, 5}, move)
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
	assert.NotEqual(t, []int8{0, 1}, move)
	assert.NotEqual(t, []int8{2, 1}, move)
	assert.NotEqual(t, []int8{6, 1}, move)
	assert.NotEqual(t, []int8{7, 0}, move)
	assert.NotEqual(t, []int8{6, 3}, move)
	assert.NotEqual(t, []int8{7, 4}, move)
	assert.NotEqual(t, []int8{4, 3}, move)
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
	assert.NotEqual(t, []int8{2, 3}, move)
	assert.NotEqual(t, []int8{1, 4}, move)
}
func TestSelect4(t *testing.T) {
	testmap := `2-1
0-0
@A@-@-@-@
-#-#A#-#-
@A@-@-@-@
-#-#-#-#B
@A@-@B@-@
-#B#-#-#-
@-@-@-@A@
B#-#A#-#B
@-@-@-@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	move := mcts.SelectMove(*gmap, minimizerSambol)
	t.Log(move)
	assert.NotEqual(t, []int8{0, 1}, move)
	assert.NotEqual(t, []int8{2, 1}, move)
	assert.NotEqual(t, []int8{0, 3}, move)
	assert.NotEqual(t, []int8{1, 6}, move)
	assert.NotEqual(t, []int8{0, 5}, move)
	assert.NotEqual(t, []int8{7, 8}, move)
	assert.NotEqual(t, []int8{6, 7}, move)
	assert.NotEqual(t, []int8{0, 5}, move)
}

func TestSelect5(t *testing.T) {
	testmap := `2-1
0-0
@B@B@B@B@
-#-#A#-#B
@A@-@-@-@
-#-#B#A#-
@A@A@-@A@
-#-#-#-#-
@-@-@-@-@
A#-#-#-#-
@-@-@-@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	move := mcts.SelectMove(*gmap, minimizerSambol)
	t.Log(move)
	assert.NotEqual(t, []int8{8, 3}, move)
}
func TestSelect6(t *testing.T) {
	testmap := `2-1
0-0
@-@-@-@-@
-#-#-#-#-
@-@A@-@-@
-#-#-#-#-
@A@-@-@-@
-#-#-#A#-
@-@-@-@-@
-#-#-#A#-
@-@B@B@B`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	move := mcts.SelectMove(*gmap, minimizerSambol)
	t.Log(move)
	assert.NotEqual(t, []int8{8, 7}, move)
}
func TestSelect7(t *testing.T) {
	testmap := `1-2
0-0
@B@B@A@-@
-#A#-#B#-
@B@-@-@B@
A#-#B#-#B
@-@B@B@-@
-#A#-#A#A
@B@-@-@-@
B#-#A#-#A
@-@A@A@A@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	move := mcts.SelectMove(*gmap, minimizerSambol)
	t.Log(move)
	assert.Equal(t, []int8{0, 1}, move)
}

func TestSelect8(t *testing.T) {
	testmap := `2-1
2-0
@-@A@A@A@
-#B#-#A#-
@A@-@-@-@
A#-#A#A#A
@-@A@-@-@
B#B#A#B#B
@-@B@-@-@
B#A#B#-#B
@B@B@B@B@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	for i := 0; i < 3; i++ {
		move := mcts.SelectMove(*gmap, minimizerSambol)
		t.Log(move)
		// assert.Contains(t, [][]int8{{0, 1}, {1, 0}}, move)
		assert.Contains(t, [][]int8{{1,6}}, move)
	}
}
func TestSelect9s(t *testing.T) {
	testmap := `1-2
0-0
@-@-@-@A@
-#B#B#-#B
@B@-@B@-@
-#-#-#-#A
@B@-@B@B@
-#A#A#-#A
@-@-@-@-@
A#-#B#-#A
@A@B@A@A@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	for i := 0; i < 3; i++ {
		move := mcts.SelectMove(*gmap, minimizerSambol)
		t.Log(move)
		assert.Contains(t, [][]int8{{0, 1}, {1, 0}, {3, 0}, {0, 3}}, move)
	}
}


func TestSelect10(t *testing.T) {
	testmap := `1-2
3-2
@A@B@A@B@
B#A#A#-#A
@B@B@-@-@
A#B#B#-#A
@B@A@B@B@
B#-#-#A#B
@-@B@-@A@
A#-#-#B#-
@A@A@A@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	gmap.Update(testmap, minimizerSambol)
	fmt.Println(gmap)
	// for i := 0; i < 3; i++ {
		move := mcts.SelectMove(*gmap, minimizerSambol)
		t.Log(move)
		// assert.Contains(t, [][]int8{{0, 1}, {1, 0}}, move)
		assert.Contains(t, [][]int8{{8,7},{7,8}}, move)
	// }
}
