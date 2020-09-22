package gamemap_test

import (
	"testing"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
	"github.com/stretchr/testify/assert"
)

func TestCreateMap(t *testing.T) {
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

	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)

	// fmt.Println(gmap)
	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[0][1].UpperEdge.State)
	assert.Equal(t, gamemap.IsBEdge, gmap.Cells[3][3].LowerEdge.State)
	assert.Equal(t, gamemap.IsBEdge, gmap.Cells[2][3].RightEdge.State)
	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[1][1].LowerEdge.State)
	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[2][1].UpperEdge.State)

}
