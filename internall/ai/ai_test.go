package ai_test

import (
	"fmt"
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/ai"
	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
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
	score := ai.Evaluate(*gmap, true, "A")
	//assert the evaluation
	assert.Equal(t, 20, score)
}

func TestMinimax(t *testing.T) {
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
	fmt.Println(gmap)
	score := ai.MiniMax(*gmap, 7, true, -999999, 999999)
	//assert the evaluation
	assert.Equal(t, 20, score)
}

func TestSelectMove(t *testing.T) {
	testmap := []struct {
		tmap   string
		depth  int
		result []int8
	}{
		{
			tmap: `2-1
0-0
@A@A@
-#-#-
@B@-@
-#A#A
@-@B@`,
			depth:  7,
			result: []int8{3, 2},
		},
	}

	for _, test := range testmap {
		t.Run("test map select", func(t *testing.T) {
			//create map
			gmap := gamemap.NewMapSquare(2)
			gmap.Update(test.tmap)
			fmt.Println(gmap)
			move := ai.SelectMove(*gmap, test.depth)
			//assert the evaluation
			fmt.Println(move)
			assert.Equal(t, test.result, move)
		})
	}
}

func TestSelect2(t *testing.T) {
	testmap := `2-1
0-0
@-@A@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@A@-@-@
-#-#-#-#B
@-@-@-@-@
-#-#-#-#-
@-@-@-@B@`

	//create map
	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)
	fmt.Println(gmap)
	move := ai.SelectMove(*gmap, 3)
	//assert the evaluation
	fmt.Println(move)
	assert.Equal(t, []int8{3, 2}, move)
}
