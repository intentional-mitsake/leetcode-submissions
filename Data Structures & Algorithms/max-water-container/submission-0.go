func maxArea(heights []int) int {
	maxArea := 0
	// so the logic is pretty much:
	// the area is ltd by the height, area = height * breadth
	// max area for any container is: shortest height * breadth(dist)
	// move the shorter bar to get better height cuz:
	// have to move one, if move higher, u lose both heigh and breadth.
	l, r := 0, len(heights)-1
	var t int
	for l < r {
		// breadth: r-l(9-2=7), height: shorter one
		if heights[l] > heights[r] {
			t = heights[r] * (r-l)
			// if left bar is bigger, move the right
			// this way the max height is still highest
			r--
		} else {
			t = heights[l] * (r-l)
			// if right is bigger, save higher bar
			l++
		}
		// update if this is the max container
		if t > maxArea {
			maxArea = t
		}
	}
	return maxArea
}
