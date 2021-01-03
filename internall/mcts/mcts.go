package mcts

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string  = "A"
	minimizerSambol string  = "B"
	UCTK            float64 = 2
)

//GameState will keep track of state of game every point
// type GameState struct {
// 	playerJustMoved int
// 	board           *gamemap.Map
// 	cachedResults   [3]float64
// }

// type MCTS struct {
// 	game       *GameState
// 	iterations int
// 	movesNode  *Node
// 	UCTK       float64
// }

//SelectMove next move based on base state of the game
func SelectMove(gmap gamemap.Map) []int8 {
	rootNode := NewNode(nil, true, &gmap)
	//for i try

	for i := 0; i < 1000; i++ {
		mcts(rootNode)
	}
	fmt.Println(gmap)

	//find best option
	move := []int8{0, 0}
	bestValue := 0.0
	for _, v := range rootNode.childNodes {
		if bestValue < v.value {
			bestValue = v.value
			move = v.move
		}
		fmt.Printf("coords: %v , visits:%f , value :%f \n", v.move, v.visits, v.value)
	}

	return move
}

func mcts(node *Node) float64 {
	if node.childNodes == nil {
		/*********
		*	RollOut node
		 */
		if node.visits == 0 {
			//rollout
			node.value = node.RollOut(3)
			node.visits = 1
			return node.value
		}

		/*********
		*	Expand the Tree
		 */
		node.Expand()
	}

	if len(node.childNodes) > 0 {
		/*********
		*	Go to Leaf Node
		 */
		chossenNode := node.childNodes[0]
		chossenUcb := chossenNode.UCB1()
		//go to leafnode
		for _, n := range node.childNodes[1:] {
			if n.UCB1() > chossenUcb {
				chossenUcb = n.UCB1()
				chossenNode = n
			}
		}
		chossenNode.value += mcts(chossenNode)
		chossenNode.visits++
		return chossenNode.value
	}
	return node.value
}
