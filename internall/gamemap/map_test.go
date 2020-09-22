package gamemap_test

import (
	"testing"
	"fmt"
	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

func TestCreateMap(t *testing.T) {
	testmap := `2-1
	0-0
	@-@A@-@-@
	-#-#-#-#-
	@-@-@-@-@
	-#-#-#-#-
	@-@-@-@-@
	-#-#-#-#-
	@-@-@-@-@
	-#-#B#-#-
	@-@-@-@-@`

	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap)

	fmt.Println(gmap)
	// assert.Equal(t, gmap.Cells[0][2].UpperEdge, gamemap.IsAEdge)
	// assert.Equal(t, gmap.Cells[3][2].RightEdge, gamemap.IsBEdge)
	// assert.Equal(t, gmap.Cells[3][3].LeftEdge, gamemap.IsBEdge)

}
