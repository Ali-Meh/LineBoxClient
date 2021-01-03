package mcts_test

import (
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateMaximizer(t *testing.T) {
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
	assert.Equal(t, 40, 40)
}
