import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var sum int64
	count := 0
	positiveMinimum := int(math.MaxInt32)
	negativeMaximum := int(math.MinInt32)

	for _, nodeValue := range nums {
		operatedNodeValue := nodeValue ^ k
		sum += int64(nodeValue)
		netChange := operatedNodeValue - nodeValue

		if netChange > 0 {
			if netChange < positiveMinimum {
				positiveMinimum = netChange
			}
			sum += int64(netChange)
			count++
		} else {
			if netChange > negativeMaximum {
				negativeMaximum = netChange
			}
		}
	}

	if count%2 == 0 {
		return sum
	}

	return max(sum-int64(positiveMinimum), sum+int64(negativeMaximum))
}
