type ProductOfNumbers struct {
	prefixProd []int
	size       int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefixProd: []int{1},
	}
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.prefixProd = []int{1}
		p.size = 0
	} else {
		p.prefixProd = append(p.prefixProd, p.prefixProd[p.size]*num)
		p.size++
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	if k > p.size {
		return 0
	}
	return p.prefixProd[p.size] / p.prefixProd[p.size-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
