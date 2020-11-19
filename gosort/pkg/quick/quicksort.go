package quick

/*
Худшее время: O(n2)
Среднее время: O(n log n)
*/

func Sort(list []int) {
	/*
		Насколько я понимаю быстрая сортировка сортирует inplace
	*/
	if len(list) <= 1 {
		return
	}

	split := partition(list)

	Sort(list[:split])
	Sort(list[split:])
}

func partition(list []int) int {
	/*
		Выбирается элемент - pivot, вокруг которого и будет происходить сортировка.
		Массив также делится на две части, как и в mergeSort, и в каждой части идет
		поиск элементов которые либо больше pivot, либо меньше. Запоминая их индексы
		чтобы потом переставить inplace.

		Получается поиск идет с двух сторон.
	*/

	pivot := list[len(list)/2] // выбор элемента, вокруг которого будет производиться сортировка

	leftIndex := 0              // первый элемент массива
	rightIndex := len(list) - 1 // последний элемент массива

	for {
		// идем слева направо, ищем такой элемент массива, который будет меньше pivot
		for ; list[leftIndex] < pivot; leftIndex++ {
		}

		// идем справа налево, ищем такой элемент массива, который будет больше pivot
		for ; list[rightIndex] > pivot; rightIndex-- {
		}

		// если индекс меньшего чем pivot больше чем индекс большего чем pivot, значит список точно
		// отсортирован до индекса rightIndex
		if leftIndex >= rightIndex {
			return rightIndex
		}

		// если левая часть не отсортирована, то inplace меняем элементы местами

		list[leftIndex], list[rightIndex] = list[rightIndex], list[leftIndex]
	}
}
