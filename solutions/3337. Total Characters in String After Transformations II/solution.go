const MOD = 1e9 + 7

func lengthAfterTransformations(s string, t int, nums []int) int {
	arr := make([]int, 26)
	for _, ch := range s {
		arr[ch-'a']++
	}

	mat := make([][]int, 26)
	for i := 0; i < 26; i++ {
		mat[i] = make([]int, 26)
	}

	for i := 0; i < 26; i++ {
		if nums[i] > 0 {
			for j := 1; j <= nums[i]; j++ {
				to := (i + j) % 26
				mat[to][i] = (mat[to][i] + 1) % MOD
			}
		}
	}

	mul := func(a, b [][]int) [][]int {
		res := make([][]int, 26)
		for i := 0; i < 26; i++ {
			res[i] = make([]int, 26)
			for j := 0; j < 26; j++ {
				for k := 0; k < 26; k++ {
					res[i][j] = (res[i][j] + a[i][k]*b[k][j]%MOD) % MOD
				}
			}
		}
		return res
	}

	pow := func(m [][]int, exp int) [][]int {
		res := make([][]int, 26)
		for i := 0; i < 26; i++ {
			res[i] = make([]int, 26)
			res[i][i] = 1
		}
		base := make([][]int, 26)
		for i := 0; i < 26; i++ {
			base[i] = make([]int, 26)
			copy(base[i], m[i])
		}

		for exp > 0 {
			if exp%2 == 1 {
				res = mul(res, base)
			}
			base = mul(base, base)
			exp /= 2
		}
		return res
	}

	apply := func(m [][]int, v []int) []int {
		res := make([]int, 26)
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				res[i] = (res[i] + m[i][j]*v[j]%MOD) % MOD
			}
		}
		return res
	}

	matT := pow(mat, t)
	final := apply(matT, arr)

	sum := 0
	for _, val := range final {
		sum = (sum + val) % MOD
	}
	return sum
}
