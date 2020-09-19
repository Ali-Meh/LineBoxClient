package gamemap

import "fmt"

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
			res += fmt.Sprint(gameMap.Cells[i][j])
		}
	}
	return res
}
