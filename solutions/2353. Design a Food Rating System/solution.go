import "container/heap"

type Food struct {
	Name   string
	Rating int
}

type FoodRatings struct {
	FoodRatingMap  map[string]int
	FoodCuisineMap map[string]string
	CuisineHeapMap map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRatingMap := make(map[string]int)
	foodCuisineMap := make(map[string]string)
	cuisineHeapMap := make(map[string]*MaxHeap)

	for i := 0; i < len(foods); i++ {
		foodRatingMap[foods[i]] = ratings[i]
		foodCuisineMap[foods[i]] = cuisines[i]

		if _, ok := cuisineHeapMap[cuisines[i]]; !ok {
			cuisineHeapMap[cuisines[i]] = &MaxHeap{}
		}

		heap.Push(cuisineHeapMap[cuisines[i]], Food{foods[i], ratings[i]})
	}

	return FoodRatings{
		FoodRatingMap:  foodRatingMap,
		FoodCuisineMap: foodCuisineMap,
		CuisineHeapMap: cuisineHeapMap,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.FoodRatingMap[food] = newRating
	cuisine := this.FoodCuisineMap[food]
	heap.Push(this.CuisineHeapMap[cuisine], Food{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	highRatedFood := (*this.CuisineHeapMap[cuisine])[0]
	for highRatedFood.Rating != this.FoodRatingMap[highRatedFood.Name] {
		heap.Pop(this.CuisineHeapMap[cuisine])
		highRatedFood = (*this.CuisineHeapMap[cuisine])[0]
	}

	return highRatedFood.Name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

type MaxHeap []Food

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	if h[i].Rating == h[j].Rating {
		return h[i].Name < h[j].Name
	}
	return h[i].Rating > h[j].Rating
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Food))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
