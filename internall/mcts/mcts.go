package mcts

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string  = "A"
	minimizerSambol string  = "B"
	uctk            float64 = 2
)

//SelectMove next move based on base state of the game
func SelectMove(gmap gamemap.Map, maximizer string) []int8 {
	rootNode := NewNode(nil, false, &gmap)
	maximizerSambol = maximizer
	if maximizerSambol == "A" {
		minimizerSambol = "B"
	} else {
		minimizerSambol = "A"
	}
	//for i try

	for i := 0; i < 15000; i++ {
		mcts(rootNode)
	}

	fmt.Println(gmap)
	//find best option
	for _, v := range rootNode.children {
		fmt.Printf("coords: %v , visits:%f , value :%f UCB:%f AVG:%f\n", v.causingAction, v.visits, v.value, v.UCB1(1.4), v.value/v.visits)
	}

	return rootNode.getBestChild().causingAction
}

func mcts(node *Node) {
	node = selectNode(node)
	value := node.RollOut(3)

	// if len(extractRemainingMoves(node.gmap)) == 0 && node.value != 0 {
	// 	value *= 100
	// }

	node.visits++
	node.value = value
	for node.parentNode != nil {
		node = node.parentNode
		node.visits++
		node.value += value
	}
}

func selectNode(node *Node) *Node {

	// for !node.IsTerminal() {
	// 	if !node.IsFullyExpanded() {
	// 		return node.Expand()
	// 	}
	// 	node = node.getBestChild()
	// }
	// return node

	// for node.children != nil {
	// 	node = node.getBestChild()
	// }
	// if len(extractRemainingMoves(node.gmap)) > 0 && node.visits == 1 || node.depth == 0 {
	// 	node = node.Expand()
	// }
	// for node != nil && (node.depth == 0 || len(extractRemainingMoves(node.gmap)) > 0) {
	// 	if node.children == nil {
	// 		node= node.Expand()
	// 	}
	// 	return  node.getBestChild()
	// }

	for len(extractRemainingMoves(node.gmap)) > 0 {
		if node.children != nil {
	node = node.getBestChild()
		} else {
			return node.Expand()
		}
	}
	return node
}

//SelectMoveRecursive next move based on base state of the game
func SelectMoveRecursive(gmap gamemap.Map, maximizer string) []int8 {
	rootNode := NewNode(nil, false, &gmap)
	maximizerSambol = maximizer
	if maximizerSambol == "A" {
		minimizerSambol = "B"
	} else {
		minimizerSambol = "A"
	}
	//for i try
	fmt.Println(gmap)

	for i := 0; i < 100000; i++ {
		mctsRecursive(rootNode)
	}

	//find best option
	for _, v := range rootNode.children {
		fmt.Printf("coords: %v , visits:%f , value :%f UCB:%f AVG:%f\n", v.causingAction, v.visits, v.value, v.UCB1(1.4), v.value/v.visits)
	}

	return rootNode.getBestChild().causingAction
}

func mctsRecursive(node *Node) float64 {
	if node.children == nil {
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

	if len(node.children) > 0 {
		/*********
		*	Go to Leaf Node
		 */
		chossenNode := node.children[0]
		chossenUcb := chossenNode.UCB1(1.4)
		//go to leafnode
		for _, n := range node.children {
			v := n.UCB1(1.4)
			if v > chossenUcb {
				chossenUcb = v
				chossenNode = n
			}
		}
		node.value += mctsRecursive(chossenNode)
		node.visits++
		return node.value
	}
	node.value += node.RollOut(5)
	node.visits++
	return node.value
}
