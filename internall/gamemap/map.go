package gamemap

import (
	"strconv"
	"strings"
)

//Map Keeps the track of Cells in game
type Map struct {
	Cells    [][]*Cell
	AIndexes map[int]interface{}
	BIndexes map[int]interface{}
}

//NewMapRect Creates ne map with diffrent hight and width
func NewMapRect(hight, width int8) *Map {
	gameMap := new(Map)
	gameMap.Cells = make([][]*Cell, hight)
	for i := int8(0); i < hight; i++ {
		gameMap.Cells[i] = make([]*Cell, 0)
		for j := int8(0); j < width; j++ {
			gameMap.Cells[i] = append(gameMap.Cells[i], NewCellXY(2*j+1, 2*i+1))
		}
	}
	gameMap.AIndexes = make(map[int]interface{})
	gameMap.BIndexes = make(map[int]interface{})
	return gameMap
}

//NewMapSquare Creates ne map with same hight and width
func NewMapSquare(length int8) *Map {
	return NewMapRect(length, length)
}

func (gameMap Map) String() string {
	res := ""
	for i := 0; i < len(gameMap.Cells); i++ {
		res += "\n"
		for j := 0; j < len(gameMap.Cells[i]); j++ {
			res += "*\t\t" + string(gameMap.Cells[i][j].UpperEdge.State) + "\t\t*"
		}
		res += "\n"
		for j := 0; j < len(gameMap.Cells[i]); j++ {
			res += string(gameMap.Cells[i][j].LeftEdge.State) + "\t" + "(" + strconv.Itoa(int(gameMap.Cells[i][j].Coordinate.X)) + ",|" + strconv.Itoa(int(gameMap.Cells[i][j].FilledEdgeCount)) /* string(gameMap.Cells[i][j].OwnedBy) */ + "|," + strconv.Itoa(int(gameMap.Cells[i][j].Coordinate.Y)) + ")" + "\t" + string(gameMap.Cells[i][j].RightEdge.State)
		}
		res += "\n"
		for j := 0; j < len(gameMap.Cells[i]); j++ {
			res += "*\t\t" + string(gameMap.Cells[i][j].LowerEdge.State) + "\t\t*"
		}
	}
	res += "\n"
	return res
}

//SetEdgeState sets the edge if is full or empity
func (gameMap Map) SetEdgeState(X, Y int, edgeState EdgeState) bool {
	res := false
	//its up and down
	if X%2 == 1 {
		//not the upest raw
		if Y > 0 && gameMap.Cells[(Y-2)/2][(X-1)/2].LowerEdge.State == IsFreeEdge {
			gameMap.Cells[(Y-2)/2][(X-1)/2].LowerEdge.State = edgeState
			// gameMap.Cells[(Y-2)/2][(X-1)/2].Edges[3] = gameMap.Cells[(Y-2)/2][(X-1)/2].LowerEdge
			gameMap.Cells[(Y-2)/2][(X-1)/2].FilledEdgeCount++
			if gameMap.Cells[(Y-2)/2][(X-1)/2].FilledEdgeCount == 4 {
				gameMap.Cells[(Y-2)/2][(X-1)/2].OwnedBy = edgeState
				res = true
			}
		}
		//not the lowest raw
		if Y < len(gameMap.Cells)*2 && gameMap.Cells[(Y)/2][(X-1)/2].UpperEdge.State == IsFreeEdge {
			gameMap.Cells[(Y)/2][(X-1)/2].UpperEdge.State = edgeState
			// gameMap.Cells[(Y-2)/2][(X-1)/2].Edges[0] = gameMap.Cells[(Y-2)/2][(X-1)/2].UpperEdge
			gameMap.Cells[(Y)/2][(X-1)/2].FilledEdgeCount++
			if gameMap.Cells[(Y)/2][(X-1)/2].FilledEdgeCount == 4 {
				gameMap.Cells[(Y)/2][(X-1)/2].OwnedBy = edgeState
				res = true
			}
		}
	} else { //its left or right
		//not the most left column
		if X > 0 && gameMap.Cells[Y/2][(X-1)/2].RightEdge.State == IsFreeEdge {
			gameMap.Cells[Y/2][(X-1)/2].RightEdge.State = edgeState
			// gameMap.Cells[(Y-2)/2][(X-1)/2].Edges[2] = gameMap.Cells[(Y-2)/2][(X-1)/2].RightEdge
			gameMap.Cells[Y/2][(X-1)/2].FilledEdgeCount++
			if gameMap.Cells[Y/2][(X-1)/2].FilledEdgeCount == 4 {
				gameMap.Cells[Y/2][(X-1)/2].OwnedBy = edgeState
				res = true
			}
		}
		//not the most right column
		if X < len(gameMap.Cells)*2 {
			if gameMap.Cells[(Y-1)/2][(X)/2].LeftEdge.State == IsFreeEdge {
				gameMap.Cells[(Y-1)/2][(X)/2].LeftEdge.State = edgeState
				// gameMap.Cells[(Y-2)/2][(X-1)/2].Edges[1] = &gameMap.Cells[(Y-2)/2][(X-1)/2].RightEdge
				gameMap.Cells[(Y-1)/2][(X)/2].FilledEdgeCount++
				if gameMap.Cells[(Y-1)/2][(X)/2].FilledEdgeCount == 4 {
					gameMap.Cells[(Y-1)/2][(X)/2].OwnedBy = edgeState
					res = true
				}
			}
		}
	}
	return res
}

//Update updates the game map according to the raw text it gets
func (gameMap Map) Update(rawMap, maximizerSambol string) {
	minimizerSambol := "B"
	if maximizerSambol == "B" {
		minimizerSambol = "A"
	}
	scoreSection := rawMap[strings.Index(rawMap, "\n"):strings.Index(rawMap, "@")]
	scores := strings.Split(scoreSection[1:len(scoreSection)-1], "-")
	minimizerScore, _ := strconv.Atoi(scores[1])
	maximizerScore, _ := strconv.Atoi(scores[0])

	rawMap = rawMap[strings.Index(rawMap, "@"):]
	rawMap = strings.ReplaceAll(rawMap, "\n", "")
	Aindexes := findIndex(rawMap, 'A')
	Bindexes := findIndex(rawMap, 'B')
	Aindexes = difference(Aindexes, gameMap.AIndexes)
	Bindexes = difference(Bindexes, gameMap.BIndexes)
	Xlength := len(gameMap.Cells[0])*2 + 1
	Ylength := len(gameMap.Cells)*2 + 1

	for _, ind := range Aindexes {
		gameMap.SetEdgeState(ind%Xlength, ind/Ylength, IsAEdge)
	}
	for _, ind := range Bindexes {
		gameMap.SetEdgeState(ind%Xlength, ind/Ylength, IsBEdge)
	}

	for _, raw := range gameMap.Cells {
		for _, cell := range raw {
			switch cell.FilledEdgeCount {
			case 4:
				if maximizerScore > 0 {
					cell.OwnedBy = EdgeState(maximizerSambol)
					maximizerScore--
				} else if minimizerScore > 0 {
					cell.OwnedBy = EdgeState(minimizerSambol)
					minimizerScore--
				}
			}
		}
	}

	gameMap.AIndexes = appendIndexes(Aindexes, gameMap.AIndexes)
	gameMap.BIndexes = appendIndexes(Bindexes, gameMap.BIndexes)
	// gameMap.bIndexes = append(*gameMap.bIndexes, Bindexes...)

	// fmt.Println(Aindexes)
	// fmt.Println(Bindexes)
}

// Clone retruns new identical map
func (gameMap Map) Clone() Map {
	gmap := NewMapRect(int8(len(gameMap.Cells)), int8(len(gameMap.Cells[0])))
	for i, raw := range gameMap.Cells {
		gmap.Cells[i] = make([]*Cell, len(raw))
		for j, cell := range raw {
			gmap.Cells[i][j] = NewCell(cell.Coordinate)
			gmap.Cells[i][j].FilledEdgeCount = cell.FilledEdgeCount
			gmap.Cells[i][j].LeftEdge = cell.LeftEdge
			gmap.Cells[i][j].RightEdge = cell.RightEdge
			gmap.Cells[i][j].UpperEdge = cell.UpperEdge
			gmap.Cells[i][j].LowerEdge = cell.LowerEdge
			gmap.Cells[i][j].OwnedBy = cell.OwnedBy
			gmap.Cells[i][j].Edges = [4]*Edge{&gmap.Cells[i][j].UpperEdge, &gmap.Cells[i][j].LeftEdge, &gmap.Cells[i][j].RightEdge, &gmap.Cells[i][j].LowerEdge}
		}
		// for j, cell := range raw {
		// 	fmt.Println(&gmap.Cells[i][j], &cell, cell, &gameMap.Cells[i][j])
		// 	gmap.Cells[i][j] = gameMap.Cells[i][j]
		// 	fmt.Println(&gmap.Cells[i][j], &cell, cell, &gameMap.Cells[i][j])
		// }

	}
	// fmt.Println("\n------------------------------------------------")
	// fmt.Println(gmap)
	// fmt.Println("cloned from")
	// fmt.Println(gameMap)
	// fmt.Println("------------------------------------------------")
	return *gmap
}

// HasFreeEdge check if map has free edge
func (gameMap Map) HasFreeEdge() bool {
	for _, raw := range gameMap.Cells {
		for _, cell := range raw {
			if cell.FilledEdgeCount != 4 {
				return true
			}
		}
	}
	return false
}

func findIndex(rawText string, char rune) []int {
	indexes := make([]int, 0)
	var i int = strings.IndexRune(rawText, char)
	j := i
	for i > -1 {
		indexes = append(indexes, j)
		j++
		i = strings.IndexRune(rawText[j:], char)
		j += i
	}

	return indexes
}

func difference(a []int, mb map[int]interface{}) []int {
	var diff []int
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func appendIndexes(a []int, mb map[int]interface{}) map[int]interface{} {
	for _, v := range a {
		mb[v] = nil
	}
	return mb
}
