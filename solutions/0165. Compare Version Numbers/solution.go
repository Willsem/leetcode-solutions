import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")

	i1, i2 := 0, 0
	for i1 < len(ver1) || i2 < len(ver2) {
		v1Num := 0
		if i1 < len(ver1) {
			v1Num, _ = strconv.Atoi(ver1[i1])
		}

		v2Num := 0
		if i2 < len(ver2) {
			v2Num, _ = strconv.Atoi(ver2[i2])
		}

		if v1Num < v2Num {
			return -1
		} else if v1Num > v2Num {
			return 1
		}

		i1++
		i2++
	}

	return 0
}
