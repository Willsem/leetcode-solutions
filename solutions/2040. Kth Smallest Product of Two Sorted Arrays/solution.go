func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n1, n2 := len(nums1), len(nums2)
	pos1, pos2 := 0, 0

	for pos1 < n1 && nums1[pos1] < 0 {
		pos1++
	}
	for pos2 < n2 && nums2[pos2] < 0 {
		pos2++
	}

	left, right := int(-1e10), int(1e10)
	for left <= right {
		mid := (left + right) / 2
		count := 0

		i1, i2 := 0, pos2-1
		for i1 < pos1 && i2 >= 0 {
			if nums1[i1]*nums2[i2] > mid {
				i1++
			} else {
				count += pos1 - i1
				i2--
			}
		}

		i1, i2 = pos1, n2-1
		for i1 < n1 && i2 >= pos2 {
			if nums1[i1]*nums2[i2] > mid {
				i2--
			} else {
				count += i2 - pos2 + 1
				i1++
			}
		}

		i1, i2 = 0, pos2
		for i1 < pos1 && i2 < n2 {
			if nums1[i1]*nums2[i2] > mid {
				i2++
			} else {
				count += n2 - i2
				i1++
			}
		}

		i1, i2 = pos1, 0
		for i1 < n1 && i2 < pos2 {
			if nums1[i1]*nums2[i2] > mid {
				i1++
			} else {
				count += n1 - i1
				i2++
			}
		}

		if count < int(k) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return int64(left)
}
