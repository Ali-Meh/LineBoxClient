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

	maximizer := "A"
	if testmap[0] == '2' {
		maximizer = "B"
	}
	move := mcts.SelectMove(*gmap, maximizer)

	assert.Equal(t, []int8{0, 1}, move)
}

func TestSelect1(t *testing.T) {
	testmap := `2-1
0-0
@A@-@B@B@
-#A#-#-#A
@-@-@A@-@
B#-#-#-#-
@-@B@-@-@
-#-#-#-#-
@-@A@-@B@
B#-#-#-#-
@A@-@A@-@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	move := mcts.SelectMove(*gmap, minimizerSambol)
	assert.Equal(t, []int8{2, 1}, move)
}
func TestSelect2(t *testing.T) {
	testmap := `2-1
0-0
@A@A@-@-@
B#-#-#-#-
@A@-@-@-@
-#-#-#-#-
@-@A@-@-@
-#-#-#-#B
@-@-@-@-@
-#-#-#-#-
@-@-@-@B@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	move := mcts.SelectMove(*gmap, minimizerSambol)
	assert.Equal(t, []int8{2, 1}, move)
}

func TestSelect3(t *testing.T) {
	testmap := `2-1
2-4
@-@B@B@B@
B#-#A#A#B
@-@-@A@A@
B#B#B#A#B
@-@-@B@A@
B#B#A#A#B
@-@-@A@A@
A#-#A#B#A
@A@B@A@A@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)
	minimizerSambol := "A"
	if testmap[0] == '2' {
		minimizerSambol = "B"
	}
	move := mcts.SelectMove(*gmap, minimizerSambol)
	assert.Equal(t, []int8{1, 0}, move)
}
