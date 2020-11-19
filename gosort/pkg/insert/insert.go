package insert

/*
сложность O(n^2)
лучшая O(n)
*/

func Sort(list []int) {
	for i := 1; i < len(list); i++ {
		x := list[i]
		j := i
		for ; j >= 1 && list[j-1] > x; j-- {
			list[j] = list[j-1]
		}
		list[j] = x
	}
}
