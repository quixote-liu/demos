package sort

var tests = []struct {
	name   string
	nums   []int
	expect []int
}{
	{
		"normal numbers",
		[]int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3},
		[]int{1, 3, 4, 5, 6, 6, 6, 8, 9, 14, 25, 49},
	},
	{
		"the same elements for numbers",
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
}
