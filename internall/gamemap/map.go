package gamemap

import (
	"strconv"
	"strings"
)

//Map Keeps the track of Cells in game
type Map struct {
	Cells [][]*Cell
}

//NewMapRect Creates ne map with diffrent hight and width
func NewMapRect(hight, width int8) *Map {
	gameMap := new(Map)
	gameMap.Cells = make([][]*Cell, hight)
	for i := int8(0); i < hight; i++ {
		gameMap.Cells[i] = make([]*Cell, 0)
		for j := int8(0); j < width; j++ {
			gameMap.Cells[i] = append(gameMap.Cells[i], NewCellXY(2*i+1, 2*j+1))
		}
	}
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
			res += "*\t" + string(gameMap.Cells[i][j].UpperEdge.State) + "\t*"
		}
		res += "\n"
		for j := 0; j < len(gameMap.Cells[i]); j++ {
			res += string(gameMap.Cells[i][j].LeftEdge.State) + "\t" + "(" + strconv.Itoa(int(gameMap.Cells[i][j].Coordinate.X)) + "," + strconv.Itoa(int(gameMap.Cells[i][j].Coordinate.Y)) + ")" + "\t" + string(gameMap.Cells[i][j].RightEdge.State)
		}
		res += "\n"
		for j := 0; j < len(gameMap.Cells[i]); j++ {
			res += "*\t" + string(gameMap.Cells[i][j].LowerEdge.State) + "\t*"
		}
	}
	res += "\n"
	return res
}

//SetEdgeState sets the edge if is full or empity
func (gameMap Map) setEdgeState(X, Y int, edgeState EdgeState) {
	//its up and down
	if X%2 == 1 {
		//not the upest raw
		if Y > 0 {
			gameMap.Cells[(Y-2)/2][(X-1)/2].LowerEdge.State = edgeState
		}
		//not the lowest raw
		if Y < len(gameMap.Cells)*2 {
			gameMap.Cells[(Y)/2][(X-1)/2].UpperEdge.State = edgeState
		}
	} else { //its left or right
		//not the most left column
		if X > 0 {
			gameMap.Cells[Y/2][(X-1)/2].RightEdge.State = edgeState
		}
		//not the most right column
		if X < len(gameMap.Cells)*2 {
			gameMap.Cells[(Y-1)/2][(X)/2].LeftEdge.State = edgeState
		}
	}
}

//Update updates the game map according to the raw text it gets
func (gameMap Map) Update(rawMap string) {
	rawMap = rawMap[strings.Index(rawMap, "@"):]
	rawMap = strings.ReplaceAll(rawMap, "\n", "")
	Aindexes := findIndex(rawMap, 'A')
	Bindexes := findIndex(rawMap, 'B')

	for _, ind := range Aindexes {
		gameMap.setEdgeState(ind%9, ind/9, IsAEdge)
	}
	for _, ind := range Bindexes {
		gameMap.setEdgeState(ind%9, ind/9, IsBEdge)
	}

	// fmt.Println(Aindexes)
	// fmt.Println(Bindexes)
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
