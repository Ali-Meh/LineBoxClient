package ai_test

import (
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
	score := ai.Evaluate(*gmap, true)
	//assert the evaluation
	assert.Equal(t, 20, score)
}
