package heap

func Sort(list []int) {
	if len(list) < 2 {
		return
	}

	heapify(list)

	list[0], list[len(list)-1] = list[len(list)-1], list[0]

	Sort(list[:len(list)-1])
}

func heapify(list []int) {
	if len(list) < 2 {
		return
	}

	if len(list) == 2 {
		if list[0] < list[1] {
			list[0], list[1] = list[1], list[0]
		}
		return
	}

	if len(list) > 3 {
		heapify(list[1:])
		heapify(list[2:])
	}

	if list[1] > list[2] {
		if list[0] < list[1] {
			list[0], list[1] = list[1], list[0]
		}
	} else {
		if list[0] < list[2] {
			list[0], list[2] = list[2], list[0]
		}
	}
}
