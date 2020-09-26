package gamemap

import "strconv"

//EdgeState state of the Edge
type EdgeState string

const (
	IsAEdge    EdgeState = "A"
	IsBEdge    EdgeState = "B"
	IsFreeEdge EdgeState = "-"
)

//Coordinates has the coordinate of map objects
type Coordinates struct {
	X int8
	Y int8
}

//Edge keeps the state of the Edge
type Edge struct {
	Coordinates
	State EdgeState
}

//Cell is the point of the map can be owned and has point
type Cell struct {
	UpperEdge       Edge
	LowerEdge       Edge
	LeftEdge        Edge
	RightEdge       Edge
	filledEdgeCount int
	Coordinate      Coordinates
}

//NewCell Creates a cell with its edges
func NewCell(cellCoord Coordinates) *Cell {
	cell := new(Cell)
	cell.Coordinate = cellCoord
	cell.filledEdgeCount = 0
	cell.LeftEdge = Edge{Coordinates: Coordinates{X: cellCoord.X - 1, Y: cellCoord.Y}, State: IsFreeEdge}
	cell.RightEdge = Edge{Coordinates: Coordinates{X: cellCoord.X + 1, Y: cellCoord.Y}, State: IsFreeEdge}
	cell.UpperEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y - 1}, State: IsFreeEdge}
	cell.LowerEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y + 1}, State: IsFreeEdge}
	return cell
}

//NewCellXY Creates a cell with its edges based on the corrdinates given to it
func NewCellXY(X, Y int8) *Cell {
	cellCoord := Coordinates{X: X, Y: Y}
	// cell := new(Cell)
	// cell.Coordinate = cellCoord
	// cell.LeftEdge = Edge{Coordinates: Coordinates{X: cellCoord.X - 1, Y: cellCoord.Y}, State: IsFreeEdge}
	// cell.RightEdge = Edge{Coordinates: Coordinates{X: cellCoord.X + 1, Y: cellCoord.Y}, State: IsFreeEdge}
	// cell.UpperEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y - 1}, State: IsFreeEdge}
	// cell.LowerEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y + 1}, State: IsFreeEdge}
	return NewCell(cellCoord)
}

func (c Cell) String() string {
	// res := "("
	// res += strconv.Itoa(int(c.Coordinate.X)) + "-" + strconv.Itoa(int(c.Coordinate.Y))
	// res += ")\t"
	//-----------------------------------------------
	// res := "\t" + "U" + string(c.UpperEdge.State) + "\t"
	// res += "L" + string(c.LeftEdge.State) + "(" + strconv.Itoa(int(c.Coordinate.X)) + "," + strconv.Itoa(int(c.Coordinate.Y)) + ")" + string(c.RightEdge.State) + "R" + ""
	// res += "\t" + "D" + string(c.LowerEdge.State) + "\t"
	//----------------------------------------------
	res := "*\t" + string(c.UpperEdge.State) + "\t*\n"
	res += string(c.LeftEdge.State) + "\t" + "(" + strconv.Itoa(int(c.Coordinate.X)) + "," + strconv.Itoa(int(c.filledEdgeCount)) + "," + strconv.Itoa(int(c.Coordinate.Y)) + ")" + "\t" + string(c.RightEdge.State) + "\n"
	res += "*\t" + string(c.LowerEdge.State) + "\t*\n\n"
	return res
}
