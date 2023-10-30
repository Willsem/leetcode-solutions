type Data struct {
    Num int
    Count int
}

func sortByBits(arr []int) []int {
    nums := make([]Data, len(arr))
    for i := range arr {
        nums[i] = Data{
            Num: arr[i],
            Count: countBits(arr[i]),
        }
    }

    sort.Slice(nums, func(i, j int) bool {
        if nums[i].Count == nums[j].Count {
            return nums[i].Num < nums[j].Num
        }
        return nums[i].Count < nums[j].Count
    })

    for i := range arr {
        arr[i] = nums[i].Num
    }

    return arr
}

func countBits(num int) int {
    count := 0

    for num > 0 {
        count += num & 1
        num >>= 1
    }

    return count
}
