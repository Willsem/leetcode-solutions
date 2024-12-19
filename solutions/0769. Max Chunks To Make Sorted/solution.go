func maxChunksToSorted(arr []int) int {
	chunks, maxEl := 0, 0
	for i := range arr {
		maxEl = max(maxEl, arr[i])
		if maxEl == i {
			chunks++
		}
	}
	return chunks
}
