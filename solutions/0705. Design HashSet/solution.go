import "math"

type MyHashSet struct {
	bitmap []byte
}

func Constructor() MyHashSet {
	return MyHashSet{
		bitmap: make([]byte, 1e6>>3+1),
	}
}

func (this *MyHashSet) Add(key int) {
	index, bitmask := this.calcKey(key)
	this.bitmap[index] |= bitmask
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		index, bitmask := this.calcKey(key)
		this.bitmap[index] &= (bitmask ^ 255)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	index, bitmask := this.calcKey(key)
	return this.bitmap[index]|(bitmask^255) == 255
}

func (this *MyHashSet) calcKey(key int) (int, byte) {
	index := key >> 3
	bitmask := math.Pow(2, float64(key&7))
	return index, byte(bitmask)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
