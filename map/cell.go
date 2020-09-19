package map

//EdgeState state of the Edge
type EdgeState string

const (
	isAEdge    EdgeState = "A"
	isBEdge    EdgeState = "B"
	isFreeEdge EdgeState = "-"
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
	UpperEdge Edge
	LowerEdge Edge
	LeftEdge  Edge
	RightEdge Edge

	Coordinate Coordinates
}

//NewCell Creates a cell with its edges
func NewCell(cellCoord Coordinates) *Cell {
	cell := new(Cell)
	cell.Coordinate = cellCoord
	cell.LeftEdge = Edge{Coordinates: Coordinates{X: cellCoord.X - 1, Y: cellCoord.Y}, State: isFreeEdge}
	cell.RightEdge = Edge{Coordinates: Coordinates{X: cellCoord.X + 1, Y: cellCoord.Y}, State: isFreeEdge}
	cell.UpperEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y - 1}, State: isFreeEdge}
	cell.LowerEdge = Edge{Coordinates: Coordinates{X: cellCoord.X, Y: cellCoord.Y + 1}, State: isFreeEdge}
	return cell
}
