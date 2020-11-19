package shell

/*
Худшее O(n^2)
Лучшее O(n log n)
*/

func Sort(list []int) {
	for gap := len(list) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(list); i++ {
			x := list[i]
			j := i
			for ; j >= gap && list[j-gap] > x; j -= gap {
				list[j] = list[j-gap]
			}
			list[j] = x
		}
	}
}
