package mcts

import (
	"fmt"

	"github.com/ali-meh/LineBoxClient/internall/gamemap"
)

var (
	//maximizerSambol indecates the maximizer player symbol
	maximizerSambol string  = "A"
	minimizerSambol string  = "B"
	uctk            float64 = 1.4
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
	rootNode.Expand()
	rootNode.visits++

	for i := 0; i < 15000; i++ {
		mcts(rootNode)
	}

	fmt.Println(gmap)
	//find best option
	var bestNode *Node
	bestAvgNode := rootNode.children[0]
	for _, v := range rootNode.children {
		fmt.Printf("coords: %v\t\tvisits:%7.0f\t\tvalue:%10.0f\t\tUCB:%f\t\tUCB0:%f\t\tAVG:%f\n", v.causingAction, v.visits, v.value, v.UCB1(uctk), v.UCB1(0), v.value/v.visits)
		if v.UCB1(0) == 0 && v.value/v.visits != 0 {
			bestNode = v
		}
		if bestAvgNode.value/bestAvgNode.visits < v.value/v.visits {
			// if
			bestAvgNode = v
		}
	}
	fmt.Println("Best Avg Node", bestAvgNode.causingAction)
	if bestNode == nil {
		return bestAvgNode.causingAction
	}
	fmt.Println("Best Node", bestNode.causingAction)
	return bestNode.causingAction
}

func mcts(root *Node) {
	leaf := selectLeaf(root)
	// println(leaf.causingAction)
	result := leaf.RollOut(5)
	// println(result)
	leaf.backpropagate(result)
}

func selectLeaf(node *Node) *Node {

	for !node.isLeaf() && !node.isTerminal() {
		if !node.isFullyExpanded() {
			return node.Expand()
		}
		node = node.getBestChild(uctk)
	}
	return node

	// for !node.isTerminal() {
	// 	if !node.isFullyExpanded() {
	// 		return node.Expand()
	// 	}
	// 	node = node.getBestChild(uctk)
	// }
	// return node
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
		fmt.Printf("coords: %v , visits:%f , value :%f UCB:%f AVG:%f\n", v.causingAction, v.visits, v.value, v.UCB1(uctk), v.value/v.visits)
	}

	return rootNode.getBestChild(uctk).causingAction
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
		chossenUcb := chossenNode.UCB1(uctk)
		//go to leafnode
		for _, n := range node.children {
			v := n.UCB1(uctk)
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
