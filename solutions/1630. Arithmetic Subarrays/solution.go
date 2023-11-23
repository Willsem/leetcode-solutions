import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		ans[i] = true
		if r[i]-l[i] > 0 {
			temp := make([]int, 0)
			temp = append(temp, nums[l[i]:r[i]+1]...)
			sort.Ints(temp)

			var diff int = temp[1] - temp[0]
			for j := 1; j < len(temp); j++ {
				if temp[j]-temp[j-1] != diff {
					ans[i] = false
					break
				}
			}
		} else {
			ans[i] = false
		}
	}

	return ans
}
