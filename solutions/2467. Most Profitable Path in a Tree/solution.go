import "math"

type ProblemContext struct {
	Amount []int
	Bob    int
	Tree   [][]int
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	tree := make([][]int, n)

	for _, e := range edges {
		tree[e[0]] = append(tree[e[0]], e[1])
		tree[e[1]] = append(tree[e[1]], e[0])
	}

	ctx := &ProblemContext{Amount: amount, Bob: bob, Tree: tree}
	aliceProfit, _ := dpRec(0, -1, 0, ctx)
	return aliceProfit
}

func dpRec(node int, parent int, currentDepth int, ctx *ProblemContext) (aliceProfit int, bobDistance int) {
	if len(ctx.Tree[node]) == 1 && currentDepth != 0 {
		if node != ctx.Bob {
			aliceProfit = ctx.Amount[node]
			bobDistance = len(ctx.Tree)
		}
		return
	}

	bobDistance = len(ctx.Tree)
	aliceProfit = math.MinInt
	if node == ctx.Bob {
		bobDistance = 0
	}

	for _, neighbor := range ctx.Tree[node] {
		if neighbor != parent {
			nProfit, nBobDistance := dpRec(neighbor, node, currentDepth+1, ctx)
			bobDistance = min(bobDistance, nBobDistance+1)
			aliceProfit = max(aliceProfit, nProfit)
		}
	}

	if currentDepth < bobDistance {
		aliceProfit += ctx.Amount[node]
	} else if currentDepth == bobDistance {
		aliceProfit += ctx.Amount[node] / 2
	}

	return
}
