package mcts

import (
	"fmt"
	"math"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string  = "A"
	minimizerSambol string  = "B"
	uctk            float64 = 1 / math.Sqrt(2)
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
func SelectMove(gmap gamemap.Map, maximizer string) []int8 {
	rootNode := NewNode(nil, true, &gmap)
	maximizerSambol = maximizer
	if maximizerSambol == "A" {
		minimizerSambol = "B"
	} else {
		minimizerSambol = "A"
	}
	//for i try

	for i := 0; i < 20000; i++ {
		mcts(rootNode)
		t++
	}
	fmt.Println(gmap)

	//find best option
	move := rootNode.childNodes[0].move
	bestValue := rootNode.childNodes[0].value
	fmt.Printf("coords: %v , visits:%f , value :%f \n", rootNode.childNodes[0].move, rootNode.childNodes[0].visits, rootNode.childNodes[0].value)

	for _, v := range rootNode.childNodes[1:] {
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
			node.value = node.RollOut(5)
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
		node.value += mcts(chossenNode)
		node.visits++
		return node.value
	}
	node.value += node.RollOut(5)
	node.visits++
	return node.value
}
