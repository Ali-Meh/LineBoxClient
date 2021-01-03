package mcts

import "github.com/ali-meh/LineBoxClient/internall/gamemap"

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string = "A"
	minimizerSambol string = "B"
)

//GameState will keep track of state of game every point
type GameState struct {
	playerJustMoved int
	board           *gamemap.Map
	cachedResults   [3]float64
}

type MCTS struct {
	game       *GameState
	iterations int
	movesNode  *Node
	UCTK       float64
}

//SelectMove next move based on base state of the game
// func SelectMove(gmap gamemap.Map, depth int, maximizer string) ([]int8, *gamemap.Tree) {

// 	return nil, nil
// }
