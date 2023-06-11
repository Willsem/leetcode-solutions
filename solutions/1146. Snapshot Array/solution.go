type value struct {
	snapID int
	value  int
}

type SnapshotArray struct {
	values [][]value
	snap   int
}

func Constructor(length int) SnapshotArray {
	values := make([][]value, length)
	for i := range values {
		values[i] = make([]value, 0)
	}
	return SnapshotArray{
		values: values,
		snap:   0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	last := len(this.values[index]) - 1
	if last >= 0 && this.values[index][last].snapID == this.snap {
		this.values[index][last].value = val
	} else {
		this.values[index] = append(this.values[index], value{
			snapID: this.snap,
			value:  val,
		})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snap++
	return this.snap - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if len(this.values[index]) == 0 || this.values[index][0].snapID > snap_id {
		return 0
	}

	return this.values[index][binarySearch(0, len(this.values[index]), func(i int) bool {
		return this.values[index][i].snapID <= snap_id
	})].value
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
