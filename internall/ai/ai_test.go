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
