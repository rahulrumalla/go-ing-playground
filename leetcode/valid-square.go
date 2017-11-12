package leetcode

import "math"

//p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if len(p1) != 2 || len(p2) != 2 || len(p3) != 2 || len(p4) != 2 {
		return false
	}

	h := make(map[float64]bool, 6)

	dp12 := getDistance(p1, p2)
	dp13 := getDistance(p1, p3)
	dp14 := getDistance(p1, p4)
	dp23 := getDistance(p2, p3)
	dp24 := getDistance(p2, p4)
	dp34 := getDistance(p3, p4)

	h[dp12] = true
	h[dp13] = true
	h[dp14] = true
	h[dp23] = true
	h[dp24] = true
	h[dp34] = true

	return !h[0] && len(h) == 2 // 1 for side and 1 for diagonal
}

func getDistance(p1 []int, p2 []int) float64 {
	return math.Abs(math.Sqrt(math.Pow(float64(p2[0]-p1[0]), 2) + math.Pow(float64(p2[1]-p1[1]), 2)))
}
