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

func TestUpdateMap(t *testing.T) {
	testmap := []string{`2-1
0-0
@-@A@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@B@`,
		`2-1
0-0
@-@A@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#A#-#-
@-@B@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@B@-@`}

	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap[0])
	// fmt.Println(gmap)

	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[0][1].UpperEdge.State)
	assert.Equal(t, gamemap.IsBEdge, gmap.Cells[3][3].LowerEdge.State)

	gmap.Update(testmap[1])
	// fmt.Println(gmap)
	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[1][1].RightEdge.State)
	assert.Equal(t, gamemap.IsAEdge, gmap.Cells[1][2].LeftEdge.State)
	assert.Equal(t, gamemap.IsBEdge, gmap.Cells[1][1].LowerEdge.State)
	assert.Equal(t, gamemap.IsBEdge, gmap.Cells[2][1].UpperEdge.State)

}

func BenchmarkUpdateMap(b *testing.B) {
	b.StopTimer()

	testmap := []string{`2-1
0-0
@-@A@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@-@B@`,
		`2-1
0-0
@-@A@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#A#-#-
@-@B@-@-@
-#-#-#-#-
@-@-@-@-@
-#-#-#-#-
@-@-@B@-@`}

	gmap := gamemap.NewMapSquare(4)
	gmap.Update(testmap[0])
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gmap.Update(testmap[1])
	}
}
